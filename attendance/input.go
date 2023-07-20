package attendance

type CreateAttendanceInput struct {
	AttendaceType string `json:"attendance_type" binding:"required"`
	AttendaceDate string `json:"attendance_date" binding:"required"`
	ClockinTime   string `json:"clockin_time" binding:"required"`
	Longitude     string `json:"longitude" binding:"required"`
	Latitude      string `json:"latitude" binding:"required"`
}

type UpdateAttendanceInput struct {
	ClockoutTime string `json:"clockout_time" binding:"required"`
}

type SearchAttendanceLogInput struct {
	Id        int    `form:"id"`
	StartDate string `form:"start_date"`
	Year      string `form:"year"`
}

type GetAttendanceDetailInput struct {
	Id int `uri:"id" binding:"required"`
}
