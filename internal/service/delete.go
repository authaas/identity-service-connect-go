//revive:disable:package-comments
package service

import (
	"context"

	"connectrpc.com/connect/v2"

	"buf.build/gen/go/authaas/identity-service/protocolbuffers/go/identity"
)

// Delete is not served. Removing an identity belongs to a flow that
// establishes who may remove it, and no such flow exists.
func (*Server) Delete(context.Context, *identity.DeleteRequest) (*identity.DeleteResponse, error) {
	return nil, connect.NewError(connect.CodeUnimplemented, "identity removal is not served")
}
