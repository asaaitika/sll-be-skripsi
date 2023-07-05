package employee

import "time"

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateEmployeeInput struct {
	EmployeeName   string    `form:"employee_name" binding:"required"`
	Email          string    `form:"email" binding:"required,email"`
	Phone          string    `form:"phone" binding:"required"`
	JenisKelamin   string    `form:"sex"`
	City           string    `form:"city"`
	Province       string    `form:"province"`
	Address        string    `form:"address"`
	DivisionId     int       `form:"division_id" binding:"required"`
	RoleId         int       `form:"role_id" binding:"required"`
	Zip            string    `form:"zip"`
	Password       string    `form:"password" binding:"required"`
	Username       string    `form:"username" binding:"required"`
	AcctName       string    `form:"acct_name"`
	BankAcct       string    `form:"bank_acct"`
	AcctNumber     string    `form:"acct_number"`
	BasicSalary    int       `form:"basic_salary" binding:"required"`
	BeginContract  time.Time `form:"begin_contract"`
	EndContract    time.Time `form:"end_contract"`
	EmployeeStatus string    `form:"employee_status"`
	IsPermanent    bool      `form:"is_permanent"`
	EmployeeNik    string    `form:"employee_nik" binding:"required"`
}

type SearchEmployeeInput struct {
	EmployeeName   string `form:"employee_name"`
	EmployeeStatus string `form:"employee_status"`
	Division       int    `form:"division_id"`
	Role           int    `form:"role_id"`
}

type GetEmployeeDetailInput struct {
	Id int `uri:"id" binding:"required"`
}
