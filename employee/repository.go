package employee

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	FindByUsername(userName string) (Employee, error)
	Save(employee Employee) (Employee, error)
	FindAll() ([]Employee, error)
	SearchEmployee(input SearchEmployeeInput) ([]Employee, error)
	FindById(id int) (Employee, error)
	Update(employee Employee) (Employee, error)
	FindByUsernameUpdate(id int, userName string) (Employee, error)
	Delete(id int) (Employee, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByUsername(userName string) (Employee, error) {
	var employee Employee
	now := time.Now()

	err := r.db.Where("username = ?", userName).Preload("Attendance", "attendances.attendance_date = ?", now.Format("2006-01-02")).Find(&employee).Error

	if err != nil {
		return employee, err
	}

	return employee, nil
}

func (r *repository) Save(employee Employee) (Employee, error) {
	err := r.db.Create(&employee).Error

	if err != nil {
		return employee, err
	}

	return employee, nil
}

func (r *repository) FindAll() ([]Employee, error) {
	var employees []Employee

	err := r.db.Find(&employees).Error
	if err != nil {
		return employees, err
	}

	return employees, nil
}

func (r *repository) SearchEmployee(input SearchEmployeeInput) ([]Employee, error) {
	var employees []Employee

	err := r.db.Where(r.db.Where("employee_name LIKE ?", "%"+input.EmployeeName+"%").
		Where("employee_status LIKE ?", "%"+input.EmployeeStatus+"%").
		// Where(r.db.Where("nullif(@division_id, 0) = 0", sql.Named("division_id", input.Division)).Or("division_id IN (?)", r.db.Table("division").Select("division_id"))).
		Where(r.db.Where("@division_id = 0", sql.Named("division_id", input.Division)).Or("division_id = ?", input.Division)).
		// Where(r.db.Where("nullif(@role_id, 0) = 0", sql.Named("role_id", input.Role)).Or("role_id IN (?)", r.db.Table("role").Select("role_id"))).
		Where(r.db.Where("@role_id = 0", sql.Named("role_id", input.Role)).Or("role_id = ?", input.Role))).
		Find(&employees).
		Error

	if err != nil {
		return employees, err
	}

	return employees, nil
}

func (r *repository) FindById(id int) (Employee, error) {
	var employee Employee

	err := r.db.Where("employee_id = ?", id).Find(&employee).Error
	if err != nil {
		return employee, err
	}

	return employee, nil
}

func (r *repository) Update(employee Employee) (Employee, error) {
	err := r.db.Model(&employee).Omit("employee_id").Where("employee_id = ?", employee.EmployeeId).Updates(&employee).Error

	if err != nil {
		return employee, err
	}

	return employee, nil
}

func (r *repository) FindByUsernameUpdate(id int, userName string) (Employee, error) {
	var employee Employee

	err := r.db.Where("employee_id <> ? AND username = ?", id, userName).Find(&employee).Error
	if err != nil {
		return employee, err
	}

	return employee, nil
}

func (r *repository) Delete(id int) (Employee, error) {
	var employee Employee

	err := r.db.Delete(&employee, "employee_id = ?", id).Error
	if err != nil {
		return employee, err
	}

	return employee, nil
}
