package config

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_logmon "github.com/labstack/gommon/log"
	"github.com/najibjodiansyah/kreditplus/app/config/database"
	_limitRepo "github.com/najibjodiansyah/kreditplus/limit/repository/mysql"
	_transactionRepo "github.com/najibjodiansyah/kreditplus/transaction/repository/mysql"
	_userHandler "github.com/najibjodiansyah/kreditplus/user/delivery/http"
	_userRepo "github.com/najibjodiansyah/kreditplus/user/repository/mysql"
	_userUseCase "github.com/najibjodiansyah/kreditplus/user/usecase"
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

	userRepo := _userRepo.NewMysqlUserRepository(db)
	limitRepo := _limitRepo.NewMysqlLimitRepository(db)
	transactionRepo := _transactionRepo.NewMysqlTransactionRepository(db)
	log.Println(transactionRepo)
	userUseCase := _userUseCase.NewUserService(userRepo, limitRepo)

	_userHandler.NewUserDelivery(e, userUseCase)

	port := ":" + viper.GetString("HTTP_PORT")

	log.Fatal(e.Start(port))

}
