package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sll-be-skripsi/attendance"
	"sll-be-skripsi/auth"
	"sll-be-skripsi/employee"
	"sll-be-skripsi/handler"
	"sll-be-skripsi/helper"
	"sll-be-skripsi/timeoff"
	"strings"

	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func main() {

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("failed to load env", err)
	// }

	// Open a connection to the database
	dbConnectionString, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatal("failed to open db connection", err)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: dbConnectionString,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	employeeRepository := employee.NewRepository(db)
	timeoffRepository := timeoff.NewRepository(db)
	attendanceRepository := attendance.NewRepository(db)

	employeeService := employee.NewService(employeeRepository)
	timeoffService := timeoff.NewService(timeoffRepository)
	attendanceService := attendance.NewService(attendanceRepository)
	authService := auth.NewService()

	employeeHandler := handler.NewEmployeeHandler(employeeService, authService)
	timeoffHandler := handler.NewTimeOffHandler(timeoffService, authService)
	attendanceHandler := handler.NewAttendanceHandler(attendanceService, authService)

	router := gin.Default()
	router.Use(CORSMiddleware())
	router.Static("/images", "./images")
	router.Static("/files", "./files")

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello.. Welcome!")
	})

	api := router.Group("/api/v1")

	fmt.Println(db.Debug())

	api.POST("/login", employeeHandler.Login)
	api.POST("/employee", authMiddleware(authService, employeeService), employeeHandler.RegisterEmployee)
	api.GET("/employee", employeeHandler.ListEmployees)
	api.GET("/employee/:id", employeeHandler.GetEmployee)
	api.PUT("/employee/:id", authMiddleware(authService, employeeService), employeeHandler.UpdateEmployee)
	api.DELETE("/employee/:id", authMiddleware(authService, employeeService), employeeHandler.DeleteEmployee)
	api.POST("/timeoff", authMiddleware(authService, employeeService), timeoffHandler.RequestTimeOff)
	api.GET("/timeoff", timeoffHandler.ListTimeOff)
	api.GET("/request-timeoff", timeoffHandler.ListRequestTimeOff)
	api.GET("/timeoff/:id", timeoffHandler.GetTimeOff)
	api.PUT("/timeoff/:id", authMiddleware(authService, employeeService), timeoffHandler.UpdateRequestTimeOff)
	api.DELETE("/timeoff/:id", authMiddleware(authService, employeeService), timeoffHandler.DeleteTimeOff)
	api.POST("/attendance", authMiddleware(authService, employeeService), attendanceHandler.CreateClockInAttendance)
	api.GET("/attendance", authMiddleware(authService, employeeService), attendanceHandler.ListAttendanceLog)
	api.PUT("/attendance/:id", authMiddleware(authService, employeeService), attendanceHandler.UpdateClockOutAttendance)

	router.Run()
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
