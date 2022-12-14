package historyController

import (
	"github.com/labstack/echo/v4"
	"hms-backend/dto"
	"hms-backend/usecases/historyUseCase"
	"net/http"
	"strconv"
)

type historyController struct {
	usecase historyUseCase.HistoryUseCase
}

func New(srv historyUseCase.HistoryUseCase) *historyController {
	return &historyController{
		srv,
	}
}
func (ctrl *historyController) GetOutpatientSessionHistory(c echo.Context) error {
	doctorId, _ := strconv.ParseInt(c.Param("doctor_id"), 0, 64)

	res, err := ctrl.usecase.GetOutpatientSessionHistory(uint(doctorId))
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
func (ctrl *historyController) GetApprovalHistory(c echo.Context) error {
	doctorId, _ := strconv.ParseInt(c.Param("doctor_id"), 0, 64)

	res, err := ctrl.usecase.GetApprovalHistory(uint(doctorId))
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
