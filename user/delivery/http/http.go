package http

import (
	"github.com/labstack/echo/v4"
	"github.com/najibjodiansyah/kreditplus/domain"
	"github.com/najibjodiansyah/kreditplus/utils"
)

type UserDelivery struct {
	UserUseCase domain.UserService
}

func NewUserDelivery(e *echo.Echo, us domain.UserService) {
	handler := &UserDelivery{
		UserUseCase: us,
	}
	e.GET("/users/:nik", handler.GetByNik)
	e.POST("/users", handler.Create)
	e.PUT("/users", handler.Update)
}

func (u *UserDelivery) GetByNik(c echo.Context) error {
	panic("implement me")
}

func (u *UserDelivery) Update(c echo.Context) error {
	panic("implement me")
}

func (u *UserDelivery) Create(c echo.Context) error {
	var input requestUser
	if err := c.Bind(&input); err != nil {
		return c.JSON(400, utils.NewResponse(utils.Failed, err.Error(), nil))
	}

	user := domain.User{
		Nik:        input.Nik,
		FullName:   input.FullName,
		Password:   input.Password,
		LegalName:  input.LegalName,
		BirthPlace: input.BirthPlace,
		BirthDate:  input.BirthDate,
		Wages:      input.Wages,
		Photo: domain.Photo{
			Selfie: input.Photo.selfie,
			Ktp:    input.Photo.ktp,
		},
	}
	ctx := c.Request().Context()
	err := u.UserUseCase.Create(ctx, user)
	if err != nil {
		return c.JSON(400, utils.NewResponse(utils.Failed, err.Error(), nil))
	}

	return c.JSON(201, utils.NewResponse(utils.Success, "user created", nil))
}
