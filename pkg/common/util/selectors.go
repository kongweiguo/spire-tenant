package util

import "github.com/kongweiguo/spire-tenant/proto/spire/common"

func EqualsSelectors(a, b []*common.Selector) bool {
	selectorsA := a
	SortSelectors(selectorsA)

	selectorsB := b
	SortSelectors(selectorsB)

	return compareSelectors(selectorsA, selectorsB) == 0
}
