package nurseController

import (
	"bytes"
	"encoding/json"
	"hms-backend/dto"
	"hms-backend/usecases/nurseUseCase/mock"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type nurseTestSuite struct {
	suite.Suite
	ctrl *nurseController
	mock *mock.NurseUseCaseMock
}

func (s *nurseTestSuite) SetupSuite() {
	mock := &mock.NurseUseCaseMock{}
	s.mock = mock

	s.ctrl = &nurseController{
		usecase: s.mock,
	}
}

func TestOrderSuite(t *testing.T) {
	suite.Run(t, new(nurseTestSuite))
}

func (s *nurseTestSuite) TestGetAll() {
	s.mock.On("GetAll").Return(
		true,
		Nurses,
	)

	testCases := []struct {
		Name               string
		Method             string
		Header             map[string][]string
		ExpectedStatusCode int
		ExpectedBody       []dto.NurseRes
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
			Nurses,
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
				var data []dto.NurseRes

				jsonStr, errMarsh := json.Marshal(response.Data.([]interface{}))
				s.NoError(errMarsh)

				errUnMarsh := json.Unmarshal(jsonStr, &data)
				s.NoError(errUnMarsh)

				s.Equal(testCase.ExpectedBody, data)
			}
		})
	}
}

func (s *nurseTestSuite) TestGetById() {
	s.mock.On("GetById", uint(1)).Return(
		true,
		Nurse,
	)

	s.mock.On("GetById", uint(2)).Return(
		false,
		dto.NurseRes{},
	)

	testCases := []struct {
		Name               string
		Method             string
		Header             map[string][]string
		Id                 uint
		Path               string
		ExpectedStatusCode int
		ExpectedBody       dto.NurseRes
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
			"nurses/:id",
			http.StatusOK,
			Nurse,
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
			"nurses/:id",
			http.StatusBadRequest,
			dto.NurseRes{},
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
				var data dto.NurseRes

				jsonStr, errMarsh := json.Marshal(response.Data)
				s.NoError(errMarsh)

				errUnMarsh := json.Unmarshal(jsonStr, &data)
				s.NoError(errUnMarsh)

				s.Equal(testCase.ExpectedBody, data)
			}
		})
	}
}

func (s *nurseTestSuite) TestGetByLicenseNumber() {
	s.mock.On("GetByLicenseNumber", "1234567890").Return(
		true,
		Nurse,
	)

	s.mock.On("GetByLicenseNumber", "").Return(
		false,
		dto.NurseRes{},
	)

	testCases := []struct {
		Name               string
		Method             string
		Header             map[string][]string
		LicenseNumber      string
		Path               string
		ExpectedStatusCode int
		ExpectedBody       dto.NurseRes
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
			"nurses/license_number/:license_number",
			http.StatusOK,
			Nurse,
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
			"nurses/:id",
			http.StatusBadRequest,
			dto.NurseRes{},
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
				var data dto.NurseRes

				jsonStr, errMarsh := json.Marshal(response.Data)
				s.NoError(errMarsh)

				errUnMarsh := json.Unmarshal(jsonStr, &data)
				s.NoError(errUnMarsh)

				s.Equal(testCase.ExpectedBody, data)
			}
		})
	}
}

func (s *nurseTestSuite) TestCreate() {
	s.mock.On("Create", NurseReq).Return(
		true,
		Nurse,
	)

	s.mock.On("Create", dto.NurseReq{}).Return(
		false,
		dto.NurseRes{},
	)

	testCases := []struct {
		Name               string
		Method             string
		Header             map[string][]string
		BodyParam          dto.NurseReq
		ExpectedStatusCode int
		ExpectedBody       dto.NurseRes
	}{
		{
			"Success",
			http.MethodPost,
			map[string][]string{
				"Content-Type":    {"application/json"},
				"Accept":          {"*/*"},
				"Accept-Encoding": {"gzip", "deflate", "br"},
			},
			NurseReq,
			http.StatusOK,
			Nurse,
		},
		{
			"Failed",
			http.MethodPost,
			map[string][]string{
				"Content-Type":    {"application/json"},
				"Accept":          {"*/*"},
				"Accept-Encoding": {"gzip", "deflate", "br"},
			},
			dto.NurseReq{},
			http.StatusBadRequest,
			Nurse,
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
				var data dto.NurseRes

				jsonStr, errMarsh := json.Marshal(response.Data)
				s.NoError(errMarsh)

				errUnMarsh := json.Unmarshal(jsonStr, &data)
				s.NoError(errUnMarsh)

				s.Equal(testCase.ExpectedBody, data)
			}
		})
	}
}

