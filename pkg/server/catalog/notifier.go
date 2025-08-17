package catalog

import (
	"github.com/kongweiguo/spire-tenant/pkg/common/catalog"
	"github.com/kongweiguo/spire-tenant/pkg/server/plugin/notifier"
	"github.com/kongweiguo/spire-tenant/pkg/server/plugin/notifier/gcsbundle"
	"github.com/kongweiguo/spire-tenant/pkg/server/plugin/notifier/k8sbundle"
)

type notifierRepository struct {
	notifier.Repository
}

func (repo *notifierRepository) Binder() any {
	return repo.AddNotifier
}

func (repo *notifierRepository) Constraints() catalog.Constraints {
	return catalog.ZeroOrMore()
}

func (repo *notifierRepository) Versions() []catalog.Version {
	return []catalog.Version{
		notifierV1{},
	}
}

func (repo *notifierRepository) BuiltIns() []catalog.BuiltIn {
	return []catalog.BuiltIn{
		gcsbundle.BuiltIn(),
		k8sbundle.BuiltIn(),
	}
}

type notifierV1 struct{}

func (notifierV1) New() catalog.Facade { return new(notifier.V1) }
func (notifierV1) Deprecated() bool    { return false }
