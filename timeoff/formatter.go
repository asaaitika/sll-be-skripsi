package timeoff

import (
	"time"
)

type TimeOffFormatter struct {
	TimeOffId     int       `json:"timeoff_id"`
	TimeOffType   string    `json:"timeoff_type"`
	TimeOffSaldo  string    `json:"timeoff_saldo"`
	StatusTimeOff string    `json:"status_timeoff"`
	StartDate     time.Time `json:"start_end"`
	EndDate       time.Time `json:"end_date"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	RequestType   string    `json:"request_type"`
	Reason        string    `json:"reason"`
	File          string    `json:"file"`
	Remark        string    `json:"remark"`
	EmployeeId    int       `json:"employee_id"`
	EmployeeName  string    `json:"employee_name"`
	RoleName      string    `json:"role_name"`
}

func FormatTimeOff(timeoff TimeOff) TimeOffFormatter {
	formatter := TimeOffFormatter{}
	formatter.TimeOffId = timeoff.TimeoffId
	formatter.TimeOffType = timeoff.TimeoffType
	formatter.TimeOffSaldo = timeoff.TimeoffSaldo
	formatter.StatusTimeOff = timeoff.StatusTimeoff
	formatter.StartDate = timeoff.StartDate
	formatter.EndDate = timeoff.EndDate
	formatter.CreatedAt = timeoff.CreatedAt
	formatter.UpdatedAt = timeoff.UpdatedAt
	formatter.RequestType = timeoff.RequestType
	formatter.Reason = timeoff.Reason
	formatter.File = timeoff.File
	formatter.EmployeeId = timeoff.EmployeeId
	formatter.EmployeeName = timeoff.Employee.EmployeeName
	formatter.RoleName = timeoff.Role.RoleName

	return formatter
}

func FormatTimeOffs(timeoffs []TimeOff) []TimeOffFormatter {
	timeoffsFormatter := []TimeOffFormatter{}

	for _, timeoff := range timeoffs {
		timeoffFormatter := FormatTimeOff(timeoff)
		timeoffsFormatter = append(timeoffsFormatter, timeoffFormatter)
	}

	return timeoffsFormatter
}
