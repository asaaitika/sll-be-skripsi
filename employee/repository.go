package employee

import "gorm.io/gorm"

type Repository interface {
	Save(user Employee) (Employee, error)
	FindByUsername(username string) (Employee, error)
	FindById(id int) (Employee, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user Employee) (Employee, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByUsername(username string) (Employee, error) {
	var user Employee

	err := r.db.Where("username = ?", username).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindById(id int) (Employee, error) {
	var user Employee

	err := r.db.Where("employee_id = ?", id).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
