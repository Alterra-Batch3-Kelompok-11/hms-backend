package specialityController

import (
	"bytes"
	"encoding/json"
	"hms-backend/dto"
	"hms-backend/usecases/specialityUseCase/mock"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type specialityTestSuite struct {
	suite.Suite
	ctrl *specialityController
	mock *mock.SpecialityUseCaseMock
}

func (s *specialityTestSuite) SetupSuite() {
	mock := &mock.SpecialityUseCaseMock{}
	s.mock = mock

	s.ctrl = &specialityController{
		usecase: s.mock,
	}
}

func TestOrderSuite(t *testing.T) {
	suite.Run(t, new(specialityTestSuite))
}

func (s *specialityTestSuite) TestGetAll() {
	s.mock.On("GetAll").Return(
		true,
		Specialities,
	)

	testCases := []struct {
		Name               string
		Method             string
		Header             map[string][]string
		ExpectedStatusCode int
		ExpectedBody       []dto.Speciality
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
			Specialities,
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
				var data []dto.Speciality

				jsonStr, errMarsh := json.Marshal(response.Data.([]interface{}))
				s.NoError(errMarsh)

				errUnMarsh := json.Unmarshal(jsonStr, &data)
				s.NoError(errUnMarsh)

				s.Equal(testCase.ExpectedBody, data)
			}
		})
	}
}

func (s *specialityTestSuite) TestGetById() {
	s.mock.On("GetById", uint(1)).Return(
		true,
		Speciality,
	)

	s.mock.On("GetById", uint(2)).Return(
		false,
		dto.Speciality{},
	)

	testCases := []struct {
		Name               string
		Method             string
		Header             map[string][]string
		Id                 uint
		Path               string
		ExpectedStatusCode int
		ExpectedBody       dto.Speciality
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
			"specialities/:id",
			http.StatusOK,
			Speciality,
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
			"specialities/:id",
			http.StatusBadRequest,
			dto.Speciality{},
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
				var data dto.Speciality

				jsonStr, errMarsh := json.Marshal(response.Data)
				s.NoError(errMarsh)

				errUnMarsh := json.Unmarshal(jsonStr, &data)
				s.NoError(errUnMarsh)

				s.Equal(testCase.ExpectedBody, data)
			}
		})
	}
}

func (s *specialityTestSuite) TestCreate() {
	s.mock.On("Create", Speciality).Return(
		true,
		Speciality,
	)

	s.mock.On("Create", dto.Speciality{}).Return(
		false,
		dto.Speciality{},
	)

	testCases := []struct {
		Name               string
		Method             string
		Header             map[string][]string
		BodyParam          dto.Speciality
		ExpectedStatusCode int
		ExpectedBody       dto.Speciality
	}{
		{
			"Success",
			http.MethodPost,
			map[string][]string{
				"Content-Type":    {"application/json"},
				"Accept":          {"*/*"},
				"Accept-Encoding": {"gzip", "deflate", "br"},
			},
			Speciality,
			http.StatusOK,
			Speciality,
		},
		{
			"Failed",
			http.MethodPost,
			map[string][]string{
				"Content-Type":    {"application/json"},
				"Accept":          {"*/*"},
				"Accept-Encoding": {"gzip", "deflate", "br"},
			},
			dto.Speciality{},
			http.StatusBadRequest,
			Speciality,
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
				var data dto.Speciality

				jsonStr, errMarsh := json.Marshal(response.Data)
				s.NoError(errMarsh)

				errUnMarsh := json.Unmarshal(jsonStr, &data)
				s.NoError(errUnMarsh)

				s.Equal(testCase.ExpectedBody, data)
			}
		})
	}
}

func (s *specialityTestSuite) TestUpdate() {
	s.mock.On("Update", uint(1), Speciality).Return(
		true,
		Speciality,
	)

	s.mock.On("Update", uint(0), dto.Speciality{}).Return(
		false,
		dto.Speciality{},
	)

	testCases := []struct {
		Name               string
		Method             string
		Header             map[string][]string
		Id                 uint
		Path               string
		BodyParam          dto.Speciality
		ExpectedStatusCode int
		ExpectedBody       dto.Speciality
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
			"specialities/:id",
			Speciality,
			http.StatusOK,
			Speciality,
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
			"specialities/:id",
			dto.Speciality{},
			http.StatusBadRequest,
			Speciality,
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
				var data dto.Speciality

				jsonStr, errMarsh := json.Marshal(response.Data)
				s.NoError(errMarsh)

				errUnMarsh := json.Unmarshal(jsonStr, &data)
				s.NoError(errUnMarsh)

				s.Equal(testCase.ExpectedBody, data)
			}
		})
	}
}

func (s *specialityTestSuite) TestDelete() {
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
		ExpectedBody       dto.Speciality
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
			"specialities/:id",
			http.StatusOK,
			Speciality,
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
			"specialities/:id",
			http.StatusBadRequest,
			Speciality,
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

var Specialities = []dto.Speciality{
	{
		ID:        1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
		Name:      "Speciality 1",
	},
	{
		ID:        2,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
		Name:      "Speciality 2",
	},
}

var Speciality = dto.Speciality{
	ID:        2,
	CreatedAt: time.Time{},
	UpdatedAt: time.Time{},
	DeletedAt: gorm.DeletedAt{},
	Name:      "Speciality 2",
}
