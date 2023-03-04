package http_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/najibjodiansyah/kreditplus/domain/mocks"
	user_handler "github.com/najibjodiansyah/kreditplus/user/delivery/http"
)

func TestUserDelivery_Login(t *testing.T) {
	t.Run("Login Success", func(t *testing.T) {
		mockUserUseCase := new(mocks.MockUserUseCase)

		mockLogin := user_handler.LoginUser{
			Nik:      "1234567890",
			Password: "1234567890",
		}

		json, err := json.Marshal(mockLogin)
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		mockUserUseCase.EXPECT().Login(mock.Anything, mockLogin.Nik, mockLogin.Password).Return("token", nil)

		e := echo.New()
		req, err := http.NewRequestWithContext(context.TODO(), echo.GET, "/users/login", strings.NewReader(string(json)))
		assert.NoError(t, err)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/login")

		h := user_handler.UserDelivery{
			UserUseCase: mockUserUseCase,
		}
		err = h.Login(c)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, rec.Code)

		mockUserUseCase.AssertExpectations(t)
	})

	t.Run("Login Error", func(t *testing.T) {
		mockUserUseCase := new(mocks.MockUserUseCase)

		mockLogin := user_handler.LoginUser{
			Nik:      "1234567890",
			Password: "1234567890",
		}

		json, err := json.Marshal(mockLogin)
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		mockUserUseCase.EXPECT().Login(mock.Anything, mockLogin.Nik, mockLogin.Password).Return("", errors.New("Error"))

		e := echo.New()
		req, err := http.NewRequestWithContext(context.TODO(), echo.GET, "/users/login", strings.NewReader(string(json)))
		assert.NoError(t, err)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/login")

		h := user_handler.UserDelivery{
			UserUseCase: mockUserUseCase,
		}
		err = h.Login(c)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rec.Code)

		mockUserUseCase.AssertExpectations(t)
	})

}
