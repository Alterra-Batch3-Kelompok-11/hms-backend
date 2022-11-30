package doctorController

import (
	"github.com/labstack/echo/v4"
	"hms-backend/dto"
	"hms-backend/usecases/doctorUseCase"
	"net/http"
	"strconv"
)

type doctorController struct {
	usecase doctorUseCase.DoctorUseCase
}

func New(srv doctorUseCase.DoctorUseCase) *doctorController {
	return &doctorController{
		srv,
	}
}

func (ctrl *doctorController) GetAll(c echo.Context) error {
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
func (ctrl *doctorController) GetById(c echo.Context) error {
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
func (ctrl *doctorController) GetByLicenseNumber(c echo.Context) error {
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
func (ctrl *doctorController) GetBySpecialityId(c echo.Context) error {
	specialityId, _ := strconv.ParseInt(c.Param("speciality_id"), 16, 64)

	res, err := ctrl.usecase.GetBySpecialityId(uint(specialityId))
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
func (ctrl *doctorController) Create(c echo.Context) error {
	var payload dto.UserReq

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
func (ctrl *doctorController) Update(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 16, 64)

	var payload dto.UserReq

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
func (ctrl *doctorController) Delete(c echo.Context) error {
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
