package client

import "github.com/kongweiguo/spire-tenant/proto/spire/common"

type Update struct {
	Entries map[string]*common.RegistrationEntry
	Bundles map[string]*common.Bundle
}
