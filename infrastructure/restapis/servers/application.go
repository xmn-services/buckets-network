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
	"github.com/xmn-services/buckets-network/application/servers/authenticates"
	init_chains "github.com/xmn-services/buckets-network/application/servers/chains"
	"github.com/xmn-services/buckets-network/domain/memory/file/contents/content"
	"github.com/xmn-services/buckets-network/domain/memory/peers/peer"
)

type application struct {
	cmdApp                commands.Application
	initChainAdapter      init_chains.Adapter
	authenticateAdapter   authenticates.Adapter
	updateIdentityAdapter identities_app.UpdateAdapter
	peerAdapter           peer.Adapter
	contentBuilder        content.Builder
	router                *mux.Router
	server                *http.Server
	authApps              map[string]identities_app.Application
	maxUploadFileSize     int64
	waitPeriod            time.Duration
	port                  uint
}

func createApplication(
	cmdApp commands.Application,
	initChainAdapter init_chains.Adapter,
	authenticateAdapter authenticates.Adapter,
	updateIdentityAdapter identities_app.UpdateAdapter,
	peerAdapter peer.Adapter,
	contentBuilder content.Builder,
	router *mux.Router,
	maxUploadFileSize int64,
	waitPeriod time.Duration,
	port uint,
) servers.Application {
	out := application{
		cmdApp:                cmdApp,
		initChainAdapter:      initChainAdapter,
		authenticateAdapter:   authenticateAdapter,
		updateIdentityAdapter: updateIdentityAdapter,
		peerAdapter:           peerAdapter,
		contentBuilder:        contentBuilder,
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
	out.router.HandleFunc("/storages/{hash:[0-9a-f]+}", out.retrieveStoredFileByHash).Methods(http.MethodGet, http.MethodOptions)

	// middleware:
	out.router.Use(mux.CORSMethodMiddleware(out.router))

	// identities:
	identityRouter := out.router.PathPrefix("/identities").Subrouter()
	identityRouter.HandleFunc("/", out.retrieveIdentity).Methods(http.MethodGet, http.MethodOptions)
	identityRouter.HandleFunc("/", out.updateIdentity).Methods(http.MethodPut, http.MethodOptions)
	identityRouter.HandleFunc("/", out.deleteIdentity).Methods(http.MethodDelete, http.MethodOptions)

	identityRouter.HandleFunc("/miners/test/{difficulty:[0-9]+}", out.identityMinerTest).Methods(http.MethodGet, http.MethodOptions)
	identityRouter.HandleFunc("/miners/block/{hash:[0-9a-f]+}", out.identityMinerBlock).Methods(http.MethodGet, http.MethodOptions)
	identityRouter.HandleFunc("/miners/link/{hash:[0-9a-f]+}", out.identityMinerLink).Methods(http.MethodGet, http.MethodOptions)

	identityRouter.HandleFunc("/chains", out.initIdentityChain).Methods(http.MethodPost, http.MethodOptions)
	identityRouter.HandleFunc("/chains/blocks/{additional:[0-9]+}", out.identityChainMineBlocks).Methods(http.MethodPost, http.MethodOptions)
	identityRouter.HandleFunc("/chains/links/{additional:[0-9]+}", out.identityChainMineLinks).Methods(http.MethodPost, http.MethodOptions)

	identityRouter.HandleFunc("/buckets", out.retrieveIdentityBuckets).Methods(http.MethodGet, http.MethodOptions)
	identityRouter.HandleFunc("/buckets", out.saveIdentityBucketPath).Methods(http.MethodPost, http.MethodOptions)
	identityRouter.HandleFunc("/buckets/{hash:[0-9a-f]+}", out.retrieveIdentityBucketByHash).Methods(http.MethodGet, http.MethodOptions)
	identityRouter.HandleFunc("/buckets/{hash:[0-9a-f]+}", out.deleteIdentityBucketByHash).Methods(http.MethodDelete, http.MethodOptions)

	identityRouter.HandleFunc("/storages/{hash:[0-9a-f]+}", out.saveStoredFile).Methods(http.MethodPost, http.MethodOptions)
	identityRouter.HandleFunc("/storages/{hash:[0-9a-f]+}", out.deleteStoredFileByHash).Methods(http.MethodDelete, http.MethodOptions)

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
	go func() {
		app.Stop()
	}()

	return nil
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
		token := r.Header.Get(tokenHeadKeyname)
		auth, err := app.authenticateAdapter.Base64ToAuthenticate(token)
		if err != nil {
			renderError(w, err, []byte(internalErrorOutput))
			return
		}

		authApp, err := app.cmdApp.Current().Authenticate(
			auth.Name(),
			auth.Seed(),
			auth.Password(),
		)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("invalid authentication"))
			return
		}

		app.authApps[token] = authApp

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
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
	}

	err := errors.New(missingIndexErrorOutput)
	renderError(w, err, []byte(internalErrorOutput))

}

