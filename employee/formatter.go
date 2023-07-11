package employee

import "time"

type UserInfoFormatter struct {
	EmployeeId   int    `json:"employee_id"`
	EmployeeNik  string `json:"employee_nik"`
	EmployeeName string `json:"employee_name"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	DivisionId   int    `json:"division_id"`
	RoleId       int    `json:"role_id"`
	Image        string `json:"image"`
	Token        string `json:"token"`
}

type EmployeeFormatter struct {
	EmployeeId     int       `json:"employee_id"`
	EmployeeName   string    `json:"employee_name"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	JenisKelamin   string    `json:"sex"`
	City           string    `json:"city"`
	Province       string    `json:"province"`
	Address        string    `json:"address"`
	DivisionId     int       `json:"division_id"`
	RoleId         int       `json:"role_id"`
	Zip            string    `json:"zip"`
	Password       string    `json:"password"`
	Username       string    `json:"username"`
	Image          string    `json:"image"`
	AcctName       string    `json:"acct_name"`
	BankAcct       string    `json:"bank_acct"`
	AcctNumber     string    `json:"acct_number"`
	BasicSalary    int       `json:"basic_salary"`
	BeginContract  time.Time `json:"begin_contract"`
	EndContract    time.Time `json:"end_contract"`
	EmployeeStatus string    `json:"employee_status"`
	IsPermanent    bool      `json:"is_permanent"`
	EmployeeNik    string    `json:"employee_nik"`
	TotalCuti      int       `json:"total_cuti"`
}

func FormatUserInfo(employee Employee, token string) UserInfoFormatter {
	formatter := UserInfoFormatter{
		EmployeeId:   employee.EmployeeId,
		EmployeeNik:  employee.EmployeeNik,
		EmployeeName: employee.EmployeeName,
		Username:     employee.Username,
		Email:        employee.Email,
		DivisionId:   employee.DivisionId,
		RoleId:       employee.RoleId,
		Image:        employee.Image,
		Token:        token,
	}

	return formatter
}

func FormatEmployee(employee Employee) EmployeeFormatter {
	formatter := EmployeeFormatter{
		EmployeeId:     employee.EmployeeId,
		EmployeeName:   employee.EmployeeName,
		Email:          employee.Email,
		Phone:          employee.Phone,
		JenisKelamin:   employee.JenisKelamin,
		City:           employee.City,
		Province:       employee.Province,
		Address:        employee.Address,
		DivisionId:     employee.DivisionId,
		RoleId:         employee.RoleId,
		Zip:            employee.Zip,
		Password:       employee.Password,
		Username:       employee.Username,
		Image:          employee.Image,
		AcctName:       employee.AcctName,
		BankAcct:       employee.BankAcct,
		AcctNumber:     employee.AcctNumber,
		BasicSalary:    employee.BasicSalary,
		BeginContract:  employee.BeginContract,
		EndContract:    employee.EndContract,
		EmployeeStatus: employee.EmployeeStatus,
		IsPermanent:    employee.IsPermanent,
		EmployeeNik:    employee.EmployeeNik,
		TotalCuti:      employee.TotalCuti,
	}

	return formatter
}

func FormatEmployees(employees []Employee) []EmployeeFormatter {
	employeesFormatter := []EmployeeFormatter{}

	for _, employee := range employees {
		employeeFormatter := FormatEmployee(employee)
		employeesFormatter = append(employeesFormatter, employeeFormatter)
	}

	return employeesFormatter
}
