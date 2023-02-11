package config

import (
	"log"

	"github.com/labstack/echo/v4"
	_userHandler "github.com/najibjodiansyah/kreditplus/user/delivery"
	_userRepo "github.com/najibjodiansyah/kreditplus/user/repository/mysql"
	_userUseCase "github.com/najibjodiansyah/kreditplus/user/usecase"
	"github.com/spf13/viper"
)

func Run() {
	db := InitDB()
	e := echo.New()
	userReoo := _userRepo.NewMysqlUserRepository(db)
	userUseCase := _userUseCase.NewUserService(userReoo)
	_userHandler.NewUserDelivery(e, userUseCase)

	log.Fatal(e.Start(viper.GetString("HTTP_PORT")))

}
