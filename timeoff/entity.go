package timeoff

import (
	"sll-be-skripsi/employee"
	"time"
)

type TimeOff struct {
	TimeoffId     int `gorm:"type:int unsigned auto_increment;PRIMARY_KEY" json:"id"`
	TimeoffType   string
	TimeoffSaldo  string
	StatusTimeoff string
	StartDate     time.Time
	EndDate       time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
	RequestType   string
	Reason        string
	File          string
	Remark        string
	EmployeeId    int
	Employee      employee.Employee `gorm:"foreignKey:EmployeeId"`
	Role          Role              `gorm:"foreignKey:RoleId"`
}

type Role struct {
	RoleId     int `gorm:"type:int unsigned auto_increment;PRIMARY_KEY" json:"id"`
	RoleName   string
	DivisionId int
}
