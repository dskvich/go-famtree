package rest

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

type server struct {
	cfg        Config
	httpServer *http.Server
}

type Config struct {
	ListenAddr string
	Router     http.Handler
	Logger     *logrus.Entry
}

func (cfg *Config) validate() error {
	var err error
	if cfg.ListenAddr == "" {
		err = multierror.Append(err, xerrors.Errorf("listen address has not been specified"))
	}
	if cfg.Router == nil {
		err = multierror.Append(err, xerrors.Errorf("router has not been specified"))
	}
	if cfg.Logger == nil {
		cfg.Logger = logrus.NewEntry(&logrus.Logger{Out: ioutil.Discard})
	}
	return err
}

func NewServer(cfg Config) *server {
	return &server{cfg: cfg}
}

func (s *server) Run(ctx context.Context) error {
	s.httpServer = &http.Server{
		Addr:              s.cfg.ListenAddr,
		Handler:           cors.Default().Handler(s.cfg.Router),
		ReadHeaderTimeout: time.Minute,
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(sigCh)

	errCh := make(chan error, 1)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		s.cfg.Logger.WithField("addr", s.cfg.ListenAddr).Info("starting REST API httpserver")
		errCh <- s.httpServer.ListenAndServe()
		close(errCh)
	}()

	select {
	case err := <-errCh:
		if err != http.ErrServerClosed {
			return fmt.Errorf("error: starting REST API httpserver: %w", err)
		}
	case sig := <-sigCh:
		s.cfg.Logger.WithField("signal", sig.String()).Infof("shutting down REST API httpserver due to signal")

		if err := s.shutdownGracefully(ctx, time.Minute); err != nil {
			return fmt.Errorf("graceful shutdown did not complete: %w", err)
		}

		wg.Wait()
		s.cfg.Logger.Info("REST API httpserver was shut down gracefully")
	}
	return nil
}

func (s *server) shutdownGracefully(ctx context.Context, timeout time.Duration) error {
	shutdownCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	if err := s.httpServer.Shutdown(shutdownCtx); err != nil {
		return err
	}

	return nil
}
