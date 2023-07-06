package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sll-be-skripsi/auth"
	"sll-be-skripsi/employee"
	"sll-be-skripsi/handler"
	"sll-be-skripsi/helper"
	"strings"

	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func main() {

	// dsn := "root:F!rentia2818@tcp(localhost:3306)/sll_dev?parseTime=true&loc=Asia%2FJakarta&charset=utf8mb4&collation=utf8mb4_unicode_ci"

	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env", err)
	}

	// Open a connection to the database
	dbConnectionString, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatal("failed to open db connection", err)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: dbConnectionString,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	employeeRepository := employee.NewRepository(db)

	employeeService := employee.NewService(employeeRepository)
	authService := auth.NewService()

	employeeHandler := handler.NewEmployeeHandler(employeeService, authService)

	router := gin.Default()
	router.Use(CORSMiddleware())
	router.Static("/images", "./images")
	api := router.Group("/api/v1")

	fmt.Println(db.Debug())

	api.POST("/login", employeeHandler.Login)
	api.POST("/employee", authMiddleware(authService, employeeService), employeeHandler.RegisterEmployee)
	api.GET("/employee", employeeHandler.ListEmployees)
	api.GET("/employee/:id", employeeHandler.GetEmployee)
	api.PUT("/employee/:id", authMiddleware(authService, employeeService), employeeHandler.UpdateEmployee)
	api.DELETE("/employee/:id", authMiddleware(authService, employeeService), employeeHandler.DeleteEmployee)

	router.Run(":8080")
}

func authMiddleware(authService auth.Service, employeeService employee.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			res := helper.APIResponse("Unauthorized Bearer", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			res := helper.APIResponse("Unauthorized Token", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			res := helper.APIResponse("Unauthorized Token", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

		userId := int(claim["employee_id"].(float64))
		user, err := employeeService.GetUserById(userId)
		if err != nil {
			res := helper.APIResponse("Unauthorized Employee ID", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

		c.Set("currentUser", user)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
