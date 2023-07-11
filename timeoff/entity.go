package timeoff

import (
	"time"
)

type TimeOff struct {
	TimeoffId     int
	TimeoffType   string
	TimeoffSaldo  string
	StatusTimeoff string
	StartDate     time.Time
	EndDate       time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
	RequestType   string
	Reason        string
	File          string
	Remark        string
	EmployeeId    int
	EmployeeName  string
	RoleName      string
}
