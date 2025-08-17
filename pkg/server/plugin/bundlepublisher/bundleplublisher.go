package bundlepublisher

import (
	"context"

	"github.com/kongweiguo/spire-tenant/pkg/common/catalog"
	"github.com/kongweiguo/spire-tenant/proto/spire/common"
)

type BundlePublisher interface {
	catalog.PluginInfo

	PublishBundle(ctx context.Context, bundle *common.Bundle) error
}
