//revive:disable:package-comments
package main

import (
	"os"

	"github.com/authaas/identity-service-connect-go/internal/cli"
)

func main() {
	os.Exit(cli.Run())
}
