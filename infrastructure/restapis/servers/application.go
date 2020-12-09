package servers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/xmn-services/buckets-network/application/commands"
	identities_app "github.com/xmn-services/buckets-network/application/commands/identities"
	"github.com/xmn-services/buckets-network/application/servers"
	"github.com/xmn-services/buckets-network/domain/memory/peers/peer"
	"github.com/xmn-services/buckets-network/infrastructure/restapis/shared"
)

type application struct {
	cmdApp                commands.Application
	updateIdentityBuilder identities_app.UpdateBuilder
	peerAdapter           peer.Adapter
	router                *mux.Router
	server                *http.Server
	authApps              map[string]identities_app.Application
	maxUploadFileSize     int64
	waitPeriod            time.Duration
	port                  uint
}

func createApplication(
	cmdApp commands.Application,
	updateIdentityBuilder identities_app.UpdateBuilder,
	peerAdapter peer.Adapter,
	router *mux.Router,
	maxUploadFileSize int64,
	waitPeriod time.Duration,
	port uint,
) servers.Application {
	out := application{
		cmdApp:                cmdApp,
		updateIdentityBuilder: updateIdentityBuilder,
		peerAdapter:           peerAdapter,
		router:                router,
		server:                nil,
		authApps:              map[string]identities_app.Application{},
		maxUploadFileSize:     maxUploadFileSize,
		waitPeriod:            waitPeriod,
		port:                  port,
	}

	// add routes:
	out.router.HandleFunc("/chains", out.retrieveChain).Methods(http.MethodGet, http.MethodOptions)
	out.router.HandleFunc("/chains/{index:[0-9]+}", out.retrieveChainAtIndex).Methods(http.MethodGet, http.MethodOptions)
	out.router.HandleFunc("/peers", out.retrievePeers).Methods(http.MethodGet, http.MethodOptions)
	out.router.HandleFunc("/peers", out.savePeer).Methods(http.MethodPost, http.MethodOptions)
	out.router.HandleFunc("/storages/{bucket_hash:[0-9a-f]+}/{chunk_hash:[0-9a-f]+}", out.retrieveChunk).Methods(http.MethodGet, http.MethodOptions)
	out.router.HandleFunc("/identities", out.newIdentity).Methods(http.MethodPost, http.MethodOptions)

	// middleware:
	out.router.Use(mux.CORSMethodMiddleware(out.router))

	// identities:
	identityRouter := out.router.PathPrefix("/identities").Subrouter()
	identityRouter.HandleFunc("", out.retrieveIdentity).Methods(http.MethodGet, http.MethodOptions)
	identityRouter.HandleFunc("", out.updateIdentity).Methods(http.MethodPut, http.MethodOptions)
	identityRouter.HandleFunc("", out.deleteIdentity).Methods(http.MethodDelete, http.MethodOptions)

	identityRouter.HandleFunc("/access/{bucket_hash:[0-9a-f]+}", out.retrieveAccess).Methods(http.MethodGet, http.MethodOptions)
	identityRouter.HandleFunc("/access/{bucket_hash:[0-9a-f]+}", out.saveAccess).Methods(http.MethodPost, http.MethodOptions)
	identityRouter.HandleFunc("/access/{bucket_hash:[0-9a-f]+}", out.deleteAccess).Methods(http.MethodDelete, http.MethodOptions)

	identityRouter.HandleFunc("/access/{bucket_hash:[0-9a-f]+}/bucket", out.deleteAccessBucket).Methods(http.MethodDelete, http.MethodOptions)
	identityRouter.HandleFunc("/access/{bucket_hash:[0-9a-f]+}/bucket", out.retrieveAccessBucket).Methods(http.MethodGet, http.MethodOptions)
	identityRouter.HandleFunc("/access/{bucket_hash:[0-9a-f]+}/bucket", out.extractAccessBucket).Methods(http.MethodPost, http.MethodOptions)

	identityRouter.HandleFunc("/lists", out.retrieveLists).Methods(http.MethodGet, http.MethodOptions)
	identityRouter.HandleFunc("/lists", out.saveList).Methods(http.MethodPost, http.MethodOptions)
	identityRouter.HandleFunc("/lists/{list_hash:[0-9a-f]+}", out.retrieveList).Methods(http.MethodGet, http.MethodOptions)
	identityRouter.HandleFunc("/lists/{list_hash:[0-9a-f]+}", out.updateList).Methods(http.MethodPut, http.MethodOptions)
	identityRouter.HandleFunc("/lists/{list_hash:[0-9a-f]+}", out.deleteList).Methods(http.MethodDelete, http.MethodOptions)

	identityRouter.HandleFunc("/lists/{list_hash:[0-9a-f]+}/contacts", out.retrieveListContacts).Methods(http.MethodGet, http.MethodOptions)
	identityRouter.HandleFunc("/lists/{list_hash:[0-9a-f]+}/contacts/{contact_hash:[0-9a-f]+}", out.retrieveListContact).Methods(http.MethodGet, http.MethodOptions)
	identityRouter.HandleFunc("/lists/{list_hash:[0-9a-f]+}/contacts/{contact_hash:[0-9a-f]+}", out.updateListContact).Methods(http.MethodPut, http.MethodOptions)
	identityRouter.HandleFunc("/lists/{list_hash:[0-9a-f]+}/contacts/{contact_hash:[0-9a-f]+}", out.deleteListContact).Methods(http.MethodDelete, http.MethodOptions)

	identityRouter.HandleFunc("/lists/{list_hash:[0-9a-f]+}/contacts/{contact_hash:[0-9a-f]+}/buckets", out.retrieveListContactBuckets).Methods(http.MethodGet, http.MethodOptions)
	identityRouter.HandleFunc("/lists/{list_hash:[0-9a-f]+}/contacts/{contact_hash:[0-9a-f]+}/buckets", out.saveListContactBucket).Methods(http.MethodPost, http.MethodOptions)
	identityRouter.HandleFunc("/lists/{list_hash:[0-9a-f]+}/contacts/{contact_hash:[0-9a-f]+}/buckets/{bucket_hash:[0-9a-f]+}/", out.retrieveListContactBucket).Methods(http.MethodGet, http.MethodOptions)
	identityRouter.HandleFunc("/lists/{list_hash:[0-9a-f]+}/contacts/{contact_hash:[0-9a-f]+}/buckets/{bucket_hash:[0-9a-f]+}/", out.deleteListContactBucket).Methods(http.MethodDelete, http.MethodOptions)

	identityRouter.HandleFunc("/miners/test/{difficulty:[0-9]+}", out.identityMinerTest).Methods(http.MethodGet, http.MethodOptions)
	identityRouter.HandleFunc("/miners/block/{hash:[0-9a-f]+}", out.identityMinerBlock).Methods(http.MethodGet, http.MethodOptions)
	identityRouter.HandleFunc("/miners/link/{hash:[0-9a-f]+}", out.identityMinerLink).Methods(http.MethodGet, http.MethodOptions)

	identityRouter.HandleFunc("/chains", out.initIdentityChain).Methods(http.MethodPost, http.MethodOptions)
	identityRouter.HandleFunc("/chains/blocks/{additional:[0-9]+}", out.identityChainMineBlocks).Methods(http.MethodPost, http.MethodOptions)
	identityRouter.HandleFunc("/chains/links/{additional:[0-9]+}", out.identityChainMineLinks).Methods(http.MethodPost, http.MethodOptions)

	identityRouter.HandleFunc("/storages/{bucket_hash:[0-9a-f]+}", out.saveChunk).Methods(http.MethodPost, http.MethodOptions)
	identityRouter.HandleFunc("/storages/{bucket_hash:[0-9a-f]+}/{chunk_hash:[0-9a-f]+}", out.deleteChunk).Methods(http.MethodDelete, http.MethodOptions)
	identityRouter.HandleFunc("/storages/{bucket_hash:[0-9a-f]+}", out.deleteBucketChunks).Methods(http.MethodDelete, http.MethodOptions)

	// identity middleware:
	identityRouter.Use(out.authenticateMiddleWare)

	return &out
}

