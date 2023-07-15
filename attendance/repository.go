package attendance

import (
	"database/sql"

	"gorm.io/gorm"
)

type Repository interface {
	Save(attendance Attendance) (Attendance, error)
	FindAll() ([]Attendance, error)
	SearchAttendanceLog(id int, input SearchAttendanceLogInput) ([]Attendance, error)
	Update(attendance Attendance) (Attendance, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(attendance Attendance) (Attendance, error) {
	err := r.db.Create(&attendance).Error

	if err != nil {
		return attendance, err
	}

	return attendance, nil
}

func (r *repository) FindAll() ([]Attendance, error) {
	var attendances []Attendance

	err := r.db.Find(&attendances).Error
	if err != nil {
		return attendances, err
	}

	return attendances, nil
}

func (r *repository) SearchAttendanceLog(id int, input SearchAttendanceLogInput) ([]Attendance, error) {
	var attendances []Attendance

	err := r.db.Where(r.db.Where("employee_id = ?", id).
		Where(r.db.Where("@start_date = 0", sql.Named("start_date", input.StartDate)).Or("attendance_date = ?", input.StartDate)).
		Where(r.db.Where("@year = 0", sql.Named("year", input.Year)).Or("YEAR(attendance_date) = ?", input.Year))).
		Find(&attendances).
		Error

	if err != nil {
		return attendances, err
	}

	return attendances, nil
}

func (r *repository) Update(attendance Attendance) (Attendance, error) {
	err := r.db.Model(&attendance).Where("attendance_id = ?", attendance.AttendanceId).Updates(&Attendance{CheckoutTime: attendance.CheckoutTime}).Error
	if err != nil {
		return attendance, err
	}

	return attendance, nil
}
