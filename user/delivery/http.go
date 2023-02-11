package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/najibjodiansyah/kreditplus/domain"
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
	type request struct {
		Nik        int
		FullName   string
		Password   string
		LegalName  string
		BirthPlace string
		BirthDate  string
		Wages      int
	}

	var input request
	if err := c.Bind(&input); err != nil {
		return c.JSON(400, err)
	}

	user := domain.User{
		Nik:        input.Nik,
		FullName:   input.FullName,
		Password:   input.Password,
		LegalName:  input.LegalName,
		BirthPlace: input.BirthPlace,
		BirthDate:  input.BirthDate,
		Wages:      input.Wages,
	}

	err := u.UserUseCase.Create(c.Request().Context(), user)
	if err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(200, "Success")
}
