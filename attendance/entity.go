package attendance

type Attendance struct {
	AttendanceId   int `gorm:"type:int unsigned auto_increment;PRIMARY_KEY" json:"id"`
	AttendanceDate string
	CheckinTime    string
	CheckoutTime   string
	Location       string
	EmployeeId     int
	Latitude       string
	Longitude      string
	AttendanceType string
}
