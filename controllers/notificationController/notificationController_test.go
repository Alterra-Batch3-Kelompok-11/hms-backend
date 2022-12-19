package notificationController

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"hms-backend/dto"
	"hms-backend/middlewares"
	"hms-backend/usecases/notificationUseCase/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

type notificationTestSuite struct {
	suite.Suite
	ctrl *notificationController
	mock *mock.NotificationUseCaseMock
}

func (s *notificationTestSuite) SetupSuite() {
	mock := &mock.NotificationUseCaseMock{}
	s.mock = mock

	s.ctrl = &notificationController{
		usecase: s.mock,
	}
}

func TestOrderSuite(t *testing.T) {
	suite.Run(t, new(notificationTestSuite))
}

func (s *notificationTestSuite) TestGetByDoctorId() {
	s.mock.On("GetByUserId", uint(1)).Return(
		true,
		Notifications,
	)

	s.mock.On("GetByUserId", uint(2)).Return(
		false,
		[]dto.Notification{},
	)

	testCases := []struct {
		Name               string
		Method             string
		Header             map[string][]string
		IsSuccess          bool
		ExpectedMessage    string
		ExpectedStatusCode int
		ExpectedBody       []dto.Notification
	}{
		{
			"Success",
			http.MethodGet,
			map[string][]string{
				"Content-Type":    {"application/json"},
				"Accept":          {"*/*"},
				"Accept-Encoding": {"gzip", "deflate", "br"},
				"Authorization":   {"Bearer " + TokenUserId1},
			},
			true,
			"",
			http.StatusOK,
			Notifications,
		},
		{
			"Failed",
			http.MethodGet,
			map[string][]string{
				"Content-Type":    {"application/json"},
				"Accept":          {"*/*"},
				"Accept-Encoding": {"gzip", "deflate", "br"},
				"Authorization":   {"Bearer " + TokenUserId2},
			},
			false,
			"testing error",
			http.StatusBadRequest,
			Notifications,
		},
	}

	for _, testCase := range testCases {
		s.T().Run(testCase.Name, func(t *testing.T) {
			req := httptest.NewRequest(testCase.Method, "/", nil)
			rec := httptest.NewRecorder()
			req.Header = testCase.Header

			e := echo.New()
			ctx := e.NewContext(req, rec)

			err := s.ctrl.GetByDoctorId(ctx)

			s.Equal(testCase.ExpectedStatusCode, rec.Result().StatusCode)

			var response dto.Response
			errBdy := json.NewDecoder(rec.Result().Body).Decode(&response)
			s.NoError(errBdy)

			if response.Data != nil {
				s.NoError(err)

				var data []dto.Notification

				jsonStr, errMarsh := json.Marshal(response.Data)
				s.NoError(errMarsh)

				errUnMarsh := json.Unmarshal(jsonStr, &data)
				s.NoError(errUnMarsh)

				s.Equal(testCase.ExpectedBody, data)
			}
		})
	}
}

var TokenUserId1, _ = middlewares.CreateToken(uint(1), "1234567890", 2)
var TokenUserId2, _ = middlewares.CreateToken(uint(2), "1234567891", 2)

var Notifications = []dto.Notification{
	{
		OutpatientSessionID: 1,
		Description:         "Lorem ipsum",
		DateString:          "2022-12-30",
		DateStringIndo:      "30 Desember 2022",
		TimeString:          "08:00",
	},
}
