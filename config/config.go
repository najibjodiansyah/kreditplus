package config

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_logmon "github.com/labstack/gommon/log"
	"github.com/najibjodiansyah/kreditplus/config/database"
	_limitRepo "github.com/najibjodiansyah/kreditplus/internal/limit/repository/mysql"
	_transactionRepo "github.com/najibjodiansyah/kreditplus/internal/transaction/repository/mysql"
	_userHandler "github.com/najibjodiansyah/kreditplus/internal/user/delivery/http"
	_userRepo "github.com/najibjodiansyah/kreditplus/internal/user/repository/mysql"
	_userUseCase "github.com/najibjodiansyah/kreditplus/internal/user/usecase"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`.env`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// if viper.GetBool(`debug`) {
	// 	log.Println("Service RUN on DEBUG mode")
	// }
}

func Run() {
	db := database.InitDB()
	e := echo.New()

	logger := logrus.New()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			logger.WithFields(logrus.Fields{
				"URI":    values.URI,
				"status": values.Status,
			}).Info("request")

			return nil
		},
	}))
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
		LogLevel:  _logmon.ERROR,
	}))
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "",
		ContentTypeNosniff:    "",
		XFrameOptions:         "",
		HSTSMaxAge:            3600,
		ContentSecurityPolicy: "default-src 'self'",
	}))

	userRepo := _userRepo.NewMysqlUserRepository(db)
	limitRepo := _limitRepo.NewMysqlLimitRepository(db)
	transactionRepo := _transactionRepo.NewMysqlTransactionRepository(db)
	log.Println(transactionRepo, limitRepo)
	userUseCase := _userUseCase.NewUserUseCase(userRepo)

	_userHandler.NewUserDelivery(e, userUseCase)

	port := ":" + viper.GetString("HTTP_PORT")

	go func() {
		if err := e.Start(port); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
