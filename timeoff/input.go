package timeoff

import "time"

type CreateRequestTimeOffInput struct {
	TimeOffType   string    `form:"timeoff_type" binding:"required"`
	TimeOffSaldo  string    `form:"timeoff_saldo" binding:"required"`
	StartDate     time.Time `form:"start_date" binding:"required"`
	EndDate       time.Time `form:"end_date"`
	RequestType   string    `form:"request_type" binding:"required"`
	Reason        string    `form:"reason"`
	StatusTimeOff string
	EmployeeId    int
}

type UpdateStatusTimeOffInput struct {
	StatusTimeoff string `form:"status_timeoff" binding:"required"`
	Remark        string `form:"remark"`
}

type SearchTimeOffInput struct {
	EmployeeId int    `form:"employee_id"`
	Status     string `form:"status"`
	Month      string `form:"month"`
	Year       string `form:"year"`
}

type SearchRequestTimeOffInput struct {
	Keyword     string `form:"keyword"`
	RequestDate string `form:"request_date"`
	Status      string `form:"status"`
}

type GetTimeOffDetailInput struct {
	Id int `uri:"id" binding:"required"`
}
