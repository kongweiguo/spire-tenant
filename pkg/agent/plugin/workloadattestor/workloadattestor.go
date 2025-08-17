package workloadattestor

import (
	"context"

	"github.com/kongweiguo/spire-tenant/pkg/common/catalog"
	"github.com/kongweiguo/spire-tenant/proto/spire/common"
)

type WorkloadAttestor interface {
	catalog.PluginInfo

	Attest(ctx context.Context, pid int) ([]*common.Selector, error)
}
