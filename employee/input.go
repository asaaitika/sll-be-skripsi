package employee

import "time"

type RegisterEmployeeInput struct {
	EmployeeName   string    `json:"employee_name" binding:"required"`
	Email          string    `json:"email" binding:"required,email"`
	Phone          string    `json:"phone" binding:"required"`
	JenisKelamin   string    `json:"sex"`
	City           string    `json:"city"`
	Province       string    `json:"province"`
	Address        string    `json:"address"`
	DivisionId     int       `json:"division_id" binding:"required"`
	RoleId         int       `json:"role_id" binding:"required"`
	Zip            string    `json:"zip"`
	Password       string    `json:"password" binding:"required"`
	Username       string    `json:"username" binding:"required"`
	Image          string    `json:"image"`
	AcctName       string    `json:"acct_name"`
	BankAcct       string    `json:"bank_acct"`
	AcctNumber     string    `json:"acct_number"`
	BasicSalary    int       `json:"basic_salary" binding:"required"`
	BeginContract  time.Time `json:"begin_contract"`
	EndContract    time.Time `json:"end_contract"`
	EmployeeStatus string    `json:"employee_status"`
	IsPermanent    bool      `json:"is_permanent"`
	EmployeeNik    string    `json:"employee_nik" binding:"required"`
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
