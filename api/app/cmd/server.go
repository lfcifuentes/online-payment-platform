// Package cmd provides the command-line interface for the server application.
// It initializes the server and handles graceful shutdown.
package cmd

// Import packages
import (
	"context"
	"errors"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lfcifuentes/online-payment-platform/api/app/infrastructure/http/router"
	"github.com/lfcifuentes/online-payment-platform/api/app/internal/adapters/pgsql"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(runServer)
}

var runServer = &cobra.Command{
	Use:   "server",
	Short: "Run app server",
	Run:   runServerHandler,
}

const (
	// ServerTearDownTimeOut sets the timout for gracefully shutdown the server.
	ServerTearDownTimeOut = 10 * time.Second
)

// runServerHandler runs the server
func runServerHandler(_ *cobra.Command, _ []string) {
	viper.SetDefault("HTTP_PORT", "8001")
	viper.SetDefault("APP_VERSION", "1.0.0")
	// Create a new mongodb adapter
	slog.Info("Creating database adapter", "event", "database_adapter_creation")
	db, err := pgsql.NewDBAdapter()
	if err != nil {
		slog.Error(err.Error(), "event", "database_connection_failed")
		return
	}
	err = db.Ping()
	if err != nil {
		slog.Error(err.Error(), "event", "database_pin_connection_failed")
		return
	}

	slog.Info("Creating server", "event", "server_creation")
	// Configura el enrutador
	routes := router.NewRouter(db)

	slog.Info("Server created", "event", "server_created")
	port := viper.GetString("HTTP_PORT")
	server := http.Server{
		Addr:    "0.0.0.0:" + port,
		Handler: routes,
	}

	err = listenWithGracefulShutdown(ServerTearDownTimeOut, &server, db)
	if err != nil {
		slog.Error(err.Error(), "event", "server_listeners_failed")
	}
}

// listenWithGracefulShutdown listens for incoming connections and gracefully shutdown the server
func listenWithGracefulShutdown(timeout time.Duration, srv *http.Server, db *pgsql.DBAdapter) error {
	shutdownCompleted := make(chan struct{})

	go func() {
		defer close(shutdownCompleted)

		osInterruptSignal := make(chan os.Signal, 1)
		signal.Notify(osInterruptSignal, syscall.SIGTERM, syscall.SIGINT)

		<-osInterruptSignal

		ctxTimeout, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		err := srv.Shutdown(ctxTimeout)
		if err != nil {
			log.Fatal("server_listeners_shutdown_failed")
		}
	}()

	if err := srv.ListenAndServe(); errors.Is(err, http.ErrServerClosed) {
		return err
	}

	// ListenAndServe returns immediately after Shutdown is invoked (see doc)
	// This means we have to wait right after it returns to give enough time to the Shutdown method
	// to wait for all connections to close down gracefully
	<-shutdownCompleted

	err := db.Close()
	if err != nil {
		slog.Error(err.Error(), "event", "database_connection_close_failed")
	}

	slog.Info("Server shutdown completed", "event", "server_shutdown_completed")

	return nil
}