func (app *application) identityMinerTest(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get(tokenHeadKeyname)
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
	token := r.Header.Get(tokenHeadKeyname)
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
	token := r.Header.Get(tokenHeadKeyname)
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
	token := r.Header.Get(tokenHeadKeyname)
	defer app.deleteAuthApp(token)

	if appIdentity, ok := app.authApps[token]; ok {
		err := r.ParseForm()
		if err != nil {
			renderError(w, err, []byte(internalErrorOutput))
			return
		}

		initChain, err := app.initChainAdapter.URLValuesToChain(r.Form)
		if err != nil {
			renderError(w, err, []byte(internalErrorOutput))
			return
		}

		err = appIdentity.Sub().Chain().Init(
			initChain.MiningValue(),
			initChain.BaseDifficulty(),
			initChain.IncreasePerBucket(),
			initChain.LinkDifficulty(),
			initChain.RootAdditionalBuckets(),
			initChain.HeadAdditionalBuckets(),
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

		token := r.Header.Get(tokenHeadKeyname)
		defer app.deleteAuthApp(token)

		if appIdentity, ok := app.authApps[token]; ok {
			err := r.ParseForm()
			if err != nil {
				renderError(w, err, []byte(internalErrorOutput))
				return
			}

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

		token := r.Header.Get(tokenHeadKeyname)
		defer app.deleteAuthApp(token)

		if appIdentity, ok := app.authApps[token]; ok {
			err := r.ParseForm()
			if err != nil {
				renderError(w, err, []byte(internalErrorOutput))
				return
			}

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

func (app *application) retrieveStoredFileByHash(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if hashStr, ok := vars["hash"]; ok {
		storedFile, err := app.cmdApp.Sub().Storage().Retrieve(hashStr)
		if err != nil {
			renderError(w, err, []byte(internalErrorOutput))
			return
		}

		js, err := json.Marshal(storedFile)
		if err != nil {
			renderError(w, err, []byte(internalErrorOutput))
			return
		}

		renderSuccess(w, js)
	}

	err := errors.New(missingHashErrorOutput)
	renderError(w, err, []byte(internalErrorOutput))

}

func (app *application) retrieveIdentity(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get(tokenHeadKeyname)
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
	token := r.Header.Get(tokenHeadKeyname)
	defer app.deleteAuthApp(token)

	if appIdentity, ok := app.authApps[token]; ok {
		err := r.ParseForm()
		if err != nil {
			renderError(w, err, []byte(internalErrorOutput))
			return
		}

		updateIdentity, err := app.updateIdentityAdapter.URLValuesToUpdate(r.Form)
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
	token := r.Header.Get(tokenHeadKeyname)
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

func (app *application) retrieveIdentityBuckets(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get(tokenHeadKeyname)
	defer app.deleteAuthApp(token)

	if appIdentity, ok := app.authApps[token]; ok {
		buckets, err := appIdentity.Sub().Bucket().RetrieveAll()
		if err != nil {
			renderError(w, err, []byte(internalErrorOutput))
			return
		}

		js, err := json.Marshal(buckets)
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

func (app *application) saveIdentityBucketPath(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get(tokenHeadKeyname)
	defer app.deleteAuthApp(token)

	if appIdentity, ok := app.authApps[token]; ok {
		path := r.Form.Get(pathKeyname)
		err := appIdentity.Sub().Bucket().Add(path)
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

func (app *application) retrieveIdentityBucketByHash(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get(tokenHeadKeyname)
	defer app.deleteAuthApp(token)

	if appIdentity, ok := app.authApps[token]; ok {
		vars := mux.Vars(r)
		if hashStr, ok := vars["hash"]; ok {
			bucket, err := appIdentity.Sub().Bucket().Retrieve(hashStr)
			if err != nil {
				renderError(w, err, []byte(internalErrorOutput))
				return
			}

			js, err := json.Marshal(bucket)
			if err != nil {
				renderError(w, err, []byte(internalErrorOutput))
				return
			}

			renderSuccess(w, js)
			return
		}

		err := errors.New(missingHashErrorOutput)
		renderError(w, err, []byte(internalErrorOutput))
	}

	str := fmt.Sprintf(authErrorOutput, token)
	err := errors.New(str)
	renderError(w, err, []byte(str))
}

func (app *application) deleteIdentityBucketByHash(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get(tokenHeadKeyname)
	defer app.deleteAuthApp(token)

	if appIdentity, ok := app.authApps[token]; ok {
		vars := mux.Vars(r)
		if hashStr, ok := vars["hash"]; ok {
			err := appIdentity.Sub().Bucket().Delete(hashStr)
			if err != nil {
				renderError(w, err, []byte(internalErrorOutput))
				return
			}

			renderSuccess(w, []byte(successPostOutput))
			return
		}

		err := errors.New(missingHashErrorOutput)
		renderError(w, err, []byte(internalErrorOutput))
	}

	str := fmt.Sprintf(authErrorOutput, token)
	err := errors.New(str)
	renderError(w, err, []byte(str))
}

func (app *application) saveStoredFile(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get(tokenHeadKeyname)
	defer app.deleteAuthApp(token)

	if appIdentity, ok := app.authApps[token]; ok {
		vars := mux.Vars(r)
		if fileHashStr, ok := vars["hash"]; ok {
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

			// retrieve the file:
			file, err := app.cmdApp.Sub().Storage().Retrieve(fileHashStr)
			if err != nil {
				renderError(w, err, []byte(internalErrorOutput))
				return
			}

			// build the content:
			content, err := app.contentBuilder.Create().WithContent(fileBytes).Now()
			if err != nil {
				renderError(w, err, []byte(internalErrorOutput))
				return
			}

			// add the content to the file:
			err = file.Contents().Add(content)
			if err != nil {
				renderError(w, err, []byte(internalErrorOutput))
				return
			}

			// save the file:
			err = appIdentity.Sub().Storage().Save(file)
			if err != nil {
				renderError(w, err, []byte(internalErrorOutput))
				return
			}

			// success:
			renderSuccess(w, []byte(successPostOutput))
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

func (app *application) deleteStoredFileByHash(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get(tokenHeadKeyname)
	defer app.deleteAuthApp(token)

	if appIdentity, ok := app.authApps[token]; ok {
		vars := mux.Vars(r)
		if hashStr, ok := vars["hash"]; ok {
			err := appIdentity.Sub().Storage().Delete(hashStr)
			if err != nil {
				renderError(w, err, []byte(internalErrorOutput))
				return
			}

			renderSuccess(w, []byte(successPostOutput))
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

func (app *application) deleteAuthApp(token string) {
	if _, ok := app.authApps[token]; ok {
		delete(app.authApps, token)
	}
}
