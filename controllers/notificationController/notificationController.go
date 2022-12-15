package notificationController

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"hms-backend/configs"
	"hms-backend/dto"
	"hms-backend/usecases/notificationUseCase"
	"net/http"
	"strings"
)

type notificationController struct {
	usecase notificationUseCase.NotificationUseCase
}

func New(srv notificationUseCase.NotificationUseCase) *notificationController {
	return &notificationController{
		srv,
	}
}

func (ctrl *notificationController) GetByDoctorId(c echo.Context) error {
	headerToken := c.Request().Header.Get("Authorization")
	token := strings.Split(headerToken, " ")[1]
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(configs.Cfg.JwtKey), nil
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unable to parse token")
	}

	userId := uint(claims["roleId"].(float64))

	res, err := ctrl.usecase.GetByDoctorId(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "success get data",
		Data:    res,
	})
}
