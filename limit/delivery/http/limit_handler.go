package http

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/najibjodiansyah/kreditplus/app/config/middleware"
	"github.com/najibjodiansyah/kreditplus/app/config/utils"
	"github.com/najibjodiansyah/kreditplus/domain"
)

type LimitHandler struct {
	LimitUseCase domain.LimitUsecase
}

func validate(m interface{}) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewLimitHandler(e *echo.Echo, lu domain.LimitUsecase) {
	handler := &LimitHandler{
		LimitUseCase: lu,
	}

	e.POST("/limit", handler.CreateLimit)
	e.GET("/limit/:id", handler.GetLimitById)
}

func (lh *LimitHandler) CreateLimit(c echo.Context) error {
	var input domain.Limit
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(400, err.Error())
	}

	if ok, err := validate(&input); !ok {
		return c.JSON(400, utils.NewResponse(utils.Failed, err.Error(), nil))
	}

	nik, err := middleware.GetNik(echo.Context(c))
	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(utils.Failed, err.Error(), nil))
	}

	err = lh.LimitUseCase.Create(c, input, nik)
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
		return c.JSON(400, utils.NewResponse(utils.Failed, err.Error(), nil))
	}

	limits, err := lh.LimitUseCase.GetById(c, input.Id)
	if err != nil {
		return c.JSON(400, err.Error())
	}

	return c.JSON(200, limits)

}
