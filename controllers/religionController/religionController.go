package religionController

import (
	"github.com/labstack/echo/v4"
	"hms-backend/dto"
	"hms-backend/usecases/religionUseCase"
	"net/http"
	"strconv"
)

type religionController struct {
	usecase religionUseCase.ReligionUseCase
}

func New(srv religionUseCase.ReligionUseCase) *religionController {
	return &religionController{
		srv,
	}
}

func (ctrl *religionController) GetAll(c echo.Context) error {
	res, err := ctrl.usecase.GetAll()
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
func (ctrl *religionController) GetById(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 0, 64)

	res, err := ctrl.usecase.GetById(uint(id))
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
