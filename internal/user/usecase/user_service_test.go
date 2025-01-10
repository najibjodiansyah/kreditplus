package usecase_test

// func TestUseCase_Login(t *testing.T) {
// 	t.Run("Success Login", func(t *testing.T) {
// 		mockUserRepo := new(mocks.MockUserRepository)
// 		mockNik := "1234567890"
// 		mockPass := "1234567890"
// 		mockUser := domain.User{
// 			Nik:      mockNik,
// 			Password: mockPass,
// 		}
// 		mockUserRepo.On("Login", mockNik).Return(mockUser, nil)
// 		token, _ := middleware.CreateToken(mockUser.Nik)
// 		assert.NotNil(t, token)
// 	})
// 	t.Run("Failed Login Repository", func(t *testing.T) {
// 		mockUserRepo := new(mocks.MockUserRepository)
// 		mockNik := "1234567890"
// 		mockPass := "1234567890"
// 		mockUser := domain.User{
// 			Nik:      mockNik,
// 			Password: mockPass,
// 		}

// 		fmt.Println(mockUser)

// 		e := echo.New()
// 		c := e.NewContext(nil, nil)

// 		mockUserRepo.On("Login", mock.Anything, mockNik).Return("", errors.New("error"))
// 		u := ucase.NewUserUseCase(mockUserRepo)

// 		a, err := u.Login(c, mockNik, mockPass)
// 		require.Error(t, err)

// 		fmt.Println(a)
// 	})
// 	t.Run("Failed Login Empty User", func(t *testing.T) {
// 		mockUserRepo := new(mocks.MockUserRepository)
// 		mockNik := "1234567890"
// 		mockPass := "1234567890"
// 		mockUser := domain.User{
// 			Nik:      mockNik,
// 			Password: mockPass,
// 		}
// 		mockUserRepo.On("Login", mockNik).Return(domain.User{}, nil)
// 		require.NotEqual(t, mockUser, domain.User{})
// 	})
// }
