//revive:disable:package-comments
package service

import (
	"testing"

	"connectrpc.com/connect/v2"

	"buf.build/gen/go/authaas/identity-service/protocolbuffers/go/identity"
	identitytypes "buf.build/gen/go/authaas/identity/protocolbuffers/go/identity"
	"github.com/authaas/identity-data-bindings-connect-go/identity/data/dataconnect"
)

const principalID = "01234567-89ab-4def-8123-456789abcdef"

// dataStub is a data service nothing calls.
type dataStub struct {
	dataconnect.UnimplementedServiceHandler
}

func TestDelete(t *testing.T) {
	t.Run("is not served", func(t *testing.T) {
		server := New(&dataStub{})

		request := identity.DeleteRequest_builder{
			Principal: identitytypes.Principal_builder{Id: principalID}.Build(),
		}.Build()

		_, err := server.Delete(t.Context(), request)

		if got := connect.CodeOf(err); got != connect.CodeUnimplemented {
			t.Errorf("code = %v, want %v", got, connect.CodeUnimplemented)
		}
	})
}
