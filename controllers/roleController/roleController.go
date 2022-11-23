package roleController

import (
	"github.com/labstack/echo/v4"
	"hms-backend/dto"
	"hms-backend/usecases/roleUseCase"
	"net/http"
	"strconv"
)

type roleController struct {
	usecase roleUseCase.RoleUseCase
}

func New(srv roleUseCase.RoleUseCase) *roleController {
	return &roleController{
		srv,
	}
}

func (ctrl *roleController) GetAll(c echo.Context) error {
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

func (ctrl *roleController) GetById(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 16, 64)

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
