package employee

type EmployeeFormatter struct {
	EmployeeName   string `json:"employee_name"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	JenisKelamin   string `json:"sex"`
	City           string `json:"city"`
	Province       string `json:"province"`
	Address        string `json:"address"`
	DivisionId     int    `json:"division_id"`
	RoleId         int    `json:"role_id"`
	Zip            string `json:"zip"`
	Password       string `json:"password"`
	Username       string `json:"username"`
	Image          string `json:"image"`
	AcctName       string `json:"acct_name"`
	BankAcct       string `json:"bank_acct"`
	AcctNumber     string `json:"acct_number"`
	BasicSalary    int    `json:"basic_salary"`
	BeginContract  string `json:"begin_contract"`
	EndContract    string `json:"end_contract"`
	EmployeeStatus string `json:"employee_status"`
	IsPermanent    bool   `json:"is_permanent"`
	EmployeeNik    string `json:"employee_nik"`
}

func FormatEmployee(user Employee, token string) EmployeeFormatter {
	formatter := EmployeeFormatter{
		EmployeeName: user.EmployeeName,
		Email:        user.Email,
		Phone:        user.Phone,
		JenisKelamin: user.JenisKelamin,
		City:         user.City,
		Province:     user.Province,
		Address:      user.Address,
		DivisionId:   user.DivisionId,
		RoleId:       user.RoleId,
		Zip:          user.Zip,
		Password:     user.Password,
		Username:     user.Username,
		Image:        user.Image,
		AcctName:     user.AcctName,
		BankAcct:     user.BankAcct,
		AcctNumber:   user.AcctNumber,
		BasicSalary:  user.BasicSalary,
		// BeginContract:  user.BeginContract,
		// EndContract:    user.EndContract,
		EmployeeStatus: user.EmployeeStatus,
		IsPermanent:    user.IsPermanent,
		EmployeeNik:    user.EmployeeNik,
	}

	return formatter
}
