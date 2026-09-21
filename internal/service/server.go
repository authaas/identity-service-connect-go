//revive:disable:package-comments
package service

import (
	"log/slog"

	"git.sonicoriginal.software/logger"

	"github.com/authaas/identity-data-bindings-connect-go/identity/data/dataconnect"
	"github.com/authaas/identity-service-bindings-connect-go/identity/identityconnect"
)

// Server serves identity.Service over the identity data service.
type Server struct {
	identityconnect.UnimplementedServiceHandler

	log  *slog.Logger
	data dataconnect.ServiceClient
}

// New returns a Server over data.
func New(log *slog.Logger, data dataconnect.ServiceClient) *Server {
	if log == nil {
		log = logger.NewNullLogger()
	}

	return &Server{log: log, data: data}
}
