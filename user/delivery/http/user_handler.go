package http

import (
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/najibjodiansyah/kreditplus/app/config/middleware"
	"github.com/najibjodiansyah/kreditplus/app/config/utils"
	"github.com/najibjodiansyah/kreditplus/domain"
)

type UserDelivery struct {
	UserUseCase domain.UserUsecase
}

func NewUserDelivery(e *echo.Echo, us domain.UserUsecase) {
	handler := &UserDelivery{
		UserUseCase: us,
	}

	e.POST("/users", handler.Create)
	e.PUT("/users", handler.Update, middleware.JWTMiddleware())
	e.POST("/users/login", handler.Login)
}

func validate(m interface{}) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *UserDelivery) Login(c echo.Context) error {
	var input LoginUser
	if err := c.Bind(&input); err != nil {
		return c.JSON(400, utils.NewResponse(utils.Failed, err.Error(), nil))
	}

	if ok, err := validate(&input); !ok {
		return c.JSON(400, utils.NewResponse(utils.Failed, err.Error(), nil))
	}

	token, err := u.UserUseCase.Login(c, input.Nik, input.Password)
	if err != nil {
		return c.JSON(400, utils.NewResponse(utils.Failed, err.Error(), nil))
	}

	return c.JSON(200, utils.NewResponse(utils.Success, "Login Success", token))
}

func (u *UserDelivery) Update(c echo.Context) error {
	var input UpdateUser
	if err := c.Bind(&input); err != nil {
		return c.JSON(400, utils.NewResponse(utils.Failed, err.Error(), nil))
	}

	if ok, err := validate(&input); !ok {
		return c.JSON(400, utils.NewResponse(utils.Failed, err.Error(), nil))
	}

	nik, err := middleware.GetNik(echo.Context(c))
	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(utils.Failed, err.Error(), nil))
	}

	birthdate, err := time.Parse("2006-01-02", input.BirthDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(utils.Failed, err.Error(), nil))
	}

	user := domain.User{
		Nik:          nik,
		FullName:     input.FullName,
		LegalName:    input.LegalName,
		BirthPlace:   input.BirthPlace,
		BirthDate:    birthdate,
		Wages:        input.Wages,
		Photo_ktp:    input.Photo_ktp,
		Photo_selfie: input.Photo_selfie,
	}

	err = u.UserUseCase.Update(c, user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(utils.Failed, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, utils.NewResponse(utils.Success, "user updated", nil))
}

func (u *UserDelivery) Create(c echo.Context) (err error) {
	var input CreateUser

	err = c.Bind(&input)
	if err != nil {
		return c.JSON(400, utils.NewResponse(utils.Failed, err.Error(), nil))
	}

	if ok, err := validate(&input); !ok {
		return c.JSON(400, utils.NewResponse(utils.Failed, err.Error(), nil))
	}

	birthdate, err := time.Parse("2006-01-02", input.BirthDate)
	if err != nil {
		return c.JSON(400, utils.NewResponse(utils.Failed, err.Error(), nil))
	}

	user := domain.User{
		Nik:        input.Nik,
		FullName:   input.FullName,
		Password:   input.Password,
		LegalName:  input.LegalName,
		BirthPlace: input.BirthPlace,
		BirthDate:  birthdate,
		Wages:      input.Wages,
	}

	err = u.UserUseCase.Create(c, user)
	if err != nil {
		return c.JSON(400, utils.NewResponse(utils.Failed, err.Error(), nil))
	}

	return c.JSON(201, utils.NewResponse(utils.Success, "user created", nil))
}
