package attendance

type Service interface {
	ClockInAttendance(id int, input CreateAttendanceInput) (Attendance, error)
	ListAttendanceLog(id int, input SearchAttendanceLogInput) ([]Attendance, error)
	ClockOutAttendance(inputId GetAttendanceDetailInput, input UpdateAttendanceInput) (Attendance, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) ClockInAttendance(id int, input CreateAttendanceInput) (Attendance, error) {
	// now := time.Now()
	attendance := Attendance{}

	attendance.EmployeeId = id
	attendance.AttendanceType = input.AttendaceType
	attendance.AttendanceDate = input.AttendaceDate
	attendance.CheckinTime = input.ClockinTime
	attendance.Longitude = input.Longitude
	attendance.Latitude = input.Latitude

	newAttendance, err := s.repository.Save(attendance)
	if err != nil {
		return newAttendance, err
	}

	return newAttendance, nil
}

func (s *service) ListAttendanceLog(id int, input SearchAttendanceLogInput) ([]Attendance, error) {
	if input.StartDate != "" || input.Year != "" {
		attendances, err := s.repository.SearchAttendanceLog(id, input)
		if err != nil {
			return attendances, err
		}

		return attendances, nil
	}

	attendances, err := s.repository.FindAll()
	if err != nil {
		return attendances, err
	}

	return attendances, nil
}

func (s *service) ClockOutAttendance(inputId GetAttendanceDetailInput, inputData UpdateAttendanceInput) (Attendance, error) {
	attendance := Attendance{}

	attendance.AttendanceId = inputId.Id
	attendance.CheckoutTime = inputData.ClockoutTime

	clockOutAttendance, err := s.repository.Update(attendance)
	if err != nil {
		return clockOutAttendance, err
	}

	return clockOutAttendance, nil
}
