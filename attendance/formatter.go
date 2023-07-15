package attendance

type AttendanceFormatter struct {
	AttendanceId   int    `json:"attendance_id"`
	EmployeeId     int    `json:"employee_id"`
	AttendanceDate string `json:"attendance_date"`
	AttendanceType string `json:"attendance_type"`
	CheckinTime    string `json:"clockin_time"`
	CheckoutTime   string `json:"clockout_time"`
	Latitude       string `json:"latitude"`
	Longitude      string `json:"longitude"`
}

func FormatAttendance(employee Attendance) AttendanceFormatter {
	formatter := AttendanceFormatter{
		AttendanceId:   employee.AttendanceId,
		EmployeeId:     employee.EmployeeId,
		AttendanceDate: employee.AttendanceDate,
		AttendanceType: employee.AttendanceType,
		CheckinTime:    employee.CheckinTime,
		CheckoutTime:   employee.CheckoutTime,
		Latitude:       employee.Latitude,
		Longitude:      employee.Longitude,
	}

	return formatter
}

func FormatAttendances(attendances []Attendance) []AttendanceFormatter {
	attendancesFormatter := []AttendanceFormatter{}

	for _, attendance := range attendances {
		attendanceFormatter := FormatAttendance(attendance)
		attendancesFormatter = append(attendancesFormatter, attendanceFormatter)
	}

	return attendancesFormatter
}
