package doctorScheduleController

import (
	"bytes"
	"encoding/json"
	"hms-backend/dto"
	"hms-backend/usecases/doctorScheduleUseCase/mock"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type doctorScheduleTestSuite struct {
	suite.Suite
	ctrl *doctorScheduleController
	mock *mock.DoctorScheduleUseCaseMock
}

func (s *doctorScheduleTestSuite) SetupSuite() {
	mock := &mock.DoctorScheduleUseCaseMock{}
	s.mock = mock

	s.ctrl = &doctorScheduleController{
		usecase: s.mock,
	}
}

func TestOrderSuite(t *testing.T) {
	suite.Run(t, new(doctorScheduleTestSuite))
}

func (s *doctorScheduleTestSuite) TestGetById() {
	s.mock.On("GetById", uint(1)).Return(
		true,
		Schedule,
	)

	s.mock.On("GetById", uint(2)).Return(
		false,
		dto.DoctorScheduleRes{},
	)

	testCases := []struct {
		Name               string
		Method             string
		Header             map[string][]string
		Id                 uint
		Path               string
		ExpectedStatusCode int
		ExpectedBody       dto.DoctorScheduleRes
	}{
		{
			"Success",
			http.MethodGet,
			map[string][]string{
				"Content-Type":    {"application/json"},
				"Accept":          {"*/*"},
				"Accept-Encoding": {"gzip", "deflate", "br"},
			},
			1,
			"schedules/:id",
			http.StatusOK,
			Schedule,
		},
		{
			"Success",
			http.MethodGet,
			map[string][]string{
				"Content-Type":    {"application/json"},
				"Accept":          {"*/*"},
				"Accept-Encoding": {"gzip", "deflate", "br"},
			},
			2,
			"schedules/:id",
			http.StatusBadRequest,
			dto.DoctorScheduleRes{},
		},
	}

	for _, testCase := range testCases {
		s.T().Run(testCase.Name, func(t *testing.T) {
			req := httptest.NewRequest(testCase.Method, "/", nil)
			rec := httptest.NewRecorder()
			req.Header = testCase.Header

			e := echo.New()
			ctx := e.NewContext(req, rec)

			ctx.SetParamNames("id")
			ctx.SetParamValues(strconv.Itoa(int(testCase.Id)))
			ctx.SetPath(testCase.Path)

			err := s.ctrl.GetById(ctx)
			s.NoError(err)
			s.Equal(testCase.ExpectedStatusCode, rec.Result().StatusCode)

			var response dto.Response
			errBdy := json.NewDecoder(rec.Result().Body).Decode(&response)
			s.NoError(errBdy)

			if response.Data != nil {
				var data dto.DoctorScheduleRes

				jsonStr, errMarsh := json.Marshal(response.Data)
				s.NoError(errMarsh)

				errUnMarsh := json.Unmarshal(jsonStr, &data)
				s.NoError(errUnMarsh)

				s.Equal(testCase.ExpectedBody, data)
			}
		})
	}
}

func (s *doctorScheduleTestSuite) TestGetByLicenseNumber() {
	s.mock.On("GetByLicenseNumber", "1234567890").Return(
		true,
		Schedule,
	)

	s.mock.On("GetByLicenseNumber", "").Return(
		false,
		dto.DoctorScheduleRes{},
	)

	testCases := []struct {
		Name               string
		Method             string
		Header             map[string][]string
		LicenseNumber      string
		Path               string
		ExpectedStatusCode int
		ExpectedBody       dto.DoctorScheduleRes
	}{
		{
			"Success",
			http.MethodGet,
			map[string][]string{
				"Content-Type":    {"application/json"},
				"Accept":          {"*/*"},
				"Accept-Encoding": {"gzip", "deflate", "br"},
			},
			"1234567890",
			"schedules/license_number/:license_number",
			http.StatusOK,
			Schedule,
		},
		{
			"Success",
			http.MethodGet,
			map[string][]string{
				"Content-Type":    {"application/json"},
				"Accept":          {"*/*"},
				"Accept-Encoding": {"gzip", "deflate", "br"},
			},
			"",
			"schedules/:id",
			http.StatusBadRequest,
			dto.DoctorScheduleRes{},
		},
	}

	for _, testCase := range testCases {
		s.T().Run(testCase.Name, func(t *testing.T) {
			req := httptest.NewRequest(testCase.Method, "/", nil)
			rec := httptest.NewRecorder()
			req.Header = testCase.Header

			e := echo.New()
			ctx := e.NewContext(req, rec)

			ctx.SetParamNames("license_number")
			ctx.SetParamValues(testCase.LicenseNumber)
			ctx.SetPath(testCase.Path)

			err := s.ctrl.GetByLicenseNumber(ctx)
			s.NoError(err)
			s.Equal(testCase.ExpectedStatusCode, rec.Result().StatusCode)

			var response dto.Response
			errBdy := json.NewDecoder(rec.Result().Body).Decode(&response)
			s.NoError(errBdy)

			if response.Data != nil {
				var data dto.DoctorScheduleRes

				jsonStr, errMarsh := json.Marshal(response.Data)
				s.NoError(errMarsh)

				errUnMarsh := json.Unmarshal(jsonStr, &data)
				s.NoError(errUnMarsh)

				s.Equal(testCase.ExpectedBody, data)
			}
		})
	}
}

func (s *doctorScheduleTestSuite) TestGetByDoctorId() {
	s.mock.On("GetBySpecialityId", uint(1)).Return(
		true,
		Schedules,
	)

	s.mock.On("GetBySpecialityId", uint(0)).Return(
		false,
		[]dto.DoctorScheduleRes{},
	)

	testCases := []struct {
		Name               string
		Method             string
		Header             map[string][]string
		SpecialityId       uint
		Path               string
		ExpectedStatusCode int
		ExpectedBody       []dto.DoctorScheduleRes
	}{
		{
			"Success",
			http.MethodGet,
			map[string][]string{
				"Content-Type":    {"application/json"},
				"Accept":          {"*/*"},
				"Accept-Encoding": {"gzip", "deflate", "br"},
			},
			1,
			"schedules/doctor/:id",
			http.StatusOK,
			Schedules,
		},
		{
			"Success",
			http.MethodGet,
			map[string][]string{
				"Content-Type":    {"application/json"},
				"Accept":          {"*/*"},
				"Accept-Encoding": {"gzip", "deflate", "br"},
			},
			0,
			"schedules/doctor/:id",
			http.StatusBadRequest,
			nil,
		},
	}

	for _, testCase := range testCases {
		s.T().Run(testCase.Name, func(t *testing.T) {
			req := httptest.NewRequest(testCase.Method, "/", nil)
			rec := httptest.NewRecorder()
			req.Header = testCase.Header

			e := echo.New()
			ctx := e.NewContext(req, rec)

			ctx.SetParamNames("doctor_id")
			ctx.SetParamValues(strconv.Itoa(int(testCase.SpecialityId)))
			ctx.SetPath(testCase.Path)

			err := s.ctrl.GetByDoctorId(ctx)
			s.NoError(err)
			s.Equal(testCase.ExpectedStatusCode, rec.Result().StatusCode)

			var response dto.Response
			errBdy := json.NewDecoder(rec.Result().Body).Decode(&response)
			s.NoError(errBdy)

			if response.Data != nil {
				var data []dto.DoctorScheduleRes

				jsonStr, errMarsh := json.Marshal(response.Data)
				s.NoError(errMarsh)

				errUnMarsh := json.Unmarshal(jsonStr, &data)
				s.NoError(errUnMarsh)

				s.Equal(testCase.ExpectedBody, data)
			}
		})
	}
}

func (s *doctorScheduleTestSuite) TestCreate() {
	s.mock.On("Create", ScheduleReq).Return(
		true,
		Schedule,
	)

	s.mock.On("Create", dto.DoctorScheduleReq{}).Return(
		false,
		dto.DoctorScheduleRes{},
	)

	testCases := []struct {
		Name               string
		Method             string
		Header             map[string][]string
		BodyParam          dto.DoctorScheduleReq
		ExpectedStatusCode int
		ExpectedBody       dto.DoctorScheduleRes
	}{
		{
			"Success",
			http.MethodPost,
			map[string][]string{
				"Content-Type":    {"application/json"},
				"Accept":          {"*/*"},
				"Accept-Encoding": {"gzip", "deflate", "br"},
			},
			ScheduleReq,
			http.StatusOK,
			Schedule,
		},
		{
			"Failed",
			http.MethodPost,
			map[string][]string{
				"Content-Type":    {"application/json"},
				"Accept":          {"*/*"},
				"Accept-Encoding": {"gzip", "deflate", "br"},
			},
			dto.DoctorScheduleReq{},
			http.StatusBadRequest,
			Schedule,
		},
	}

	for _, testCase := range testCases {
		s.T().Run(testCase.Name, func(t *testing.T) {
			body, _ := json.Marshal(testCase.BodyParam)

			req := httptest.NewRequest(testCase.Method, "/", bytes.NewBuffer(body))
			rec := httptest.NewRecorder()
			req.Header = testCase.Header

			e := echo.New()
			ctx := e.NewContext(req, rec)

			err := s.ctrl.Create(ctx)
			s.NoError(err)
			s.Equal(testCase.ExpectedStatusCode, rec.Result().StatusCode)

			var response dto.Response
			errBdy := json.NewDecoder(rec.Result().Body).Decode(&response)
			s.NoError(errBdy)

			if response.Data != nil {
				var data dto.DoctorScheduleRes

				jsonStr, errMarsh := json.Marshal(response.Data)
				s.NoError(errMarsh)

				errUnMarsh := json.Unmarshal(jsonStr, &data)
				s.NoError(errUnMarsh)

				s.Equal(testCase.ExpectedBody, data)
			}
		})
	}
}

func (s *doctorScheduleTestSuite) TestUpdate() {
	s.mock.On("Update", uint(1), ScheduleReq).Return(
		true,
		Schedule,
	)

	s.mock.On("Update", uint(0), dto.DoctorScheduleReq{}).Return(
		false,
		dto.NurseRes{},
	)

	testCases := []struct {
		Name               string
		Method             string
		Header             map[string][]string
		Id                 uint
		Path               string
		BodyParam          dto.DoctorScheduleReq
		ExpectedStatusCode int
		ExpectedBody       dto.DoctorScheduleRes
	}{
		{
			"Success",
			http.MethodPut,
			map[string][]string{
				"Content-Type":    {"application/json"},
				"Accept":          {"*/*"},
				"Accept-Encoding": {"gzip", "deflate", "br"},
			},
			1,
			"schedules/:id",
			ScheduleReq,
			http.StatusOK,
			Schedule,
		},
		{
			"Failed",
			http.MethodPut,
			map[string][]string{
				"Content-Type":    {"application/json"},
				"Accept":          {"*/*"},
				"Accept-Encoding": {"gzip", "deflate", "br"},
			},
			0,
			"Nurses/:id",
			dto.DoctorScheduleReq{},
			http.StatusBadRequest,
			Schedule,
		},
	}

	for _, testCase := range testCases {
		s.T().Run(testCase.Name, func(t *testing.T) {
			body, _ := json.Marshal(testCase.BodyParam)

			req := httptest.NewRequest(testCase.Method, "/", bytes.NewBuffer(body))
			rec := httptest.NewRecorder()
			req.Header = testCase.Header

			e := echo.New()
			ctx := e.NewContext(req, rec)

			ctx.SetParamNames("id")
			ctx.SetParamValues(strconv.Itoa(int(testCase.Id)))
			ctx.SetPath(testCase.Path)

			err := s.ctrl.Update(ctx)
			s.NoError(err)
			s.Equal(testCase.ExpectedStatusCode, rec.Result().StatusCode)

			var response dto.Response
			errBdy := json.NewDecoder(rec.Result().Body).Decode(&response)
			s.NoError(errBdy)

			if response.Data != nil {
				var data dto.DoctorScheduleRes

				jsonStr, errMarsh := json.Marshal(response.Data)
				s.NoError(errMarsh)

				errUnMarsh := json.Unmarshal(jsonStr, &data)
				s.NoError(errUnMarsh)

				s.Equal(testCase.ExpectedBody, data)
			}
		})
	}
}

func (s *doctorScheduleTestSuite) TestDelete() {
	s.mock.On("Delete", uint(1)).Return(
		true,
	)

	s.mock.On("Delete", uint(0)).Return(
		false,
	)

	testCases := []struct {
		Name               string
		Method             string
		Header             map[string][]string
		Id                 uint
		Path               string
		ExpectedStatusCode int
		ExpectedBody       dto.DoctorScheduleRes
	}{
		{
			"Success",
			http.MethodPut,
			map[string][]string{
				"Content-Type":    {"application/json"},
				"Accept":          {"*/*"},
				"Accept-Encoding": {"gzip", "deflate", "br"},
			},
			1,
			"schedules/:id",
			http.StatusOK,
			Schedule,
		},
		{
			"Failed",
			http.MethodPut,
			map[string][]string{
				"Content-Type":    {"application/json"},
				"Accept":          {"*/*"},
				"Accept-Encoding": {"gzip", "deflate", "br"},
			},
			0,
			"schedules/:id",
			http.StatusBadRequest,
			Schedule,
		},
	}

	for _, testCase := range testCases {
		s.T().Run(testCase.Name, func(t *testing.T) {

			req := httptest.NewRequest(testCase.Method, "/", nil)
			rec := httptest.NewRecorder()
			req.Header = testCase.Header

			e := echo.New()
			ctx := e.NewContext(req, rec)

			ctx.SetParamNames("id")
			ctx.SetParamValues(strconv.Itoa(int(testCase.Id)))
			ctx.SetPath(testCase.Path)

			err := s.ctrl.Delete(ctx)
			s.NoError(err)
			s.Equal(testCase.ExpectedStatusCode, rec.Result().StatusCode)

			var response dto.Response
			errBdy := json.NewDecoder(rec.Result().Body).Decode(&response)
			s.NoError(errBdy)
		})
	}
}

// Data Mock

var Schedules = []dto.DoctorScheduleRes{
	{
		ID:        1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
		DoctorId:  1,
		DayInt:    1,
		DayString: "Monday",
		StartTime: "08:00:00",
		EndTime:   "16:00:00",
	},
	{
		ID:        2,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
		DoctorId:  2,
		DayInt:    2,
		DayString: "Monday",
		StartTime: "08:00:00",
		EndTime:   "16:00:00",
	},
}

var Schedule = dto.DoctorScheduleRes{

	ID:        1,
	CreatedAt: time.Time{},
	UpdatedAt: time.Time{},
	DeletedAt: gorm.DeletedAt{},
	DoctorId:  1,
	DayInt:    1,
	DayString: "Monday",
	StartTime: "08:00:00",
	EndTime:   "16:00:00",
}

var ScheduleReq = dto.DoctorScheduleReq{
	DoctorId:  1,
	DayInt:    1,
	StartTime: "08:00:00",
	EndTime:   "16:00:00",
}
