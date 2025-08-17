package main

import (
	"os"

	"github.com/kongweiguo/spire-tenant/cmd/spire-server/cli"
	"github.com/kongweiguo/spire-tenant/pkg/common/entrypoint"
)

func main() {
	os.Exit(entrypoint.NewEntryPoint(new(cli.CLI).Run).Main())
}