// Start starts the application
func (app *application) Start() error {
	app.server = &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%d", app.port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      app.router,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := app.server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	//block until we receive signal
	<-c

	return app.Stop()
}

// Stop stops the application
func (app *application) Stop() error {
	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), app.waitPeriod)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	app.server.Shutdown(ctx)

	// shut down:
	log.Println("shutting down")
	os.Exit(0)
	return nil
}

func (app *application) authenticateMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get(shared.TokenHeadKeyname)
		auth, err := shared.Base64ToAuthenticate(token)
		if err != nil {
			renderError(w, err, []byte(internalErrorOutput))
			return
		}

		authApp, err := app.cmdApp.Current().Authenticate(
			auth.Name,
			auth.Seed,
			auth.Password,
		)

		if err != nil {
			renderInvalidAuthentication(w, err, []byte("invalid authentication"))
			return
		}

		app.authApps[token] = authApp

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func (app *application) newIdentity(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		renderError(w, err, []byte(internalErrorOutput))
		return
	}

	identity, err := shared.URLValuesToIdentity(r.Form)
	if err != nil {
		renderError(w, err, []byte(internalErrorOutput))
		return
	}

	err = app.cmdApp.Current().NewIdentity(
		identity.Authenticate.Name,
		identity.Authenticate.Password,
		identity.Authenticate.Seed,
		identity.Root,
	)

	if err != nil {
		renderError(w, err, []byte(internalErrorOutput))
		return
	}

	renderSuccess(w, []byte(successPostOutput))
}

