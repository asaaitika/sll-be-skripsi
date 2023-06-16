package main

import (
	"fmt"
	"log"
	"sll-be-skripsi/user"

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

	fmt.Println("connection to database is good")

	var users []user.Employee
	length := len(users)

	fmt.Println(length)

	db.Find(&users)
	length = len(users)

	fmt.Println(length)

	for _, user := range users {
		fmt.Println(user.EmployeeName)
		fmt.Println(user.EmployeeNik)
		fmt.Println(user.Address)
		fmt.Println(user.Email)
		fmt.Println("=================================")
	}
}
