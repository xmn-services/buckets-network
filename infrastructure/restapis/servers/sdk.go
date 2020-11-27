package servers

import (
	"time"

	"github.com/gorilla/mux"
	"github.com/xmn-services/buckets-network/application/commands"
	identities_app "github.com/xmn-services/buckets-network/application/commands/identities"
	"github.com/xmn-services/buckets-network/application/servers"
	"github.com/xmn-services/buckets-network/domain/memory/file/contents/content"
	"github.com/xmn-services/buckets-network/domain/memory/peers/peer"
)

const successPostOutput = "success"

const internalErrorOutput = "server error"

const authErrorOutput = "then authentication token (%s) is invalid"

const missingHashErrorOutput = "the hash was expected in the vars"

const missingIndexErrorOutput = "the index was expected in the vars"

// NewApplication creates a new restful server application
func NewApplication(
	cmdApp commands.Application,
	router *mux.Router,
	maxUploadFileSize int64,
	waitPeriod time.Duration,
	port uint,
) servers.Application {
	updateIdentityBuilder := identities_app.NewUpdateBuilder()
	peerAdapter := peer.NewAdapter()
	contentBuilder := content.NewBuilder()
	return createApplication(
		cmdApp,
		updateIdentityBuilder,
		peerAdapter,
		contentBuilder,
		router,
		maxUploadFileSize,
		waitPeriod,
		port,
	)
}
