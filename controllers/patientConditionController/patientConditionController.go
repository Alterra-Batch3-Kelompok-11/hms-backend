package patientConditionController

import (
	"github.com/labstack/echo/v4"
	"hms-backend/dto"
	"hms-backend/usecases/patientConditionUseCase"
	"net/http"
	"strconv"
)

type patientConditionController struct {
	usecase patientConditionUseCase.PatientConditionUseCase
}

func New(srv patientConditionUseCase.PatientConditionUseCase) *patientConditionController {
	return &patientConditionController{
		srv,
	}
}

func (ctrl *patientConditionController) GetAll(c echo.Context) error {
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
func (ctrl *patientConditionController) GetById(c echo.Context) error {
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
func (ctrl *patientConditionController) GetByDoctorId(c echo.Context) error {
	doctorId, _ := strconv.ParseInt(c.Param("doctor_id"), 16, 64)

	res, err := ctrl.usecase.GetByDoctorId(uint(doctorId))
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
func (ctrl *patientConditionController) GetByPatientId(c echo.Context) error {
	patientId, _ := strconv.ParseInt(c.Param("patient_id"), 16, 64)

	res, err := ctrl.usecase.GetByPatientId(uint(patientId))
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
func (ctrl *patientConditionController) Create(c echo.Context) error {
	var payload dto.InsertPatientCondition

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

	return c.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "success create data",
		Data:    res,
	})
}
