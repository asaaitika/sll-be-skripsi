package employee

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterEmployee(input RegisterEmployeeInput) (Employee, error)
	Login(input LoginInput) (Employee, error)
	GetUserById(id int) (Employee, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterEmployee(input RegisterEmployeeInput) (Employee, error) {
	user := Employee{}

	user.EmployeeName = input.EmployeeName
	user.Email = input.Email
	user.Phone = input.Phone
	user.JenisKelamin = input.JenisKelamin
	user.City = input.City
	user.Province = input.Province
	user.Address = input.Address
	user.DivisionId = input.DivisionId
	user.RoleId = input.RoleId
	user.Zip = input.Zip
	user.Username = input.Username
	user.Image = input.Image
	user.AcctName = input.AcctName
	user.BankAcct = input.BankAcct
	user.AcctNumber = input.AcctNumber
	user.BasicSalary = input.BasicSalary
	user.BeginContract = input.BeginContract
	user.EndContract = input.EndContract
	user.EmployeeStatus = input.EmployeeStatus
	user.IsPermanent = input.IsPermanent
	user.EmployeeNik = input.EmployeeNik

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)
	user.RoleId = 1

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input LoginInput) (Employee, error) {
	username := input.Username
	password := input.Password

	user, err := s.repository.FindByUsername(username)
	if err != nil {
		return user, err
	}

	if user.EmployeeId == 0 {
		return user, errors.New("no employee found on that username")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) GetUserById(id int) (Employee, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		return user, err
	}

	if user.EmployeeId == 0 {
		return user, errors.New("no user found on with that id")
	}

	return user, nil
}
