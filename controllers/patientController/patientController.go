package patientController

import (
	"hms-backend/dto"
	"hms-backend/usecases/patientUseCase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type patientController struct {
	usecase patientUseCase.PatientUseCase
}

func New(srv patientUseCase.PatientUseCase) *patientController {
	return &patientController{
		srv,
	}
}

func (ctrl *patientController) GetAll(c echo.Context) error {
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

func (ctrl *patientController) GetById(c echo.Context) error {
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

func (ctrl *patientController) Create(c echo.Context) error {
	var payload dto.Patient

	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	res, err := ctrl.usecase.Create(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = payload.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "success create data",
		Data:    res,
	})

}

func (ctrl *patientController) Update(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 0, 64)

	var payload dto.Patient

	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	res, err := ctrl.usecase.Update(uint(id), payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = payload.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "success update data",
		Data:    res,
	})
}

func (ctrl *patientController) Delete(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 0, 64)

	err := ctrl.usecase.Delete(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "success delete data",
		Data:    nil,
	})
}
