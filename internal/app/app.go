package app

import (
	"URL-Shortener/internal/handler"
	database "URL-Shortener/internal/repository"
	"URL-Shortener/internal/server"
	service "URL-Shortener/internal/services"
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// TODO добавить makefile

func Run() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.SetLevel(logrus.TraceLevel)

	if err := initConfig(); err != nil {
		logrus.Fatalf("Failed to initialize config: %v", err)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error loading .env file: %v", err)
	}

	conn := dbInit()
	repo := database.New(conn)
	services := service.New(repo)

	handlerInit := handler.NewHandlers(services)
	handler := handler.New(handlerInit)

	srv := start(handler)
	// TODO сделать нормальное отлов паники

	if err := stop(srv, conn); err != nil {
		logrus.Errorf("Error stopping the server: %v", err)
	}
}

func start(h *handler.Handler) *server.New {
	srv := new(server.New)

	go func() {
		logrus.Debug("1")
		if err := srv.Start(
			viper.GetString("server.port"),
			h.InitRoutes(),
		); err != nil {
			logrus.Errorf("Failed to run server: %v", err)
			os.Exit(1)
		}
	}()

	logrus.Info("Server started")
	return srv
}

func stop(srv *server.New, conn *pgx.Conn) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := database.Close(conn); err != nil {
		return err
	}

	if err := srv.Stop(context.Background()); err != nil {
		return err
	}

	logrus.Info("Server stopped")
	return nil
}
