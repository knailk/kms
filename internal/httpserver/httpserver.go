// Package httpserver provides a preconfigured HTTP server.
package httpserver

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"kms/app/api/routes"
	"kms/app/config"
	"kms/app/registry"
	"kms/internal/shutdown"
	"kms/pkg/logger"

	"github.com/sirupsen/logrus"
)

type task struct {
	srv *http.Server
}

// Init initializes a new Server and registers
// routes to the given router
func Init(ctx context.Context, cfg *config.Config, tasks *shutdown.Tasks, provider *registry.Provider) {
	if tasks.HasStopSignal() {
		return
	}

	engine := routes.NewGinRouter(ctx, cfg, tasks, provider)

	t := &task{
		srv: &http.Server{
			Addr:              fmt.Sprintf(":%d", cfg.HTTPServer.Port),
			Handler:           engine,
			ReadHeaderTimeout: 10 * time.Second,
		},
	}

	go func() {
		if err := t.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.ErrorF(err.Error(), logrus.Fields{"Message": "HTTP server ListenAndServe failed"})
			tasks.GetSigChan() <- os.Interrupt
		}
	}()

	// Add shutdown task
	tasks.Add(t)
}

func (t *task) Name() string {
	return "httpserver"
}

func (t *task) Shutdown(ctx context.Context) error {
	if err := t.srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown http server: %w", err)
	}
	return nil
}
