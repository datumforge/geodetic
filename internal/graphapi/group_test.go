package graphapi_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	ent "github.com/datumforge/geodetic/internal/ent/generated"
)

func (suite *GraphTestSuite) TestQueryGroup() {
	t := suite.T()

	group := (&GroupBuilder{client: suite.client}).MustNew(context.Background(), t)

	testCases := []struct {
		name     string
		query    string
		expected *ent.Group
		errorMsg string
	}{
		{
			name:     "happy path group",
			query:    group.Name,
			expected: group,
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			resp, err := suite.client.geodetic.GetGroup(context.Background(), tc.query)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.Group)
		})
	}
}
