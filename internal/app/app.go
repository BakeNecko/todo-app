// Package app configures and runs application.
package app

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xhit/go-str2duration/v2"
	"os"
	"os/signal"
	"syscall"
	"todo-app/config"
	v1 "todo-app/internal/controller/http/v1"
	"todo-app/internal/usecase"
	"todo-app/internal/usecase/repository"
	"todo-app/internal/usecase/service"
	"todo-app/pkg/httpserver"
	"todo-app/pkg/jwt"
	"todo-app/pkg/logger"
	"todo-app/pkg/mongodb"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)
	mongo, err := mdb.New(context.TODO(), cfg.MongoDB.URL, cfg.MongoDB.Database)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - mongodb.New: %w", err))
	}

	// Repository
	reps := repository.NewRepository(mongo.Database)

	// UseCase layer
	jwtManger, err := jwt.NewTokenManger(cfg.Secret.Key)
	if err != nil {
		panic(fmt.Sprintf("can't create JWT manger. Errror %v", err))
	}
	accessTokenTTL, err := str2duration.ParseDuration(cfg.Auth.AccessTokenTTL)
	if err != nil {
		panic(err)
	}
	refreshTokenTTL, err := str2duration.ParseDuration(cfg.Auth.RefreshTokenTTL)
	if err != nil {
		panic(err)
	}

	useCases := usecase.NewUseCases(
		&usecase.Deps{
			AccountRepository: reps.AccountRepository,
			AuthRepository:    reps.AuthRepository,
			TodoRepository:    reps.TodoRepository,

			Hasher:          service.NewSHA1Hasher(cfg.Secret.Salt),
			JwtManager:      jwtManger,
			RefreshTokenTTL: refreshTokenTTL,
			AccessTokenTTL:  accessTokenTTL,
		},
	)

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, l, &v1.Deps{
		UseCases:        useCases,
		JwtTokenManager: jwtManger,
	})
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
