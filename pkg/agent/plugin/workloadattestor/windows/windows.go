package windows

import "github.com/kongweiguo/spire-tenant/pkg/common/catalog"

const (
	pluginName = "windows"
)

func BuiltIn() catalog.BuiltIn {
	return builtin(New())
}
