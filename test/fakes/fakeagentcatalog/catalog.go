package fakeagentcatalog

import (
	"github.com/kongweiguo/spire-tenant/pkg/agent/plugin/keymanager"
	"github.com/kongweiguo/spire-tenant/pkg/agent/plugin/nodeattestor"
	"github.com/kongweiguo/spire-tenant/pkg/agent/plugin/svidstore"
	"github.com/kongweiguo/spire-tenant/pkg/agent/plugin/workloadattestor"
)

func New() *Catalog {
	return new(Catalog)
}

type Catalog struct {
	keyManagerRepository
	nodeAttestorRepository
	svidStoreRepository
	workloadAttestorRepository
}

// We need distinct type names to embed in the Catalog above, since the types
// we want to actually embed are all named the same.
type keyManagerRepository struct{ keymanager.Repository }
type nodeAttestorRepository struct{ nodeattestor.Repository }
type svidStoreRepository struct{ svidstore.Repository }
type workloadAttestorRepository struct{ workloadattestor.Repository }
