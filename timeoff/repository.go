package timeoff

import (
	"database/sql"

	"gorm.io/gorm"
)

type Repository interface {
	Save(timeoff TimeOff) (TimeOff, error)
	FindAllByEmployeeId(id int) ([]TimeOff, error)
	SearchTimeOff(input SearchTimeOffInput, id int) ([]TimeOff, error)
	FindAll() ([]TimeOff, error)
	SearchRequestTimeOff(input SearchRequestTimeOffInput) ([]TimeOff, error)
	FindById(id int) (TimeOff, error)
	Update(timeoff TimeOff) (TimeOff, error)
	Delete(id int) (TimeOff, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(timeoff TimeOff) (TimeOff, error) {
	err := r.db.Create(&timeoff).Error

	if err != nil {
		return timeoff, err
	}

	return timeoff, nil
}

func (r *repository) FindAllByEmployeeId(id int) ([]TimeOff, error) {
	var timeoff []TimeOff

	rows, err := r.db.Table("vad_time_off").Select("*").Where("employee_id = ?", id).Rows()
	for rows.Next() {
		r.db.ScanRows(rows, &timeoff)
	}

	if err != nil {
		return timeoff, err
	}

	return timeoff, nil
}

func (r *repository) SearchTimeOff(input SearchTimeOffInput, id int) ([]TimeOff, error) {
	var timeoff []TimeOff

	rows, err := r.db.Table("vad_time_off").Select("*").
		Where(r.db.Where("employee_id = ?", id).
			Where(r.db.Where("@status_timeoff = ''", sql.Named("status_timeoff", input.Status)).Or("status_timeoff = ?", input.Status)).
			Where(r.db.Where("@month = 0", sql.Named("month", input.Month)).Or("MONTH(start_date) = ?", input.Month)).
			Where(r.db.Where("@year = 0", sql.Named("year", input.Year)).Or("YEAR(start_date) = ?", input.Year))).
		Rows()

	for rows.Next() {
		r.db.ScanRows(rows, &timeoff)
	}

	if err != nil {
		return timeoff, err
	}

	return timeoff, nil
}

func (r *repository) FindAll() ([]TimeOff, error) {
	var timeoff []TimeOff

	rows, err := r.db.Table("vad_time_off").Select("*").Rows()
	for rows.Next() {
		r.db.ScanRows(rows, &timeoff)
	}

	if err != nil {
		return timeoff, err
	}

	return timeoff, nil
}

func (r *repository) SearchRequestTimeOff(input SearchRequestTimeOffInput) ([]TimeOff, error) {
	var timeoff []TimeOff

	rows, err := r.db.Table("vad_time_off").Select("*").
		Where(r.db.Where("remark = ?", input.Keyword).Or("employee_name = ?", input.Keyword).Or("role_name = ?", input.Keyword).Or("timeoff_type_name = ?", input.Keyword)).
		Where(r.db.Where("@status_timeoff = ''", sql.Named("status_timeoff", input.Status)).Or("status_timeoff = ?", input.Status)).
		Where(r.db.Where("@start_date = ''", sql.Named("start_date", input.RequestDate)).Or("DATE(start_date) = date(?)", input.RequestDate)).
		Rows()

	for rows.Next() {
		r.db.ScanRows(rows, &timeoff)
	}

	if err != nil {
		return timeoff, err
	}

	return timeoff, nil
}

func (r *repository) FindById(id int) (TimeOff, error) {
	var timeoff TimeOff

	err := r.db.Table("vad_time_off").Select("*").Where("timeoff_id = ?", id).Find(&timeoff).Error
	if err != nil {
		return timeoff, err
	}

	return timeoff, nil
}

func (r *repository) Update(timeoff TimeOff) (TimeOff, error) {
	err := r.db.Model(&timeoff).Where("timeoff_id = ?", timeoff.TimeoffId).Updates(&TimeOff{Remark: timeoff.Remark, StatusTimeoff: timeoff.StatusTimeoff}).Error
	if err != nil {
		return timeoff, err
	}

	if timeoff.StatusTimeoff == "A" {
		err = r.db.Exec("UPDATE employee SET total_cuti = total_cuti - 1 where employee_id = ?", timeoff.EmployeeId).Error

		if err != nil {
			return timeoff, err
		}
	}

	return timeoff, nil
}

func (r *repository) Delete(id int) (TimeOff, error) {
	var timeoff TimeOff

	err := r.db.Delete(&timeoff, "timeoff_id = ?", id).Error
	if err != nil {
		return timeoff, err
	}

	return timeoff, nil
}
