package dashboardController

import (
	"github.com/labstack/echo/v4"
	"hms-backend/dto"
	"hms-backend/usecases/dashboardUseCase"
	"net/http"
	"strconv"
)

type dashboardController struct {
	usecase dashboardUseCase.DashboardUseCase
}

func New(srv dashboardUseCase.DashboardUseCase) *dashboardController {
	return &dashboardController{
		srv,
	}
}

func (ctrl *dashboardController) GetDataDashboardWeb(c echo.Context) error {
	res, err := ctrl.usecase.GetDataDashboardWeb()
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

func (ctrl *dashboardController) GetDataDashboardMobile(c echo.Context) error {
	doctorId, _ := strconv.ParseInt(c.Param("doctor_id"), 0, 64)

	res, err := ctrl.usecase.GetDataDashboardMobile(uint(doctorId))
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
