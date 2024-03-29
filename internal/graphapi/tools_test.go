package graphapi_test

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/Yamashou/gqlgenc/clientv2"
	"github.com/datumforge/datum/pkg/testutils"
	"github.com/datumforge/go-turso"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"

	ent "github.com/datumforge/geodetic/internal/ent/generated"
	"github.com/datumforge/geodetic/internal/entdb"
	"github.com/datumforge/geodetic/internal/graphapi"
	"github.com/datumforge/geodetic/pkg/geodeticclient"
)

// TestGraphTestSuite runs all the tests in the GraphTestSuite
func TestGraphTestSuite(t *testing.T) {
	suite.Run(t, new(GraphTestSuite))
}

// GraphTestSuite handles the setup and teardown between tests
type GraphTestSuite struct {
	suite.Suite
	client *client
	tc     *testutils.TC
}

// client contains all the clients the test need to interact with
type client struct {
	db       *ent.Client
	geodetic geodeticclient.GeodeticClient
}

type graphClient struct {
	srvURL     string
	httpClient *http.Client
}

func (suite *GraphTestSuite) SetupSuite() {
	ctx := context.Background()

	suite.tc = entdb.NewTestContainer(ctx)
}

func (suite *GraphTestSuite) SetupTest() {
	t := suite.T()

	ctx := context.Background()

	// setup logger
	logger := zap.NewNop().Sugar()

	// setup mock turso client
	tc := turso.NewMockClient()

	opts := []ent.Option{
		ent.Logger(*logger),
		ent.Turso(tc),
	}

	// create database connection
	db, err := entdb.NewTestClient(ctx, suite.tc, opts)
	if err != nil {
		require.NoError(t, err, "failed opening connection to database")
	}

	// assign values
	c := &client{
		db:       db,
		geodetic: graphTestClient(t, db),
	}

	suite.client = c
}

func (suite *GraphTestSuite) TearDownTest() {
	if err := suite.client.db.Close(); err != nil {
		log.Fatalf("failed to close database: %s", err)
	}
}

func (suite *GraphTestSuite) TearDownSuite() {
	if suite.tc.Container != nil {
		if err := suite.tc.Container.Terminate(context.Background()); err != nil {
			log.Fatalf("failed to terminate container: %s", err)
		}
	}
}

func graphTestClient(t *testing.T, c *ent.Client) geodeticclient.GeodeticClient {
	logger := zaptest.NewLogger(t, zaptest.Level(zap.ErrorLevel)).Sugar()

	srv := handler.NewDefaultServer(
		graphapi.NewExecutableSchema(
			graphapi.Config{Resolvers: graphapi.NewResolver(c).WithLogger(logger)},
		))

	graphapi.WithTransactions(srv, c)

	g := &graphClient{
		srvURL:     "query",
		httpClient: &http.Client{Transport: localRoundTripper{handler: srv}},
	}

	// set options
	opt := &clientv2.Options{
		ParseDataAlongWithErrors: false,
	}

	// setup interceptors
	i := geodeticclient.WithEmptyInterceptor()

	return geodeticclient.NewClient(g.httpClient, g.srvURL, opt, i)
}

// localRoundTripper is an http.RoundTripper that executes HTTP transactions
// by using handler directly, instead of going over an HTTP connection.
type localRoundTripper struct {
	handler http.Handler
}

func (l localRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	w := httptest.NewRecorder()
	l.handler.ServeHTTP(w, req)

	return w.Result(), nil
}
