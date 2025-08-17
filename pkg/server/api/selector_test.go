package api_test

import (
	"testing"

	"github.com/kongweiguo/spire-api-sdk-tenant/proto/spire/api/types"
	"github.com/kongweiguo/spire-tenant/pkg/server/api"
	"github.com/kongweiguo/spire-tenant/proto/spire/common"
	"github.com/stretchr/testify/require"
)

func TestSelectorsFromProto(t *testing.T) {
	testCases := []struct {
		name     string
		proto    []*types.Selector
		expected []*common.Selector
		err      string
	}{
		{
			name: "happy path",
			proto: []*types.Selector{
				{Type: "unix", Value: "uid:1000"},
				{Type: "unix", Value: "gid:1000"},
			},
			expected: []*common.Selector{
				{Type: "unix", Value: "uid:1000"},
				{Type: "unix", Value: "gid:1000"},
			},
		},
		{
			name:     "nil input",
			proto:    nil,
			expected: nil,
		},
		{
			name:     "empty slice",
			proto:    []*types.Selector{},
			expected: nil,
		},
		{
			name: "missing type",
			proto: []*types.Selector{
				{Type: "unix", Value: "uid:1000"},
				{Type: "", Value: "gid:1000"},
			},
			expected: nil,
			err:      "missing selector type",
		},
		{
			name: "missing value",
			proto: []*types.Selector{
				{Type: "unix", Value: ""},
				{Type: "unix", Value: "gid:1000"},
			},
			expected: nil,
			err:      "missing selector value",
		},
		{
			name: "type contains ':'",
			proto: []*types.Selector{
				{Type: "unix:uid", Value: "1000"},
				{Type: "unix", Value: "gid:1000"},
			},
			expected: nil,
			err:      "selector type contains ':'",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			selectors, err := api.SelectorsFromProto(testCase.proto)
			if testCase.err != "" {
				require.EqualError(t, err, testCase.err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, testCase.expected, selectors)

			// assert that a conversion in the opposite direction yields the
			// original types slice. In the special case that the input slice
			// is non-nil but empty, SelectorsFromProto returns nil so we
			// need to adjust the expected type accordingly.
			expected := testCase.proto
			if len(testCase.proto) == 0 {
				expected = nil
			}
			require.Equal(t, expected, api.ProtoFromSelectors(selectors))
		})
	}
}
