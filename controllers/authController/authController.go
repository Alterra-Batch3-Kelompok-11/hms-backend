package authController

import (
	"github.com/labstack/echo/v4"
	"hms-backend/dto"
	"hms-backend/usecases/authUseCase"
	"net/http"
)

type authController struct {
	usecase authUseCase.AuthUseCase
}

func New(srv authUseCase.AuthUseCase) *authController {
	return &authController{
		srv,
	}
}

func (ctrl *authController) Login(c echo.Context) error {
	payload := dto.UserReq{}

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	loginRes, err := ctrl.usecase.Login(payload.Username, payload.Password)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "login success",
		Data:    loginRes,
	})

}

func (ctrl *authController) SignUp(c echo.Context) error {
	payload := dto.UserReq{}

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if !(payload.RoleID == 2 || payload.RoleID == 3) { // jika role bukan doctor[2] atau nurse[3] maka ditolak
		return c.JSON(http.StatusBadRequest, dto.Response{
			Status:  http.StatusBadRequest,
			Message: "role id is not valid",
			Data:    nil,
		})
	}

	user, err := ctrl.usecase.SignUp(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "signup success",
		Data:    user,
	})

}
