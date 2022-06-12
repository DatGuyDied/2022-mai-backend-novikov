package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/DatGuyDied/2022-mai-backend-novikov/internal/auth"
	"github.com/DatGuyDied/2022-mai-backend-novikov/internal/handlers"
	"github.com/DatGuyDied/2022-mai-backend-novikov/internal/repository"
	"github.com/DatGuyDied/2022-mai-backend-novikov/internal/services"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const shutdownTimeout = 10 * time.Second

func main() {
	logger := newLogger()

	repo := repository.NewMemoryRepository()

	pk, err := rsa.GenerateKey(rand.Reader, rsaBits)
	if err != nil {
		logger.Fatal("Can't generate rsa key", zap.Error(err))
	}
	authHandler := newAuthHandler(pk, repo)
	messagesHandler := newMessagesHandler(repo)
	usersHandler := newUsersHandler(repo)

	router := chi.NewRouter()
	router.Mount(`/auth`, authHandler.Routes())

	authRouter := router.With(
		auth.Middleware(&pk.PublicKey),
	)
	authRouter.Mount(`/messages`, messagesHandler.Routes())
	authRouter.Mount(`/users`, usersHandler.Routes())

	httpServer := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	done := make(chan struct{})
	go func() {
		defer close(done)

		logger.Info("Server started",
			zap.String(`addr`, httpServer.Addr),
		)
		serveErr := httpServer.ListenAndServe()
		if errors.Is(serveErr, http.ErrServerClosed) {
			logger.Info("Server closed")
		} else {
			logger.Error("Serve http failed", zap.Error(serveErr))
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	select {
	case <-done:
		// Я не знаю английский.
		logger.Fatal("Unexpected server closed")
	case <-signals:
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		logger.Error("shutdown failed")
	}
}

const rsaBits = 2048

func newAuthHandler(
	pk *rsa.PrivateKey,
	repo services.UsersRepository,
) *handlers.Auth {
	serivce := services.NewAuthService(repo)

	return handlers.NewAuth(serivce, pk)
}

func newMessagesHandler(
	repo services.MessagesRepository,
) *handlers.Messages {
	service := services.NewMessagesService(repo)

	return handlers.NewMessages(service)
}

func newUsersHandler(
	repo services.UserMessagesRepository,
) *handlers.Users {
	service := services.NewUserMessagesService(repo)

	return handlers.NewUsers(service)
}

func newLogger() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.Development = false

	logger, err := config.Build()
	if err != nil {
		panic(err)
	}

	return logger
}