func (app *application) retrievePeers(w http.ResponseWriter, r *http.Request) {
	peers, err := app.cmdApp.Sub().Peers().Retrieve()
	if err != nil {
		renderError(w, err, []byte(internalErrorOutput))
		return
	}

	js, err := json.Marshal(peers)
	if err != nil {
		renderError(w, err, []byte(internalErrorOutput))
		return
	}

	renderSuccess(w, js)
}

func (app *application) savePeer(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		renderError(w, err, []byte(internalErrorOutput))
		return
	}

	peer, err := app.peerAdapter.URLValuesToPeer(r.Form)
	if err != nil {
		renderError(w, err, []byte(internalErrorOutput))
		return
	}

	err = app.cmdApp.Sub().Peers().Save(peer)
	if err != nil {
		renderError(w, err, []byte(internalErrorOutput))
		return
	}

	renderSuccess(w, []byte(successPostOutput))
}

func (app *application) retrieveChain(w http.ResponseWriter, r *http.Request) {
	chain, err := app.cmdApp.Sub().Chain().Retrieve()
	if err != nil {
		renderError(w, err, []byte(internalErrorOutput))
		return
	}

	js, err := json.Marshal(chain)
	if err != nil {
		renderError(w, err, []byte(internalErrorOutput))
		return
	}

	renderSuccess(w, js)
}

func (app *application) retrieveChainAtIndex(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if indexAsStr, ok := vars["index"]; ok {
		index, err := strconv.Atoi(indexAsStr)
		if err != nil {
			renderError(w, err, []byte(internalErrorOutput))
			return
		}

		chain, err := app.cmdApp.Sub().Chain().RetrieveAtIndex(uint(index))
		if err != nil {
			renderError(w, err, []byte(internalErrorOutput))
			return
		}

		js, err := json.Marshal(chain)
		if err != nil {
			renderError(w, err, []byte(internalErrorOutput))
			return
		}

		renderSuccess(w, js)
		return
	}

	err := errors.New(missingIndexErrorOutput)
	renderError(w, err, []byte(internalErrorOutput))
}

func (app *application) identityMinerTest(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get(shared.TokenHeadKeyname)
	defer app.deleteAuthApp(token)

	if appIdentity, ok := app.authApps[token]; ok {
		vars := mux.Vars(r)
		if diffAsStr, ok := vars["difficulty"]; ok {
			difficulty, err := strconv.Atoi(diffAsStr)
			if err != nil {
				renderError(w, err, []byte(internalErrorOutput))
				return
			}

			results, err := appIdentity.Sub().Miner().Test(uint(difficulty))
			if err != nil {
				renderError(w, err, []byte(internalErrorOutput))
				return
			}

			renderSuccess(w, []byte(results))
			return
		}

		err := errors.New(missingIndexErrorOutput)
		renderError(w, err, []byte(internalErrorOutput))
		return
	}

	str := fmt.Sprintf(authErrorOutput, token)
	err := errors.New(str)
	renderError(w, err, []byte(str))
}

func (app *application) identityMinerBlock(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get(shared.TokenHeadKeyname)
	defer app.deleteAuthApp(token)

	if appIdentity, ok := app.authApps[token]; ok {
		vars := mux.Vars(r)
		if blockHashStr, ok := vars["hash"]; ok {
			results, err := appIdentity.Sub().Miner().Block(blockHashStr)
			if err != nil {
				renderError(w, err, []byte(internalErrorOutput))
				return
			}

			renderSuccess(w, []byte(results))
			return
		}

		err := errors.New(missingHashErrorOutput)
		renderError(w, err, []byte(internalErrorOutput))
		return
	}

	str := fmt.Sprintf(authErrorOutput, token)
	err := errors.New(str)
	renderError(w, err, []byte(str))
}

