package http

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/najibjodiansyah/kreditplus/config/middleware"
	"github.com/najibjodiansyah/kreditplus/config/rest"
	"github.com/najibjodiansyah/kreditplus/internal/limit"
)

type LimitUsecase interface {
	Create(ctx context.Context, lm limit.Limit, nik string) error
	GetById(ctx context.Context, Id int) ([]limit.Limit, error)
}

type LimitHandler struct {
	LimitUseCase LimitUsecase
}

func validate(m interface{}) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewLimitHandler(e *echo.Echo, lu LimitUsecase) {
	handler := &LimitHandler{
		LimitUseCase: lu,
	}

	e.POST("/limit", handler.CreateLimit)
	e.GET("/limit/:id", handler.GetLimitById)
}

func (lh *LimitHandler) CreateLimit(c echo.Context) error {
	var input limit.Limit
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(400, err.Error())
	}

	if ok, err := validate(&input); !ok {
		return c.JSON(400, rest.NewResponse(rest.Failed, err.Error(), nil))
	}

	nik, err := middleware.GetNik(echo.Context(c))
	if err != nil {
		c.JSON(http.StatusUnauthorized, rest.NewResponse(rest.Failed, err.Error(), nil))
	}

	err = lh.LimitUseCase.Create(c.Request().Context(), input, nik)
	if err != nil {
		return c.JSON(400, err.Error())
	}

	return c.JSON(200, "Success")
}

func (lh *LimitHandler) GetLimitById(c echo.Context) error {
	type Limit_user_id struct {
		Id int `param:"id"`
	}
	var input Limit_user_id
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(400, err.Error())
	}

	if ok, err := validate(&input); !ok {
		return c.JSON(400, rest.NewResponse(rest.Failed, err.Error(), nil))
	}

	limits, err := lh.LimitUseCase.GetById(c.Request().Context(), input.Id)
	if err != nil {
		return c.JSON(400, err.Error())
	}

	return c.JSON(200, limits)

}
