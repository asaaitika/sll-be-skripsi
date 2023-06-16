package main

import (
	"log"
	"sll-be-skripsi/user"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	dsn := "root:F!rentia2818@tcp(localhost:3306)/sll_dev?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	user := user.Employee{
		EmployeeName:  "test simpan part 2",
		EndContract:   time.Date(2026, 8, 15, 14, 30, 45, 100, time.Local),
		BeginContract: time.Now(),
	}

	userRepository.Save(user)
}
