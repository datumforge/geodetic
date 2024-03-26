package graphapi_test

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/datumforge/geodetic/internal/ent/enums"
	ent "github.com/datumforge/geodetic/internal/ent/generated"
	"github.com/datumforge/geodetic/internal/geodeticclient"
)

func (suite *GraphTestSuite) TestQueryDatabase() {
	t := suite.T()

	db := (&DatabaseBuilder{client: suite.client}).MustNew(context.Background(), t)

	testCases := []struct {
		name     string
		query    string
		expected *ent.Database
		errorMsg string
	}{
		{
			name:     "happy path database",
			query:    db.Name,
			expected: db,
		},
		{
			name:     "database not found",
			query:    "notfound",
			expected: nil,
			errorMsg: "database not found",
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			resp, err := suite.client.geodetic.GetDatabase(context.Background(), tc.query)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.Database)
		})
	}
}

func (suite *GraphTestSuite) TestListDatabases() {
	t := suite.T()

	_ = (&DatabaseBuilder{client: suite.client}).MustNew(context.Background(), t)
	_ = (&DatabaseBuilder{client: suite.client}).MustNew(context.Background(), t)

	t.Run("List Databases", func(t *testing.T) {
		resp, err := suite.client.geodetic.GetAllDatabases(context.Background())

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, resp.Databases)
		require.Len(t, resp.Databases.Edges, 2)
	})
}

func (suite *GraphTestSuite) TestCreateDatabase() {
	t := suite.T()

	group := (&GroupBuilder{client: suite.client}).MustNew(context.Background(), t)

	testCases := []struct {
		name     string
		orgID    string
		groupID  string
		provider *enums.DatabaseProvider
		errorMsg string
	}{
		{
			name:     "happy path, turso database",
			orgID:    "01HSCAGDJ1XZ12Y06FESH4VEC1",
			groupID:  group.ID,
			provider: &enums.Turso,
		},
		{
			name:     "happy path, local database",
			orgID:    "01HSCAGDJ1XZ12Y06FESH4VEC2",
			groupID:  group.ID,
			provider: &enums.Local,
		},
		{
			name:     "duplicate org id",
			orgID:    "01HSCAGDJ1XZ12Y06FESH4VEC2",
			groupID:  group.ID,
			provider: &enums.Local,
			errorMsg: "constraint failed",
		},
		{
			name:     "missing group",
			orgID:    "01HSCAGDJ1XZ12Y06FESH4VEC3",
			groupID:  "notfound",
			provider: &enums.Turso,
			errorMsg: "group not found",
		},
		{
			name:     "missing org id",
			orgID:    "",
			groupID:  group.ID,
			provider: &enums.Turso,
			errorMsg: "invalid or unparsable field: organization_id",
		},
	}

	for _, tc := range testCases {
		t.Run("Create "+tc.name, func(t *testing.T) {
			g := geodeticclient.CreateDatabaseInput{
				OrganizationID: tc.orgID,
				Provider:       tc.provider,
				GroupID:        tc.groupID,
			}

			resp, err := suite.client.geodetic.CreateDatabase(context.Background(), g)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.CreateDatabase)

			assert.Contains(t, resp.CreateDatabase.Database.Name, strings.ToLower(tc.orgID))
			assert.Equal(t, *tc.provider, resp.CreateDatabase.Database.Provider)
			assert.Equal(t, tc.orgID, resp.CreateDatabase.Database.OrganizationID)
		})
	}
}

func (suite *GraphTestSuite) TestDeleteDatabase() {
	t := suite.T()

	db := (&DatabaseBuilder{client: suite.client}).MustNew(context.Background(), t)

	testCases := []struct {
		name     string
		dbName   string
		errorMsg string
	}{
		{
			name:   "happy path database",
			dbName: db.Name,
		},
		{
			name:     "db does not exist",
			dbName:   "lost-ark",
			errorMsg: "database not found",
		},
	}

	for _, tc := range testCases {
		t.Run("Delete "+tc.name, func(t *testing.T) {
			resp, err := suite.client.geodetic.DeleteDatabase(context.Background(), tc.dbName)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.DeleteDatabase)

			assert.NotEmpty(t, resp.DeleteDatabase.DeletedID)
		})
	}
}
