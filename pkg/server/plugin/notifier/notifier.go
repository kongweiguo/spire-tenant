package notifier

import (
	"context"

	"github.com/kongweiguo/spire-tenant/pkg/common/catalog"
	"github.com/kongweiguo/spire-tenant/proto/spire/common"
)

type Notifier interface {
	catalog.PluginInfo

	NotifyAndAdviseBundleLoaded(ctx context.Context, bundle *common.Bundle) error
	NotifyBundleUpdated(ctx context.Context, bundle *common.Bundle) error
}
