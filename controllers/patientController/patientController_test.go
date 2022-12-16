package patientController

import (
	"bytes"
	"encoding/json"
	"hms-backend/dto"
	"hms-backend/usecases/patientUseCase/mock"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type patientTestSuite struct {
	suite.Suite
	ctrl *patientController
	mock *mock.PatientUseCaseMock
}

func (s *patientTestSuite) SetupSuite() {
	mock := &mock.PatientUseCaseMock{}
	s.mock = mock

	s.ctrl = &patientController{
		usecase: s.mock,
	}
}

func TestOrderSuite(t *testing.T) {
	suite.Run(t, new(patientTestSuite))
}

func (s *patientTestSuite) TestGetAll() {
	s.mock.On("GetAll").Return(
		true,
		Patients,
	)

	testCases := []struct {
		Name               string
		Method             string
		Header             map[string][]string
		ExpectedStatusCode int
		ExpectedBody       []dto.PatientRes
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
			Patients,
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
				var data []dto.PatientRes

				jsonStr, errMarsh := json.Marshal(response.Data.([]interface{}))
				s.NoError(errMarsh)

				errUnMarsh := json.Unmarshal(jsonStr, &data)
				s.NoError(errUnMarsh)

				s.Equal(testCase.ExpectedBody, data)
			}
		})
	}
}

func (s *patientTestSuite) TestGetById() {
	s.mock.On("GetById", uint(1)).Return(
		true,
		Patient,
	)

	s.mock.On("GetById", uint(2)).Return(
		false,
		dto.PatientRes{},
	)

	testCases := []struct {
		Name               string
		Method             string
		Header             map[string][]string
		Id                 uint
		Path               string
		ExpectedStatusCode int
		ExpectedBody       dto.PatientRes
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
			"patients/:id",
			http.StatusOK,
			Patient,
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
			"patients/:id",
			http.StatusBadRequest,
			dto.PatientRes{},
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
				var data dto.PatientRes

				jsonStr, errMarsh := json.Marshal(response.Data)
				s.NoError(errMarsh)

				errUnMarsh := json.Unmarshal(jsonStr, &data)
				s.NoError(errUnMarsh)

				s.Equal(testCase.ExpectedBody, data)
			}
		})
	}
}

func (s *patientTestSuite) TestCreate() {
	s.mock.On("Create", PatientReq).Return(
		true,
		Patient,
	)

	s.mock.On("Create", dto.Patient{}).Return(
		false,
		dto.PatientRes{},
	)

	testCases := []struct {
		Name               string
		Method             string
		Header             map[string][]string
		BodyParam          dto.Patient
		ExpectedStatusCode int
		ExpectedBody       dto.PatientRes
	}{
		{
			"Success",
			http.MethodPost,
			map[string][]string{
				"Content-Type":    {"application/json"},
				"Accept":          {"*/*"},
				"Accept-Encoding": {"gzip", "deflate", "br"},
			},
			PatientReq,
			http.StatusOK,
			Patient,
		},
		{
			"Failed",
			http.MethodPost,
			map[string][]string{
				"Content-Type":    {"application/json"},
				"Accept":          {"*/*"},
				"Accept-Encoding": {"gzip", "deflate", "br"},
			},
			dto.Patient{},
			http.StatusBadRequest,
			Patient,
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
				var data dto.PatientRes

				jsonStr, errMarsh := json.Marshal(response.Data)
				s.NoError(errMarsh)

				errUnMarsh := json.Unmarshal(jsonStr, &data)
				s.NoError(errUnMarsh)

				s.Equal(testCase.ExpectedBody, data)
			}
		})
	}
}

func (s *patientTestSuite) TestUpdate() {
	s.mock.On("Update", uint(1), PatientReq).Return(
		true,
		Patient,
	)

	s.mock.On("Update", uint(0), dto.Patient{}).Return(
		false,
		dto.PatientRes{},
	)

	testCases := []struct {
		Name               string
		Method             string
		Header             map[string][]string
		Id                 uint
		Path               string
		BodyParam          dto.Patient
		ExpectedStatusCode int
		ExpectedBody       dto.PatientRes
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
			"patients/:id",
			PatientReq,
			http.StatusOK,
			Patient,
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
			"patients/:id",
			dto.Patient{},
			http.StatusBadRequest,
			Patient,
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
				var data dto.PatientRes

				jsonStr, errMarsh := json.Marshal(response.Data)
				s.NoError(errMarsh)

				errUnMarsh := json.Unmarshal(jsonStr, &data)
				s.NoError(errUnMarsh)

				s.Equal(testCase.ExpectedBody, data)
			}
		})
	}
}

func (s *patientTestSuite) TestDelete() {
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
		ExpectedBody       dto.PatientRes
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
			"patients/:id",
			http.StatusOK,
			Patient,
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
			"patients/:id",
			http.StatusBadRequest,
			Patient,
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

var Patients = []dto.PatientRes{
	{
		ID:            1,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     gorm.DeletedAt{},
		Nik:           "1234567890",
		Name:          "Patient 1",
		BirthDate:     time.Time{},
		Gender:        1,
		Address:       "Address 1",
		Phone:         "081234567890",
		MaritalStatus: false,
		ReligionID:    2,
	},
	{
		ID:            2,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     gorm.DeletedAt{},
		Nik:           "12121234567890",
		Name:          "Patient 2",
		BirthDate:     time.Time{},
		Gender:        1,
		Address:       "Address 2",
		Phone:         "0812345678121290",
		MaritalStatus: false,
		ReligionID:    1,
	},
}

var Patient = dto.PatientRes{
	ID:            1,
	CreatedAt:     time.Time{},
	UpdatedAt:     time.Time{},
	DeletedAt:     gorm.DeletedAt{},
	Nik:           "1234567890",
	Name:          "Patient 1",
	BirthDate:     time.Time{},
	Gender:        1,
	Address:       "Address 1",
	Phone:         "081234567890",
	MaritalStatus: false,
	ReligionID:    2,
}

var PatientReq = dto.Patient{
	NIK:           "1234567890",
	Name:          "Patient 1",
	BirthDate:     "2021-01-01",
	Gender:        1,
	Phone:         "081234567890",
	Address:       "Address 1",
	MaritalStatus: false,
	ReligionID:    1,
}
