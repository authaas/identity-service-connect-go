// The identity service: identity.Service served with connect-go over the
// identity data service.
package main

import (
	"os"

	"github.com/authaas/identity-service-connect-go/internal/cli"
)

func main() {
	os.Exit(cli.Run())
}
