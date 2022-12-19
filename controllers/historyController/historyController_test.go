package historyController

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"hms-backend/dto"
	"hms-backend/usecases/historyUseCase/mock"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

type dashboardTestSuite struct {
	suite.Suite
	ctrl *historyController
	mock *mock.HistoryUseCaseMock
}

func (s *dashboardTestSuite) SetupSuite() {
	mock := &mock.HistoryUseCaseMock{}
	s.mock = mock

	s.ctrl = &historyController{
		usecase: s.mock,
	}
}

func TestOrderSuite(t *testing.T) {
	suite.Run(t, new(dashboardTestSuite))
}

func (s *dashboardTestSuite) TestGetOutpatientSessionHistory() {
	s.mock.On("GetOutpatientSessionHistory", uint(1)).Return(
		true,
		Histories,
	)

	s.mock.On("GetOutpatientSessionHistory", uint(2)).Return(
		false,
		[]dto.History{},
	)

	testCases := []struct {
		Name               string
		Method             string
		Header             map[string][]string
		Id                 uint
		Path               string
		ExpectedStatusCode int
		ExpectedBody       []dto.History
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
			"/histories/doctor/:doctor_id/outpatient_sessions",
			http.StatusOK,
			Histories,
		},
		{
			"Failed",
			http.MethodGet,
			map[string][]string{
				"Content-Type":    {"application/json"},
				"Accept":          {"*/*"},
				"Accept-Encoding": {"gzip", "deflate", "br"},
			},
			2,
			"/histories/doctor/:doctor_id/outpatient_sessions",
			http.StatusBadRequest,
			Histories,
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
			ctx.SetParamValues(strconv.Itoa(int(testCase.Id)))
			ctx.SetPath(testCase.Path)

			err := s.ctrl.GetOutpatientSessionHistory(ctx)

			s.Equal(testCase.ExpectedStatusCode, rec.Result().StatusCode)

			var response dto.Response
			errBdy := json.NewDecoder(rec.Result().Body).Decode(&response)
			s.NoError(errBdy)

			if response.Data != nil {
				s.NoError(err)

				var data []dto.History

				jsonStr, errMarsh := json.Marshal(response.Data)
				s.NoError(errMarsh)

				errUnMarsh := json.Unmarshal(jsonStr, &data)
				s.NoError(errUnMarsh)

				s.Equal(testCase.ExpectedBody, data)
			}
		})
	}
}
func (s *dashboardTestSuite) TestGetApprovalHistory() {
	s.mock.On("GetApprovalHistory", uint(1)).Return(
		true,
		Histories,
	)

	s.mock.On("GetApprovalHistory", uint(2)).Return(
		false,
		[]dto.History{},
	)

	testCases := []struct {
		Name               string
		Method             string
		Header             map[string][]string
		Id                 uint
		Path               string
		ExpectedStatusCode int
		ExpectedBody       []dto.History
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
			"/histories/doctor/:doctor_id/approvals",
			http.StatusOK,
			Histories,
		},
		{
			"Failed",
			http.MethodGet,
			map[string][]string{
				"Content-Type":    {"application/json"},
				"Accept":          {"*/*"},
				"Accept-Encoding": {"gzip", "deflate", "br"},
			},
			2,
			"/histories/doctor/:doctor_id/approvals",
			http.StatusBadRequest,
			Histories,
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
			ctx.SetParamValues(strconv.Itoa(int(testCase.Id)))
			ctx.SetPath(testCase.Path)

			err := s.ctrl.GetApprovalHistory(ctx)

			s.Equal(testCase.ExpectedStatusCode, rec.Result().StatusCode)

			var response dto.Response
			errBdy := json.NewDecoder(rec.Result().Body).Decode(&response)
			s.NoError(errBdy)

			if response.Data != nil {
				s.NoError(err)

				var data []dto.History

				jsonStr, errMarsh := json.Marshal(response.Data)
				s.NoError(errMarsh)

				errUnMarsh := json.Unmarshal(jsonStr, &data)
				s.NoError(errUnMarsh)

				s.Equal(testCase.ExpectedBody, data)
			}
		})
	}
}

var Histories = []dto.History{
	{
		PatientName:      "Ahmad",
		Schedule:         time.Time{},
		ScheduleDate:     "2022-12-18",
		ScheduleDateIndo: "18 Desember 2022",
		ScheduleTime:     "08:00",
		Status:           "Proses",
	},
}
