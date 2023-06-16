package handler

import (
	"net/http"
	"sll-be-skripsi/employee"
	"sll-be-skripsi/helper"

	"github.com/gin-gonic/gin"
)

type employeeHandler struct {
	employeeService employee.Service
}

func NewEmployeeHandler(userService employee.Service) *employeeHandler {
	return &employeeHandler{userService}
}

func (h *employeeHandler) RegisterEmployee(c *gin.Context) {
	var input employee.RegisterEmployeeInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Save data employee failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.employeeService.RegisterEmployee(input)
	if err != nil {
		response := helper.APIResponse("Save data employee failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(newUser.EmployeeId)
	if err != nil {
		response := helper.APIResponse("Save data employee failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := employee.FormatEmployee(newUser, token)

	response := helper.APIResponse("Data employee has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

// func (h *employeeHandler) Login(c *gin.Context) {
// 	var input employee.LoginInput

// 	err := c.ShouldBind(&input)
// 	if err != nil {
// 		errors := helper.FormatValidationError(err)
// 		errorMessage := gin.H{"errors": errors}

// 		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}

// 	loggedInUser, err := h.employeeService.Login(input)
// 	if err != nil {
// 		errorMessage := gin.H{"errors": err.Error()}

// 		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}

// 	token, err := h.authService.GenerateToken(loggedInUser.EmployeeId)
// 	if err != nil {
// 		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	formatter := employee.FormatEmployee(loggedInUser, token)

// 	response := helper.APIResponse("Successfully logged in", http.StatusOK, "succes", formatter)
// 	c.JSON(http.StatusOK, response)
// }
