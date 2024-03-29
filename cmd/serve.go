package cmd

import (
	"context"

	"github.com/datumforge/datum/pkg/otelx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/datumforge/datum/pkg/cache"
	"github.com/datumforge/go-turso"

	ent "github.com/datumforge/geodetic/internal/ent/generated"
	"github.com/datumforge/geodetic/internal/entdb"
	"github.com/datumforge/geodetic/internal/httpserve/config"
	"github.com/datumforge/geodetic/internal/httpserve/server"
	"github.com/datumforge/geodetic/internal/httpserve/serveropts"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start the server",
	RunE: func(cmd *cobra.Command, args []string) error {
		return serve(cmd.Context())
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.PersistentFlags().String("config", "./config/.config.yaml", "config file location")
	viperBindFlag("config", serveCmd.PersistentFlags().Lookup("config"))
}

func serve(ctx context.Context) error {
	// setup db connection for server
	var (
		err error
	)

	serverOpts := []serveropts.ServerOption{}
	serverOpts = append(serverOpts,
		serveropts.WithConfigProvider(&config.ConfigProviderWithRefresh{}),
		serveropts.WithLogger(logger),
		serveropts.WithHTTPS(),
		serveropts.WithMiddleware(),
	)

	so := serveropts.NewServerOptions(serverOpts, viper.GetString("config"))

	err = otelx.NewTracer(so.Config.Settings.Tracer, appName, logger)
	if err != nil {
		logger.Fatalw("failed to initialize tracer", "error", err)
	}

	tursoClient, err := turso.NewClient(so.Config.Settings.Turso)
	if err != nil {
		logger.Fatalw("failed to initialize turso client", "error", err)
	}

	// create ent dependency injection
	entOpts := []ent.Option{ent.Logger(*logger), ent.Turso(tursoClient)}

	// Setup DB connection
	entdbClient, dbConfig, err := entdb.NewMultiDriverDBClient(ctx, so.Config.Settings.DB, logger, entOpts)
	if err != nil {
		return err
	}

	defer entdbClient.Close()

	// Setup Redis connection
	redisClient := cache.New(so.Config.Settings.Redis)
	defer redisClient.Close()

	// Add Driver to the Handlers Config
	so.Config.Handler.DBClient = entdbClient

	// Add redis client to Handlers Config
	so.Config.Handler.RedisClient = redisClient

	// add ready checks
	so.AddServerOptions(
		serveropts.WithReadyChecks(dbConfig, redisClient),
	)

	srv := server.NewServer(so.Config, so.Config.Logger)

	// Setup Graph API Handlers
	so.AddServerOptions(serveropts.WithGraphRoute(srv, entdbClient))

	if err := srv.StartEchoServer(ctx); err != nil {
		logger.Error("failed to run server", zap.Error(err))
	}

	return nil
}
