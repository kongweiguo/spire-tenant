package nodeattestor

import (
	"context"

	"github.com/kongweiguo/spire-tenant/pkg/common/catalog"
	"github.com/kongweiguo/spire-tenant/proto/spire/common"
)

type NodeAttestor interface {
	catalog.PluginInfo

	Attest(ctx context.Context, payload []byte, challengeFn func(ctx context.Context, challenge []byte) ([]byte, error)) (*AttestResult, error)
}

type AttestResult struct {
	AgentID     string
	Selectors   []*common.Selector
	CanReattest bool
}
