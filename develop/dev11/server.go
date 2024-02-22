package server

import (
	"context"
	"go.uber.org/zap"
	"net/http"
	"os"
	"server/pkg/handler"
	store2 "server/pkg/service/store"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string) error {
	logger, err := zap.NewProduction()
	if err != nil {
		os.Exit(1)
	}
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {

		}
	}(logger)

	zap.ReplaceGlobals(logger)

	logger.Info("reading config")

	s.httpServer = &http.Server{
		Addr:           ":" + port,
		MaxHeaderBytes: 1 << 20, // MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	str := store2.NewInMemoryStore()

	h := handler.Handler{
		Store: &str,
	}

	err = h.InitRoutes()
	if err != nil {
		return err
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
