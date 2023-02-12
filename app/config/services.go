package config

import (
	"log"

	"github.com/labstack/echo/v4"
	_userHandler "github.com/najibjodiansyah/kreditplus/user/delivery/http"
	_userRepo "github.com/najibjodiansyah/kreditplus/user/repository/mysql"
	_userUseCase "github.com/najibjodiansyah/kreditplus/user/usecase"
	"github.com/spf13/viper"
)

func Run() {
	db := InitDB()
	e := echo.New()
	userRepo := _userRepo.NewMysqlUserRepository(db)
	userUseCase := _userUseCase.NewUserService(userRepo)
	_userHandler.NewUserDelivery(e, userUseCase)

	port := ":" + viper.GetString("HTTP_PORT")

	log.Fatal(e.Start(port))

}
