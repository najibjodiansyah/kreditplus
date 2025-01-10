package http

import (
	"context"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/najibjodiansyah/kreditplus/config/middleware"
	"github.com/najibjodiansyah/kreditplus/config/rest"
	"github.com/najibjodiansyah/kreditplus/internal/user"
)

//go:generate mockery --name UserUsecase
type UserUsecase interface {
	Login(ctx context.Context, nik, pass string) (string, error)
	Create(ctx context.Context, us user.User) error
	Update(ctx context.Context, us user.User) error
}

type UserDelivery struct {
	UserUseCase UserUsecase
}

func NewUserDelivery(e *echo.Echo, us UserUsecase) {
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
		return c.JSON(400, rest.NewResponse(rest.Failed, err.Error(), nil))
	}

	if ok, err := validate(&input); !ok {
		return c.JSON(400, rest.NewResponse(rest.Failed, err.Error(), nil))
	}

	token, err := u.UserUseCase.Login(c.Request().Context(), input.Nik, input.Password)
	if err != nil {
		return c.JSON(400, rest.NewResponse(rest.Failed, err.Error(), nil))
	}

	return c.JSON(200, rest.NewResponse(rest.Success, "Login Success", token))
}

func (u *UserDelivery) Update(c echo.Context) error {
	var input UpdateUser
	if err := c.Bind(&input); err != nil {
		return c.JSON(400, rest.NewResponse(rest.Failed, err.Error(), nil))
	}

	if ok, err := validate(&input); !ok {
		return c.JSON(400, rest.NewResponse(rest.Failed, err.Error(), nil))
	}

	nik, err := middleware.GetNik(echo.Context(c))
	if err != nil {
		c.JSON(http.StatusUnauthorized, rest.NewResponse(rest.Failed, err.Error(), nil))
	}

	birthdate, err := time.Parse("2006-01-02", input.BirthDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, rest.NewResponse(rest.Failed, err.Error(), nil))
	}

	user := user.User{
		Nik:          nik,
		FullName:     input.FullName,
		LegalName:    input.LegalName,
		BirthPlace:   input.BirthPlace,
		BirthDate:    birthdate,
		Wages:        input.Wages,
		Photo_ktp:    input.Photo_ktp,
		Photo_selfie: input.Photo_selfie,
	}

	err = u.UserUseCase.Update(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, rest.NewResponse(rest.Failed, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, rest.NewResponse(rest.Success, "user updated", nil))
}

func (u *UserDelivery) Create(c echo.Context) (err error) {
	var input CreateUser

	err = c.Bind(&input)
	if err != nil {
		return c.JSON(400, rest.NewResponse(rest.Failed, err.Error(), nil))
	}

	if ok, err := validate(&input); !ok {
		return c.JSON(400, rest.NewResponse(rest.Failed, err.Error(), nil))
	}

	birthdate, err := time.Parse("2006-01-02", input.BirthDate)
	if err != nil {
		return c.JSON(400, rest.NewResponse(rest.Failed, err.Error(), nil))
	}

	user := user.User{
		Nik:        input.Nik,
		FullName:   input.FullName,
		Password:   input.Password,
		LegalName:  input.LegalName,
		BirthPlace: input.BirthPlace,
		BirthDate:  birthdate,
		Wages:      input.Wages,
	}

	err = u.UserUseCase.Create(c.Request().Context(), user)
	if err != nil {
		return c.JSON(400, rest.NewResponse(rest.Failed, err.Error(), nil))
	}

	return c.JSON(201, rest.NewResponse(rest.Success, "user created", nil))
}
