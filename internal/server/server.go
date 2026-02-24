package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type New struct {
	httpServer *http.Server
}

func (s *New) Start(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,          // 1 MB
		ReadTimeout:    10 * time.Second, // 10 minutes
		WriteTimeout:   10 * time.Second, // 10 minutes
	}

	return s.httpServer.ListenAndServe()
}

func (s *New) Stop(ctx context.Context) error {
	logrus.Info("initiating server shutdown")
    
    shutdownCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
    defer cancel()

    var shutdownErrors []error

    if err := s.httpServer.Shutdown(shutdownCtx); err != nil {
        shutdownErrors = append(shutdownErrors, fmt.Errorf("HTTP server shutdown: %w", err))
        
        if err := s.httpServer.Close(); err != nil {
            shutdownErrors = append(shutdownErrors, fmt.Errorf("HTTP server force close: %w", err))
        }
    }

    if len(shutdownErrors) > 0 {
        logrus.Error("server shutdown completed with errors")
        return fmt.Errorf("shutdown errors: %v", shutdownErrors)
    }

    logrus.Info("server shutdown completed successfully")
    return nil
}
