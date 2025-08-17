package systemd

import "github.com/kongweiguo/spire-tenant/pkg/common/catalog"

const (
	pluginName = "systemd"
)

func BuiltIn() catalog.BuiltIn {
	return builtin(New())
}
