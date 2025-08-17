//go:build windows

package clitest

import (
	"net"

	"github.com/kongweiguo/spire-tenant/pkg/common/namedpipe"
)

func GetAddr(addr net.Addr) string {
	return namedpipe.GetPipeName(addr.String())
}