func (app *application) identityMinerLink(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get(shared.TokenHeadKeyname)
	defer app.deleteAuthApp(token)

	if appIdentity, ok := app.authApps[token]; ok {
		vars := mux.Vars(r)
		if linkHashStr, ok := vars["hash"]; ok {
			results, err := appIdentity.Sub().Miner().Link(linkHashStr)
			if err != nil {
				renderError(w, err, []byte(internalErrorOutput))
				return
			}

			renderSuccess(w, []byte(results))
			return
		}

		err := errors.New(missingHashErrorOutput)
		renderError(w, err, []byte(internalErrorOutput))
		return
	}

	str := fmt.Sprintf(authErrorOutput, token)
	err := errors.New(str)
	renderError(w, err, []byte(str))
}

func (app *application) initIdentityChain(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get(shared.TokenHeadKeyname)
	defer app.deleteAuthApp(token)

	if appIdentity, ok := app.authApps[token]; ok {
		err := r.ParseForm()
		if err != nil {
			renderError(w, err, []byte(internalErrorOutput))
			return
		}

		initChain, err := shared.URLValuesToInitChain(r.Form)
		if err != nil {
			renderError(w, err, []byte(internalErrorOutput))
			return
		}

		err = appIdentity.Sub().Chain().Init(
			initChain.MiningValue,
			initChain.BaseDifficulty,
			initChain.IncreasePerBucket,
			initChain.LinkDifficulty,
			initChain.RootAdditionalBuckets,
			initChain.HeadAdditionalBuckets,
		)

		if err != nil {
			renderError(w, err, []byte(internalErrorOutput))
			return
		}

		renderSuccess(w, []byte(successPostOutput))
		return
	}

	str := fmt.Sprintf(authErrorOutput, token)
	err := errors.New(str)
	renderError(w, err, []byte(str))
}

func (app *application) identityChainMineBlocks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if additionalStr, ok := vars["additional"]; ok {
		additional, err := strconv.Atoi(additionalStr)
		if err != nil {
			renderError(w, err, []byte(internalErrorOutput))
			return
		}

		token := r.Header.Get(shared.TokenHeadKeyname)
		defer app.deleteAuthApp(token)

		if appIdentity, ok := app.authApps[token]; ok {
			err = appIdentity.Sub().Chain().Block(uint(additional))
			if err != nil {
				renderError(w, err, []byte(internalErrorOutput))
				return
			}

			renderSuccess(w, []byte(successPostOutput))
			return
		}

		str := fmt.Sprintf(authErrorOutput, token)
		err = errors.New(str)
		renderError(w, err, []byte(str))
		return
	}

	err := errors.New(missingHashErrorOutput)
	renderError(w, err, []byte(internalErrorOutput))
}

func (app *application) identityChainMineLinks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if additionalStr, ok := vars["additional"]; ok {
		additional, err := strconv.Atoi(additionalStr)
		if err != nil {
			renderError(w, err, []byte(internalErrorOutput))
			return
		}

		token := r.Header.Get(shared.TokenHeadKeyname)
		defer app.deleteAuthApp(token)

		if appIdentity, ok := app.authApps[token]; ok {
			err = appIdentity.Sub().Chain().Link(uint(additional))
			if err != nil {
				renderError(w, err, []byte(internalErrorOutput))
				return
			}

			renderSuccess(w, []byte(successPostOutput))
			return
		}

		str := fmt.Sprintf(authErrorOutput, token)
		err = errors.New(str)
		renderError(w, err, []byte(str))
		return
	}

	err := errors.New(missingHashErrorOutput)
	renderError(w, err, []byte(internalErrorOutput))
}

func (app *application) retrieveChunk(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if bucketHash, ok := vars["bucket_hash"]; ok {
		if chunkHash, ok := vars["chunk_hash"]; ok {
			data, err := app.cmdApp.Sub().Storage().Retrieve(bucketHash, chunkHash)
			if err != nil {
				renderError(w, err, []byte(internalErrorOutput))
				return
			}

			renderSuccess(w, data)
			return
		}

		err := errors.New(missingChunkHashErrorOutput)
		renderError(w, err, []byte(internalErrorOutput))
	}

	err := errors.New(missingBucketHashErrorOutput)
	renderError(w, err, []byte(internalErrorOutput))

}

