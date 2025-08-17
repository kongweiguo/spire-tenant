package fakeserverkeymanager

import (
	"testing"

	keymanagerv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/server/keymanager/v1"
	"github.com/kongweiguo/spire-tenant/pkg/common/catalog"
	"github.com/kongweiguo/spire-tenant/pkg/server/plugin/keymanager"
	keymanagerbase "github.com/kongweiguo/spire-tenant/pkg/server/plugin/keymanager/base"
	"github.com/kongweiguo/spire-tenant/test/plugintest"
	"github.com/kongweiguo/spire-tenant/test/testkey"
)

func New(t *testing.T) keymanager.KeyManager {
	plugin := keyManager{
		Base: keymanagerbase.New(keymanagerbase.Config{
			Generator: &testkey.Generator{},
		}),
	}

	v1 := new(keymanager.V1)
	plugintest.Load(t, catalog.MakeBuiltIn("fake", keymanagerv1.KeyManagerPluginServer(plugin)), v1)
	return v1
}

type keyManager struct {
	*keymanagerbase.Base
}
