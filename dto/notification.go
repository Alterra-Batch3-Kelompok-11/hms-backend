package dto

type Notification struct {
	OutpatientSessionID uint   `json:"outpatient_session_id" form:"outpatient_session_id"`
	Description         string `json:"description" form:"description"`
	DateString          string `json:"schedule_date" form:"schedule_date"`
	DateStringIndo      string `json:"schedule_date_indo" form:"schedule_date_indo"`
	TimeString          string `json:"schedule_time" form:"schedule_time"`
}
