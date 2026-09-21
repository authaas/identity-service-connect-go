//revive:disable:package-comments
package service

import (
	"github.com/authaas/identity-data-bindings-connect-go/identity/data/dataconnect"
	"github.com/authaas/identity-service-bindings-connect-go/identity/identityconnect"
)

// Server serves identity.Service over the identity data service.
type Server struct {
	identityconnect.UnimplementedServiceHandler

	data dataconnect.ServiceClient
}

// New returns a Server over data.
func New(data dataconnect.ServiceClient) *Server {
	return &Server{data: data}
}
