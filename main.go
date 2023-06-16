package main

import (
	"log"
	"sll-be-skripsi/employee"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	dsn := "root:F!rentia2818@tcp(localhost:3306)/sll_dev?charset=utf8mb4&parseTime=True&loc=Asia%2FJakarta"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	employeeRepository := employee.NewRepository(db)

	employeeService := employee.NewService(employeeRepository)
	// authService := auth.NewService()

	// employeeHandler := handler.NewEmployeeHandler(employeeService, authService)

	userInput := employee.RegisterEmployeeInput{}
	userInput.EmployeeName = "test dr vscode part 3"
	userInput.Email = ""
	userInput.Phone = ""
	userInput.JenisKelamin = ""
	userInput.City = ""
	userInput.Province = ""
	userInput.Address = ""
	userInput.DivisionId = 0
	userInput.RoleId = 0
	userInput.Zip = ""
	userInput.Password = "test123"
	userInput.Username = "test"
	userInput.Image = ""
	userInput.AcctName = ""
	userInput.BankAcct = ""
	userInput.AcctNumber = ""
	userInput.BasicSalary = 0
	userInput.BeginContract = "2022-02-02 00:00:00"
	userInput.EndContract = "2026-02-02 00:00:00"
	userInput.EmployeeStatus = ""
	userInput.IsPermanent = false
	userInput.EmployeeNik = ""

	employeeService.RegisterEmployee(userInput)

	// router := gin.Default()
	// router.Use(CORSMiddleware())
	// router.Static("/images", "./images")
	// api := router.Group("/api/v1")

	// api.POST("/employee", employeeHandler.RegisterEmployee)
	// api.POST("/sessions", employeeHandler.Login)

	// router.Run()
}

// func authMiddleware(authService auth.Service, employeeService employee.Service) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")

// 		if !strings.Contains(authHeader, "Bearer") {
// 			res := helper.APIResponse("Unauthorized Bearer", http.StatusUnauthorized, "error", nil)
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
// 			return
// 		}

// 		tokenString := ""
// 		arrayToken := strings.Split(authHeader, " ")
// 		if len(arrayToken) == 2 {
// 			tokenString = arrayToken[1]
// 		}

// 		token, err := authService.ValidateToken(tokenString)
// 		if err != nil {
// 			res := helper.APIResponse("Unauthorized Token", http.StatusUnauthorized, "error", nil)
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
// 			return
// 		}

// 		claim, ok := token.Claims.(jwt.MapClaims)
// 		if !ok || !token.Valid {
// 			res := helper.APIResponse("Unauthorized Okay?", http.StatusUnauthorized, "error", nil)
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
// 			return
// 		}

// 		userId := int(claim["user_id"].(float64))
// 		user, err := employeeService.GetUserById(userId)
// 		if err != nil {
// 			res := helper.APIResponse("Unauthorized Employee ID", http.StatusUnauthorized, "error", nil)
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
// 			return
// 		}

// 		c.Set("currentUser", user)
// 	}
// }

// func CORSMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(204)
// 			return
// 		}

// 		c.Next()
// 	}
// }
