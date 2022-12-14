package doctorScheduleController

import (
	"hms-backend/dto"
	"hms-backend/usecases/doctorScheduleUseCase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type doctorScheduleController struct {
	usecase doctorScheduleUseCase.DoctorScheduleUseCase
}

func New(srv doctorScheduleUseCase.DoctorScheduleUseCase) *doctorScheduleController {
	return &doctorScheduleController{
		srv,
	}
}

func (ctrl *doctorScheduleController) GetById(c echo.Context) error {
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
func (ctrl *doctorScheduleController) GetByDoctorId(c echo.Context) error {
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
func (ctrl *doctorScheduleController) GetByLicenseNumber(c echo.Context) error {
	licenseNumber := c.Param("license_number")

	res, err := ctrl.usecase.GetByLicenseNumber(licenseNumber)
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
func (ctrl *doctorScheduleController) Create(c echo.Context) error {
	var payload dto.DoctorScheduleReq

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
func (ctrl *doctorScheduleController) Update(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 16, 64)

	var payload dto.DoctorScheduleReq

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
func (ctrl *doctorScheduleController) Delete(c echo.Context) error {
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
