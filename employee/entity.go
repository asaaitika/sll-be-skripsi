package employee

import (
	"sll-be-skripsi/attendance"
	"time"
)

type Employee struct {
	EmployeeId     int `gorm:"type:int unsigned auto_increment;PRIMARY_KEY" json:"id"`
	EmployeeName   string
	Address        string
	Phone          string
	Email          string
	RoleId         int
	AcctNumber     string
	JenisKelamin   string
	AcctName       string
	BankAcct       string
	City           string
	Province       string
	DivisionId     int
	Zip            string
	Image          string
	Username       string
	Password       string
	EmployeeStatus string
	EndContract    time.Time
	BeginContract  time.Time
	IsPermanent    bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
	EmployeeNik    string
	BasicSalary    int
	Token          string
	TotalCuti      int
	Attendance     attendance.Attendance `gorm:"foreignKey:EmployeeId"`
}
