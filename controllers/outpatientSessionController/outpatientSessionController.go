package outpatientSessionController

import (
	"github.com/labstack/echo/v4"
	"hms-backend/dto"
	"hms-backend/usecases/outpatientSessionUseCase"
	"net/http"
	"strconv"
)

type outpatientSessionController struct {
	usecase outpatientSessionUseCase.OutpatientSessionUseCase
}

func New(srv outpatientSessionUseCase.OutpatientSessionUseCase) *outpatientSessionController {
	return &outpatientSessionController{
		srv,
	}
}
func (ctrl *outpatientSessionController) GetAll(c echo.Context) error {
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
func (ctrl *outpatientSessionController) GetById(c echo.Context) error {
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
func (ctrl *outpatientSessionController) GetByDoctorId(c echo.Context) error {
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
func (ctrl *outpatientSessionController) GetByPatientId(c echo.Context) error {
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
func (ctrl *outpatientSessionController) GetUnprocessedByDoctorId(c echo.Context) error {
	doctorId, _ := strconv.ParseInt(c.Param("doctor_id"), 16, 64)

	res, err := ctrl.usecase.GetUnprocessedByDoctorId(uint(doctorId))
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
func (ctrl *outpatientSessionController) GetProcessedByDoctorId(c echo.Context) error {
	doctorId, _ := strconv.ParseInt(c.Param("doctor_id"), 16, 64)

	res, err := ctrl.usecase.GetProcessedByDoctorId(uint(doctorId))
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
func (ctrl *outpatientSessionController) Create(c echo.Context) error {
	var payload dto.OutpatientSessionReq

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
func (ctrl *outpatientSessionController) Update(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 16, 64)

	var payload dto.OutpatientSessionReq

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

	return c.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "success update data",
		Data:    res,
	})
}
func (ctrl *outpatientSessionController) Approval(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 16, 64)

	var payload dto.ApprovalReq

	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	res, err := ctrl.usecase.Approval(uint(id), payload)
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
func (ctrl *outpatientSessionController) Delete(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 16, 64)

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