func (app *application) retrieveIdentity(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get(shared.TokenHeadKeyname)
	defer app.deleteAuthApp(token)

	if appIdentity, ok := app.authApps[token]; ok {
		identity, err := appIdentity.Current().Retrieve()
		if err != nil {
			renderError(w, err, []byte(internalErrorOutput))
			return
		}

		js, err := json.Marshal(identity)
		if err != nil {
			renderError(w, err, []byte(internalErrorOutput))
			return
		}

		renderSuccess(w, js)
		return
	}

	str := fmt.Sprintf(authErrorOutput, token)
	err := errors.New(str)
	renderError(w, err, []byte(str))
}

func (app *application) updateIdentity(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get(shared.TokenHeadKeyname)
	defer app.deleteAuthApp(token)

	if appIdentity, ok := app.authApps[token]; ok {
		err := r.ParseForm()
		if err != nil {
			renderError(w, err, []byte(internalErrorOutput))
			return
		}

		identity, err := shared.URLValuesToIdentity(r.Form)
		if err != nil {
			renderError(w, err, []byte(internalErrorOutput))
			return
		}

		updateIdentityBuilder := app.updateIdentityBuilder.Create()
		if identity.Authenticate.Seed != "" {
			updateIdentityBuilder.WithSeed(identity.Authenticate.Seed)
		}

		if identity.Authenticate.Name != "" {
			updateIdentityBuilder.WithName(identity.Authenticate.Name)
		}

		if identity.Authenticate.Password != "" {
			updateIdentityBuilder.WithPassword(identity.Authenticate.Password)
		}

		if identity.Root != "" {
			updateIdentityBuilder.WithRoot(identity.Root)
		}

		updateIdentity, err := updateIdentityBuilder.Now()
		if err != nil {
			renderError(w, err, []byte(internalErrorOutput))
			return
		}

		err = appIdentity.Current().Update(updateIdentity)
		if err != nil {
			renderError(w, err, []byte(internalErrorOutput))
			return
		}

		renderSuccess(w, []byte(successPostOutput))
		return
	}

	str := fmt.Sprintf(authErrorOutput, token)
	err := errors.New(str)
	renderError(w, err, []byte(str))
}

func (app *application) deleteIdentity(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get(shared.TokenHeadKeyname)
	defer app.deleteAuthApp(token)

	if appIdentity, ok := app.authApps[token]; ok {
		err := appIdentity.Current().Delete()
		if err != nil {
			renderError(w, err, []byte(internalErrorOutput))
			return
		}

		renderSuccess(w, []byte(successPostOutput))
		return
	}

	str := fmt.Sprintf(authErrorOutput, token)
	err := errors.New(str)
	renderError(w, err, []byte(str))
}

func (app *application) retrieveAccess(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get(shared.TokenHeadKeyname)
	defer app.deleteAuthApp(token)

	if appIdentity, ok := app.authApps[token]; ok {
		vars := mux.Vars(r)
		if bucketHashStr, ok := vars["bucket_hash"]; ok {
			access, err := appIdentity.Sub().Access().Retrieve(bucketHashStr)
			if err != nil {
				renderError(w, err, []byte(internalErrorOutput))
				return
			}

			js, err := json.Marshal(access)
			if err != nil {
				renderError(w, err, []byte(internalErrorOutput))
				return
			}

			// success:
			renderSuccess(w, []byte(js))
			return
		}

		err := errors.New(missingBucketHashErrorOutput)
		renderError(w, err, []byte(internalErrorOutput))
		return
	}

	str := fmt.Sprintf(authErrorOutput, token)
	err := errors.New(str)
	renderError(w, err, []byte(str))
}

func (app *application) saveAccess(w http.ResponseWriter, r *http.Request) {

}

func (app *application) deleteAccess(w http.ResponseWriter, r *http.Request) {

}

func (app *application) deleteAccessBucket(w http.ResponseWriter, r *http.Request) {

}

func (app *application) retrieveAccessBucket(w http.ResponseWriter, r *http.Request) {

}

func (app *application) extractAccessBucket(w http.ResponseWriter, r *http.Request) {

}

func (app *application) retrieveLists(w http.ResponseWriter, r *http.Request) {

}

