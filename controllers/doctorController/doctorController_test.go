package doctorController

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"hms-backend/dto"
	"hms-backend/usecases/doctorUseCase/mock"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

type doctorTestSuite struct {
	suite.Suite
	ctrl *doctorController
	mock *mock.DoctorUseCaseMock
}

func (s *doctorTestSuite) SetupSuite() {
	mock := &mock.DoctorUseCaseMock{}
	s.mock = mock

	s.ctrl = &doctorController{
		usecase: s.mock,
	}
}

func TestOrderSuite(t *testing.T) {
	suite.Run(t, new(doctorTestSuite))
}

func (s *doctorTestSuite) TestGetAll() {
	s.mock.On("GetAll").Return(
		true,
		Doctors,
	)

	testCases := []struct {
		Name               string
		Method             string
		Header             map[string][]string
		ExpectedStatusCode int
		ExpectedBody       []dto.DoctorRes
	}{
		{
			"Success",
			http.MethodGet,
			map[string][]string{
				"Content-Type":    {"application/json"},
				"Accept":          {"*/*"},
				"Accept-Encoding": {"gzip", "deflate", "br"},
			},
			http.StatusOK,
			Doctors,
		},
	}

	for _, testCase := range testCases {
		s.T().Run(testCase.Name, func(t *testing.T) {
			req := httptest.NewRequest(testCase.Method, "/", nil)
			rec := httptest.NewRecorder()
			req.Header = testCase.Header

			e := echo.New()
			ctx := e.NewContext(req, rec)

			err := s.ctrl.GetAll(ctx)
			s.NoError(err)
			s.Equal(testCase.ExpectedStatusCode, rec.Result().StatusCode)

			var response dto.Response
			errBdy := json.NewDecoder(rec.Result().Body).Decode(&response)
			s.NoError(errBdy)

			if response.Data != nil {
				var data []dto.DoctorRes

				jsonStr, errMarsh := json.Marshal(response.Data.([]interface{}))
				s.NoError(errMarsh)

				errUnMarsh := json.Unmarshal(jsonStr, &data)
				s.NoError(errUnMarsh)

				s.Equal(testCase.ExpectedBody, data)
			}
		})
	}
}
func (s *doctorTestSuite) TestGetById() {
	s.mock.On("GetById", uint(1)).Return(
		true,
		Doctor,
	)

	s.mock.On("GetById", uint(2)).Return(
		false,
		dto.DoctorRes{},
	)

	testCases := []struct {
		Name               string
		Method             string
		Header             map[string][]string
		Id                 uint
		Path               string
		ExpectedStatusCode int
		ExpectedBody       dto.DoctorRes
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
			"doctors/:id",
			http.StatusOK,
			Doctor,
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
			"doctors/:id",
			http.StatusBadRequest,
			dto.DoctorRes{},
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
				var data dto.DoctorRes

				jsonStr, errMarsh := json.Marshal(response.Data)
				s.NoError(errMarsh)

				errUnMarsh := json.Unmarshal(jsonStr, &data)
				s.NoError(errUnMarsh)

				s.Equal(testCase.ExpectedBody, data)
			}
		})
	}
}
func (s *doctorTestSuite) TestGetByLicenseNumber() {
	s.mock.On("GetByLicenseNumber", "1234567890").Return(
		true,
		Doctor,
	)

	s.mock.On("GetByLicenseNumber", "").Return(
		false,
		dto.DoctorRes{},
	)

	testCases := []struct {
		Name               string
		Method             string
		Header             map[string][]string
		LicenseNumber      string
		Path               string
		ExpectedStatusCode int
		ExpectedBody       dto.DoctorRes
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
			"doctors/license_number/:license_number",
			http.StatusOK,
			Doctor,
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
			"doctors/:id",
			http.StatusBadRequest,
			dto.DoctorRes{},
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
				var data dto.DoctorRes

				jsonStr, errMarsh := json.Marshal(response.Data)
				s.NoError(errMarsh)

				errUnMarsh := json.Unmarshal(jsonStr, &data)
				s.NoError(errUnMarsh)

				s.Equal(testCase.ExpectedBody, data)
			}
		})
	}
}
func (s *doctorTestSuite) TestGetBySpecialityId() {
	s.mock.On("GetBySpecialityId", uint(1)).Return(
		true,
		Doctors,
	)

	s.mock.On("GetBySpecialityId", uint(0)).Return(
		false,
		[]dto.DoctorRes{},
	)

	testCases := []struct {
		Name               string
		Method             string
		Header             map[string][]string
		SpecialityId       uint
		Path               string
		ExpectedStatusCode int
		ExpectedBody       []dto.DoctorRes
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
			"doctors/speciality/:id",
			http.StatusOK,
			Doctors,
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
			"doctors/speciality/:id",
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

			ctx.SetParamNames("speciality_id")
			ctx.SetParamValues(strconv.Itoa(int(testCase.SpecialityId)))
			ctx.SetPath(testCase.Path)

			err := s.ctrl.GetBySpecialityId(ctx)
			s.NoError(err)
			s.Equal(testCase.ExpectedStatusCode, rec.Result().StatusCode)

			var response dto.Response
			errBdy := json.NewDecoder(rec.Result().Body).Decode(&response)
			s.NoError(errBdy)

			if response.Data != nil {
				var data []dto.DoctorRes

				jsonStr, errMarsh := json.Marshal(response.Data)
				s.NoError(errMarsh)

				errUnMarsh := json.Unmarshal(jsonStr, &data)
				s.NoError(errUnMarsh)

				s.Equal(testCase.ExpectedBody, data)
			}
		})
	}
}
func (s *doctorTestSuite) TestGetToday() {
	s.mock.On("GetToday").Return(
		true,
		Doctors,
	)

	testCases := []struct {
		Name               string
		Method             string
		Header             map[string][]string
		ExpectedStatusCode int
		ExpectedBody       []dto.DoctorRes
	}{
		{
			"Success",
			http.MethodGet,
			map[string][]string{
				"Content-Type":    {"application/json"},
				"Accept":          {"*/*"},
				"Accept-Encoding": {"gzip", "deflate", "br"},
			},
			http.StatusOK,
			Doctors,
		},
	}

	for _, testCase := range testCases {
		s.T().Run(testCase.Name, func(t *testing.T) {
			req := httptest.NewRequest(testCase.Method, "/", nil)
			rec := httptest.NewRecorder()
			req.Header = testCase.Header

			e := echo.New()
			ctx := e.NewContext(req, rec)

			err := s.ctrl.GetToday(ctx)
			s.NoError(err)
			s.Equal(testCase.ExpectedStatusCode, rec.Result().StatusCode)

			var response dto.Response
			errBdy := json.NewDecoder(rec.Result().Body).Decode(&response)
			s.NoError(errBdy)

			if response.Data != nil {
				var data []dto.DoctorRes

				jsonStr, errMarsh := json.Marshal(response.Data.([]interface{}))
				s.NoError(errMarsh)

				errUnMarsh := json.Unmarshal(jsonStr, &data)
				s.NoError(errUnMarsh)

				s.Equal(testCase.ExpectedBody, data)
			}
		})
	}
}
func (s *doctorTestSuite) TestCreate() {
	s.mock.On("Create", DoctorReq).Return(
		true,
		Doctor,
	)

	s.mock.On("Create", dto.DoctorReq{}).Return(
		false,
		dto.DoctorRes{},
	)

	testCases := []struct {
		Name               string
		Method             string
		Header             map[string][]string
		BodyParam          dto.DoctorReq
		ExpectedStatusCode int
		ExpectedBody       dto.DoctorRes
	}{
		{
			"Success",
			http.MethodPost,
			map[string][]string{
				"Content-Type":    {"application/json"},
				"Accept":          {"*/*"},
				"Accept-Encoding": {"gzip", "deflate", "br"},
			},
			DoctorReq,
			http.StatusOK,
			Doctor,
		},
		{
			"Failed",
			http.MethodPost,
			map[string][]string{
				"Content-Type":    {"application/json"},
				"Accept":          {"*/*"},
				"Accept-Encoding": {"gzip", "deflate", "br"},
			},
			dto.DoctorReq{},
			http.StatusBadRequest,
			Doctor,
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
				var data dto.DoctorRes

				jsonStr, errMarsh := json.Marshal(response.Data)
				s.NoError(errMarsh)

				errUnMarsh := json.Unmarshal(jsonStr, &data)
				s.NoError(errUnMarsh)

				s.Equal(testCase.ExpectedBody, data)
			}
		})
	}
}
func (s *doctorTestSuite) TestUpdate() {
	s.mock.On("Update", uint(1), DoctorReq).Return(
		true,
		Doctor,
	)

	s.mock.On("Update", uint(0), dto.DoctorReq{}).Return(
		false,
		dto.DoctorRes{},
	)

	testCases := []struct {
		Name               string
		Method             string
		Header             map[string][]string
		Id                 uint
		Path               string
		BodyParam          dto.DoctorReq
		ExpectedStatusCode int
		ExpectedBody       dto.DoctorRes
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
			"doctors/:id",
			DoctorReq,
			http.StatusOK,
			Doctor,
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
			"doctors/:id",
			dto.DoctorReq{},
			http.StatusBadRequest,
			Doctor,
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
				var data dto.DoctorRes

				jsonStr, errMarsh := json.Marshal(response.Data)
				s.NoError(errMarsh)

				errUnMarsh := json.Unmarshal(jsonStr, &data)
				s.NoError(errUnMarsh)

				s.Equal(testCase.ExpectedBody, data)
			}
		})
	}
}
func (s *doctorTestSuite) TestDelete() {
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
		ExpectedBody       dto.DoctorRes
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
			"doctors/:id",
			http.StatusOK,
			Doctor,
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
			"doctors/:id",
			http.StatusBadRequest,
			Doctor,
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
var Doctors = []dto.DoctorRes{
	{
		ID:                  1,
		CreatedAt:           time.Time{},
		UpdatedAt:           time.Time{},
		DeletedAt:           gorm.DeletedAt{},
		Name:                "Dr. Fulan",
		SpecialityId:        1,
		LicenseNumber:       "1234567890",
		SpecialityName:      "Umum",
		ProfilePic:          "",
		BirthDate:           time.Time{},
		BirthDateString:     "",
		BirthDateStringIndo: "",
		Phone:               "081234567890",
		MaritalStatus:       false,
		Email:               "fulan@mail.com",
		DoctorSchedules:     nil,
	},
	{
		ID:                  2,
		CreatedAt:           time.Time{},
		UpdatedAt:           time.Time{},
		DeletedAt:           gorm.DeletedAt{},
		Name:                "Dr. Mulan",
		SpecialityId:        1,
		LicenseNumber:       "1234567891",
		SpecialityName:      "Umum",
		ProfilePic:          "",
		BirthDate:           time.Time{},
		BirthDateString:     "",
		BirthDateStringIndo: "",
		Phone:               "081234567891",
		MaritalStatus:       false,
		Email:               "mulan@mail.com",
		DoctorSchedules:     nil,
	},
}
var Doctor = dto.DoctorRes{
	ID:                  1,
	CreatedAt:           time.Time{},
	UpdatedAt:           time.Time{},
	DeletedAt:           gorm.DeletedAt{},
	Name:                "Dr. Fulan",
	SpecialityId:        1,
	LicenseNumber:       "1234567890",
	SpecialityName:      "Umum",
	ProfilePic:          "",
	BirthDate:           time.Time{},
	BirthDateString:     "",
	BirthDateStringIndo: "",
	Phone:               "081234567890",
	MaritalStatus:       false,
	Email:               "fulan@mail.com",
	DoctorSchedules:     nil,
}
var DoctorReq = dto.DoctorReq{
	Name:          "Dr. Fulan",
	LicenseNumber: "1234567890",
	Password:      "fulan123",
	SpecialityID:  1,
	ProfilePic:    "",
	BirthDate:     "1998-01-02",
	Phone:         "081234567890",
	MaritalStatus: false,
	Email:         "fulan@mail.com",
}
