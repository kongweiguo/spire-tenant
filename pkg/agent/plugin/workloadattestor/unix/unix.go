package unix

import "github.com/kongweiguo/spire-tenant/pkg/common/catalog"

const (
	pluginName = "unix"
)

func BuiltIn() catalog.BuiltIn {
	return builtin(New())
}
