package roleController

import (
	"encoding/json"
	"hms-backend/dto"
	"hms-backend/usecases/roleUseCase/mock"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type roleTestSuite struct {
	suite.Suite
	ctrl *roleController
	mock *mock.RoleUseCaseMock
}

func (s *roleTestSuite) SetupSuite() {
	mock := &mock.RoleUseCaseMock{}
	s.mock = mock

	s.ctrl = &roleController{
		usecase: s.mock,
	}
}

func TestOrderSuite(t *testing.T) {
	suite.Run(t, new(roleTestSuite))
}

func (s *roleTestSuite) TestGetAll() {
	s.mock.On("GetAll").Return(
		true,
		Roles,
	)

	testCases := []struct {
		Name               string
		Method             string
		Header             map[string][]string
		ExpectedStatusCode int
		ExpectedBody       []dto.Role
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
			Roles,
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
				var data []dto.Role

				jsonStr, errMarsh := json.Marshal(response.Data.([]interface{}))
				s.NoError(errMarsh)

				errUnMarsh := json.Unmarshal(jsonStr, &data)
				s.NoError(errUnMarsh)

				s.Equal(testCase.ExpectedBody, data)
			}
		})
	}
}

func (s *roleTestSuite) TestGetById() {
	s.mock.On("GetById", uint(1)).Return(
		true,
		Role,
	)

	s.mock.On("GetById", uint(2)).Return(
		false,
		dto.Role{},
	)

	testCases := []struct {
		Name               string
		Method             string
		Header             map[string][]string
		Id                 uint
		Path               string
		ExpectedStatusCode int
		ExpectedBody       dto.Role
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
			"roless/:id",
			http.StatusOK,
			Role,
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
			"roles/:id",
			http.StatusBadRequest,
			dto.Role{},
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
				var data dto.Role

				jsonStr, errMarsh := json.Marshal(response.Data)
				s.NoError(errMarsh)

				errUnMarsh := json.Unmarshal(jsonStr, &data)
				s.NoError(errUnMarsh)

				s.Equal(testCase.ExpectedBody, data)
			}
		})
	}
}

// Data mock

var Roles = []dto.Role{
	{
		ID:        1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
		Name:      "admin",
	},
	{
		ID:        2,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
		Name:      "doctor",
	},
}

var Role = dto.Role{
	ID:        1,
	CreatedAt: time.Time{},
	UpdatedAt: time.Time{},
	DeletedAt: gorm.DeletedAt{},
	Name:      "admin",
}