func (app *application) saveList(w http.ResponseWriter, r *http.Request) {

}

func (app *application) retrieveList(w http.ResponseWriter, r *http.Request) {

}

func (app *application) updateList(w http.ResponseWriter, r *http.Request) {

}

func (app *application) deleteList(w http.ResponseWriter, r *http.Request) {

}

func (app *application) retrieveListContacts(w http.ResponseWriter, r *http.Request) {

}

func (app *application) retrieveListContact(w http.ResponseWriter, r *http.Request) {

}

func (app *application) updateListContact(w http.ResponseWriter, r *http.Request) {

}

func (app *application) deleteListContact(w http.ResponseWriter, r *http.Request) {

}

func (app *application) retrieveListContactBuckets(w http.ResponseWriter, r *http.Request) {

}

func (app *application) saveListContactBucket(w http.ResponseWriter, r *http.Request) {

}

func (app *application) retrieveListContactBucket(w http.ResponseWriter, r *http.Request) {

}

func (app *application) deleteListContactBucket(w http.ResponseWriter, r *http.Request) {

}

func (app *application) saveChunk(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get(shared.TokenHeadKeyname)
	defer app.deleteAuthApp(token)

	if appIdentity, ok := app.authApps[token]; ok {
		vars := mux.Vars(r)
		if bucketHashStr, ok := vars["bucket_hash"]; ok {
			err := r.ParseMultipartForm(app.maxUploadFileSize)
			if err != nil {
				renderError(w, err, []byte(internalErrorOutput))
				return
			}

			uploadedChk, _, err := r.FormFile("chunk")
			if err != nil {
				renderError(w, err, []byte(internalErrorOutput))
				return
			}
			defer uploadedChk.Close()

			fileBytes, err := ioutil.ReadAll(uploadedChk)
			if err != nil {
				renderError(w, err, []byte(internalErrorOutput))
				return
			}

			// save the file:
			err = appIdentity.Sub().Storage().Save(bucketHashStr, fileBytes)
			if err != nil {
				renderError(w, err, []byte(internalErrorOutput))
				return
			}

			// success:
			renderSuccess(w, []byte(successPostOutput))
			return
		}

		err := errors.New(missingBucketHashErrorOutput)
		renderError(w, err, []byte(internalErrorOutput))
		return
	}

	str := fmt.Sprintf(authErrorOutput, token)
	err := errors.New(str)
	renderError(w, err, []byte(str))
}

func (app *application) deleteChunk(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get(shared.TokenHeadKeyname)
	defer app.deleteAuthApp(token)

	if appIdentity, ok := app.authApps[token]; ok {
		vars := mux.Vars(r)
		if bucketHashStr, ok := vars["bucket_hash"]; ok {
			if chunkHashStr, ok := vars["chunk_hash"]; ok {
				err := appIdentity.Sub().Storage().Delete(bucketHashStr, chunkHashStr)
				if err != nil {
					renderError(w, err, []byte(internalErrorOutput))
					return
				}

				renderSuccess(w, []byte(successPostOutput))
				return
			}

			err := errors.New(missingChunkHashErrorOutput)
			renderError(w, err, []byte(internalErrorOutput))
			return
		}

		err := errors.New(missingBucketHashErrorOutput)
		renderError(w, err, []byte(internalErrorOutput))
		return
	}

	str := fmt.Sprintf(authErrorOutput, token)
	err := errors.New(str)
	renderError(w, err, []byte(str))
}

func (app *application) deleteBucketChunks(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get(shared.TokenHeadKeyname)
	defer app.deleteAuthApp(token)

	if appIdentity, ok := app.authApps[token]; ok {
		vars := mux.Vars(r)
		if bucketHashStr, ok := vars["bucket_hash"]; ok {
			err := appIdentity.Sub().Storage().DeleteAll(bucketHashStr)
			if err != nil {
				renderError(w, err, []byte(internalErrorOutput))
				return
			}

			renderSuccess(w, []byte(successPostOutput))
			return
		}

		err := errors.New(missingBucketHashErrorOutput)
		renderError(w, err, []byte(internalErrorOutput))
		return
	}

	str := fmt.Sprintf(authErrorOutput, token)
	err := errors.New(str)
	renderError(w, err, []byte(str))
}

func (app *application) deleteAuthApp(token string) {
	if _, ok := app.authApps[token]; ok {
		delete(app.authApps, token)
	}
}
