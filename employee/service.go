package employee

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(input LoginInput) (Employee, error)
	GetUserById(id int) (Employee, error)
	CreateEmployee(input CreateEmployeeInput, images string) (Employee, error)
	ListEmployee(data SearchEmployeeInput) ([]Employee, error)
	GetEmployeeById(input GetEmployeeDetailInput) (Employee, error)
	UpdateEmployee(inputId GetEmployeeDetailInput, inputData CreateEmployeeInput, images string) (Employee, error)
	DeleteEmployee(input GetEmployeeDetailInput, id int) (Employee, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
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

func (s *service) CreateEmployee(input CreateEmployeeInput, images string) (Employee, error) {
	employee := Employee{}

	employee.EmployeeName = input.EmployeeName
	employee.Email = input.Email
	employee.Phone = input.Phone
	employee.JenisKelamin = input.JenisKelamin
	employee.City = input.City
	employee.Province = input.Province
	employee.Address = input.Address
	employee.DivisionId = input.DivisionId
	employee.RoleId = input.RoleId
	employee.Zip = input.Zip
	employee.Image = images
	employee.AcctName = input.AcctName
	employee.BankAcct = input.BankAcct
	employee.AcctNumber = input.AcctNumber
	employee.BasicSalary = input.BasicSalary
	employee.BeginContract = input.BeginContract
	employee.EndContract = input.EndContract
	employee.EmployeeStatus = input.EmployeeStatus
	employee.IsPermanent = input.IsPermanent
	employee.EmployeeNik = input.EmployeeNik

	newUsername, err := s.repository.FindByUsername(input.Username)
	if err != nil {
		return employee, err
	}

	if newUsername.EmployeeId > 0 {
		return employee, errors.New("this username is already in use")
	}

	employee.Username = input.Username

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return employee, err
	}

	employee.Password = string(passwordHash)

	newEmployee, err := s.repository.Save(employee)
	if err != nil {
		return newEmployee, err
	}

	return newEmployee, nil
}

func (s *service) ListEmployee(input SearchEmployeeInput) ([]Employee, error) {
	if input.EmployeeName != "" || input.EmployeeStatus != "" || input.Division != 0 || input.Role != 0 {
		employees, err := s.repository.SearchEmployee(input)
		if err != nil {
			return employees, err
		}

		return employees, nil
	}

	employees, err := s.repository.FindAll()
	if err != nil {
		return employees, err
	}

	return employees, nil
}

func (s *service) GetEmployeeById(input GetEmployeeDetailInput) (Employee, error) {
	employees, err := s.repository.FindById(input.Id)

	if err != nil {
		return employees, err
	}

	return employees, nil
}

func (s *service) UpdateEmployee(inputId GetEmployeeDetailInput, inputData CreateEmployeeInput, images string) (Employee, error) {
	employee, err := s.repository.FindById(inputId.Id)
	if err != nil {
		return employee, err
	}

	employee.EmployeeName = inputData.EmployeeName
	employee.Email = inputData.Email
	employee.Phone = inputData.Phone
	employee.JenisKelamin = inputData.JenisKelamin
	employee.City = inputData.City
	employee.Province = inputData.Province
	employee.Address = inputData.Address
	employee.DivisionId = inputData.DivisionId
	employee.RoleId = inputData.RoleId
	employee.Zip = inputData.Zip
	employee.Image = images
	employee.AcctName = inputData.AcctName
	employee.BankAcct = inputData.BankAcct
	employee.AcctNumber = inputData.AcctNumber
	employee.BasicSalary = inputData.BasicSalary
	employee.BeginContract = inputData.BeginContract
	employee.EndContract = inputData.EndContract
	employee.EmployeeStatus = inputData.EmployeeStatus
	employee.IsPermanent = inputData.IsPermanent
	employee.EmployeeNik = inputData.EmployeeNik

	newUsername, err := s.repository.FindByUsernameUpdate(inputId.Id, inputData.Username)
	if err != nil {
		return employee, err
	}

	if newUsername.EmployeeId > 0 {
		return employee, errors.New("this username is already in use")
	}

	employee.Username = inputData.Username

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(inputData.Password), bcrypt.MinCost)
	if err != nil {
		return employee, err
	}

	employee.Password = string(passwordHash)

	updateEmployee, err := s.repository.Update(employee)
	if err != nil {
		return updateEmployee, err
	}

	return updateEmployee, nil
}

func (s *service) DeleteEmployee(input GetEmployeeDetailInput, id int) (Employee, error) {
	employees, err := s.repository.FindById(input.Id)

	if err != nil {
		return employees, err
	}

	if id == input.Id {
		return employees, errors.New("admin can't delete his own data")
	}

	employee, err := s.repository.Delete(input.Id)

	if err != nil {
		return employee, err
	}

	return employee, nil
}
