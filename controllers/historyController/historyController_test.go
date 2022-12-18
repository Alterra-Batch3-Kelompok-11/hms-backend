package historyController

import (
	"github.com/stretchr/testify/suite"
	"hms-backend/usecases/historyUseCase/mock"
	"testing"
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

func (s *dashboardTestSuite) GetOutpatientSessionHistory() {

}
func (s *dashboardTestSuite) GetApprovalHistory() {

}

/*
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
*/
