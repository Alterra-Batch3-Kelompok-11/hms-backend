package dashboardController

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"hms-backend/dto"
	"hms-backend/usecases/dashboardUseCase/mock"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

type dashboardTestSuite struct {
	suite.Suite
	ctrl *dashboardController
	mock *mock.DashboardUseCaseMock
}

func (s *dashboardTestSuite) SetupSuite() {
	mock := &mock.DashboardUseCaseMock{}
	s.mock = mock

	s.ctrl = &dashboardController{
		usecase: s.mock,
	}
}

func TestOrderSuite(t *testing.T) {
	suite.Run(t, new(dashboardTestSuite))
}

func (s *dashboardTestSuite) TestGetDataDashboardWeb() {
	s.mock.On("GetDataDashboardWeb").Return(
		true,
		DataWeb,
	)

	testCases := []struct {
		Name               string
		Method             string
		Header             map[string][]string
		ExpectedStatusCode int
		ExpectedBody       dto.DashboardWeb
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
			DataWeb,
		},
	}

	for _, testCase := range testCases {
		s.T().Run(testCase.Name, func(t *testing.T) {
			req := httptest.NewRequest(testCase.Method, "/", nil)
			rec := httptest.NewRecorder()
			req.Header = testCase.Header

			e := echo.New()
			ctx := e.NewContext(req, rec)

			err := s.ctrl.GetDataDashboardWeb(ctx)
			s.NoError(err)
			s.Equal(testCase.ExpectedStatusCode, rec.Result().StatusCode)

			var response dto.Response
			errBdy := json.NewDecoder(rec.Result().Body).Decode(&response)
			s.NoError(errBdy)

			if response.Data != nil {
				var data dto.DashboardWeb

				jsonStr, errMarsh := json.Marshal(response.Data)
				s.NoError(errMarsh)

				errUnMarsh := json.Unmarshal(jsonStr, &data)
				s.NoError(errUnMarsh)

				s.Equal(testCase.ExpectedBody, data)
			}
		})
	}
}
func (s *dashboardTestSuite) TestGetDataDashboardMobile() {
	s.mock.On("GetDataDashboardMobile", uint(1)).Return(
		true,
		DataMobile,
	)

	s.mock.On("GetDataDashboardMobile", uint(2)).Return(
		false,
		dto.DashboardMobile{},
	)

	testCases := []struct {
		Name               string
		Method             string
		Header             map[string][]string
		Id                 uint
		Path               string
		ExpectedStatusCode int
		ExpectedBody       dto.DashboardMobile
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
			"dashboard/mobile/doctor/:doctor_id",
			http.StatusOK,
			DataMobile,
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
			"dashboard/mobile/doctor/:doctor_id",
			http.StatusBadRequest,
			DataMobile,
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

			err := s.ctrl.GetDataDashboardMobile(ctx)
			s.NoError(err)
			s.Equal(testCase.ExpectedStatusCode, rec.Result().StatusCode)

			var response dto.Response
			errBdy := json.NewDecoder(rec.Result().Body).Decode(&response)
			s.NoError(errBdy)

			if response.Data != nil {
				var data dto.DashboardMobile

				jsonStr, errMarsh := json.Marshal(response.Data)
				s.NoError(errMarsh)

				errUnMarsh := json.Unmarshal(jsonStr, &data)
				s.NoError(errUnMarsh)

				s.Equal(testCase.ExpectedBody, data)
			} else {
				s.Equal("testing error", response.Message)
			}
		})
	}
}

var DataWeb = dto.DashboardWeb{
	TotalDoctors:  2,
	TotalNurses:   1,
	TotalPatients: 1,
	TodayDoctors: []dto.TodayDoctorRes{
		{
			Name:           "Dr. Fulan",
			LicenseNumber:  "1234567890",
			SpecialityName: "Umum",
			ProfilePic:     "",
			DayInt:         1,
			DayString:      "Senin",
			StartTime:      "08:00",
			EndTime:        "16:00",
		},
	},
	TodayOutpatientSessions: []dto.OutpatientSessionDashboardRes{
		{
			Patient: struct {
				NIK           string    `json:"nik"`
				Name          string    `json:"name"`
				BirthDate     time.Time `json:"birth_date"`
				Gender        int       `json:"gender"`
				Age           int       `json:"age"`
				Phone         string    `json:"phone"`
				Address       string    `json:"address"`
				MaritalStatus bool      `json:"marital_status"`
				ReligionName  string    `json:"religion_name"`
			}{
				NIK:           "1234567890",
				Name:          "Ahmad",
				BirthDate:     time.Time{},
				Gender:        1,
				Age:           20,
				Phone:         "081234567890",
				Address:       "Surabaya",
				MaritalStatus: false,
				ReligionName:  "Islam",
			},
			Doctor: struct {
				Name           string `json:"name"`
				LicenseNumber  string `json:"license_number" form:"license_number"`
				SpecialityName string `json:"speciality_name" form:"speciality_name"`
			}{
				Name:           "Dr. Fulan",
				LicenseNumber:  "1234567890",
				SpecialityName: "Umum",
			},
			Complaint:      "Batuk Berdahak",
			IsApproved:     1,
			IsFinish:       false,
			FinishedAt:     time.Time{},
			FinishedAtDate: "",
			FinishedAtTime: "",
			Schedule:       time.Time{},
			ScheduleDate:   "",
			ScheduleTime:   "",
		},
	},
	Patients: nil,
}

var DataMobile = dto.DashboardMobile{
	TotalQueueToday:    1,
	TotalFinishedToday: 10,
	PatientsToday: []dto.PatientToday{
		{
			Name:             "Ahmad",
			ScheduleDate:     "2022-12-18",
			ScheduleDateIndo: "18 Desember 2022",
			ScheduleTime:     "10:00",
			Schedule:         time.Time{},
			Complaint:        "Batuk Berdahak",
		},
	},
}
