package memory_test

import (
	"testing"

	"github.com/kongweiguo/spire-tenant/pkg/agent/plugin/keymanager"
	"github.com/kongweiguo/spire-tenant/pkg/agent/plugin/keymanager/memory"
	keymanagertest "github.com/kongweiguo/spire-tenant/pkg/agent/plugin/keymanager/test"
	"github.com/kongweiguo/spire-tenant/test/plugintest"
)

func TestKeyManagerContract(t *testing.T) {
	keymanagertest.Test(t, keymanagertest.Config{
		Create: func(t *testing.T) keymanager.KeyManager {
			km := new(keymanager.V1)
			plugintest.Load(t, memory.TestBuiltIn(keymanagertest.NewGenerator()), km)
			return km
		},
	})
}
