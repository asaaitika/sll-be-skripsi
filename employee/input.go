package employee

type RegisterEmployeeInput struct {
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

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
