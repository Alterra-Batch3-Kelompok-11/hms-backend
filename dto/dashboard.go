package dto

type DashboardWeb struct {
	TotalDoctors            int64                           `json:"total_doctors" form:"total_doctors"`
	TotalNurses             int64                           `json:"total_nurses" form:"total_nurses"`
	TotalPatients           int64                           `json:"total_patients" form:"total_patients"`
	TodayDoctors            []TodayDoctorRes                `json:"today_doctors" form:"today_doctors"`
	TodayOutpatientSessions []OutpatientSessionDashboardRes `json:"today_outpatient_sessions" form:"today_outpatient_sessions"`
	Patients                []OutpatientSessionDashboardRes `json:"patients" form:"patients"`
}
