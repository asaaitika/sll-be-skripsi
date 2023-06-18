package main

import (
	"fmt"
	"log"
	"sll-be-skripsi/employee"
	"sll-be-skripsi/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	dsn := "root:F!rentia2818@tcp(localhost:3306)/sll_dev?parseTime=true&loc=Asia%2FJakarta&charset=utf8mb4&collation=utf8mb4_unicode_ci"
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

	input := employee.LoginInput{
		Username: "johndoemc",
		Password: "test123",
	}
	user, err := employeeService.Login(input)
	if err != nil {
		fmt.Println("salah euy")
		fmt.Println(err.Error())
	}

	fmt.Println(user.Email)

	employeeHandler := handler.NewEmployeeHandler(employeeService)

	router := gin.Default()
	// router.Use(CORSMiddleware())
	// router.Static("/images", "./images")
	api := router.Group("/api/v1")

	api.POST("/employee", employeeHandler.RegisterEmployee)
	// api.POST("/sessions", employeeHandler.Login)

	router.Run()
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