func (s *nurseTestSuite) TestUpdate() {
	s.mock.On("Update", uint(1), NurseReq).Return(
		true,
		Nurse,
	)

	s.mock.On("Update", uint(0), dto.NurseReq{}).Return(
		false,
		dto.NurseRes{},
	)

	testCases := []struct {
		Name               string
		Method             string
		Header             map[string][]string
		Id                 uint
		Path               string
		BodyParam          dto.NurseReq
		ExpectedStatusCode int
		ExpectedBody       dto.NurseRes
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
			"nurses/:id",
			NurseReq,
			http.StatusOK,
			Nurse,
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
			dto.NurseReq{},
			http.StatusBadRequest,
			Nurse,
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
				var data dto.NurseRes

				jsonStr, errMarsh := json.Marshal(response.Data)
				s.NoError(errMarsh)

				errUnMarsh := json.Unmarshal(jsonStr, &data)
				s.NoError(errUnMarsh)

				s.Equal(testCase.ExpectedBody, data)
			}
		})
	}
}

func (s *nurseTestSuite) TestDelete() {
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
		ExpectedBody       dto.NurseRes
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
			"nurses/:id",
			http.StatusOK,
			Nurse,
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
			"nurses/:id",
			http.StatusBadRequest,
			Nurse,
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

// Data mock

var Nurses = []dto.NurseRes{
	{
		ID:                  1,
		UserID:              1,
		DoctorID:            1,
		Name:                "Nurse 1",
		SpecialityId:        1,
		LicenseNumber:       "1234567890",
		SpecialityName:      "Speciality 1",
		BirthDate:           time.Time{},
		BirthDateString:     "",
		BirthDateStringIndo: "",
		ProfilePic:          "",
		Phone:               "121212",
		MaritalStatus:       false,
		Email:               "nurse@gmail.com",
		Nira:                "1234567890",
		SIP:                 "1234567890",
		CreatedAt:           time.Time{},
		UpdatedAt:           time.Time{},
		DeletedAt:           gorm.DeletedAt{},
	},
	{
		ID:                  2,
		UserID:              2,
		DoctorID:            2,
		Name:                "Nurse 2",
		SpecialityId:        2,
		LicenseNumber:       "12345678901212",
		SpecialityName:      "Speciality 2",
		BirthDate:           time.Time{},
		BirthDateString:     "",
		BirthDateStringIndo: "",
		ProfilePic:          "",
		Phone:               "1212121212",
		MaritalStatus:       false,
		Email:               "nurse2@gmail.com",
		Nira:                "12345678901212",
		SIP:                 "12345678901212",
		CreatedAt:           time.Time{},
		UpdatedAt:           time.Time{},
		DeletedAt:           gorm.DeletedAt{},
	},
}

var Nurse = dto.NurseRes{
	ID:                  1,
	UserID:              1,
	DoctorID:            1,
	Name:                "Nurse 1",
	SpecialityId:        1,
	LicenseNumber:       "1234567890",
	SpecialityName:      "Speciality 1",
	BirthDate:           time.Time{},
	BirthDateString:     "",
	BirthDateStringIndo: "",
	ProfilePic:          "",
	Phone:               "121212",
	MaritalStatus:       false,
	Email:               "nurse@gmail.com",
	Nira:                "1234567890",
	SIP:                 "1234567890",
	CreatedAt:           time.Time{},
	UpdatedAt:           time.Time{},
	DeletedAt:           gorm.DeletedAt{},
}

var NurseReq = dto.NurseReq{
	DoctorID:      1,
	Name:          "Nurse 1",
	LicenseNumber: "1234567890",
	Password:      "1234567890",
	SpecialityID:  1,
	ProfilePic:    "",
	BirthDate:     "2001-01-01",
	Phone:         "121212",
	MaritalStatus: false,
	Email:         "nurse1@gmail.com",
	Nira:          "1234567890",
	SIP:           "1234567890",
}
