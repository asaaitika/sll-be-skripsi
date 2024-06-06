package handler

import (
	"fmt"
	"net/http"
	"sll-be-skripsi/auth"
	"sll-be-skripsi/employee"
	"sll-be-skripsi/helper"

	"github.com/gin-gonic/gin"
)

type employeeHandler struct {
	employeeService employee.Service
	authService     auth.Service
}

func NewEmployeeHandler(employeeService employee.Service, authService auth.Service) *employeeHandler {
	return &employeeHandler{employeeService, authService}
}

func (h *employeeHandler) Login(c *gin.Context) {
	var input employee.LoginInput

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedInUser, err := h.employeeService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.authService.GenerateToken(loggedInUser.EmployeeId)
	if err != nil {
		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := employee.FormatUserInfo(loggedInUser, token)

	response := helper.APIResponse("Successfully logged in", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *employeeHandler) Logout(c *gin.Context) {
	response := helper.APIResponse("Successfully logged out", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *employeeHandler) RegisterEmployee(c *gin.Context) {
	var input employee.CreateEmployeeInput

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Save data employee failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(employee.Employee)
	employeeId := currentUser.EmployeeId

	file, err := c.FormFile("images")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)

		return
	}

	path := fmt.Sprintf("images/%d-%s", employeeId, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)

		return
	}

	newEmployee, err := h.employeeService.CreateEmployee(input, path)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Save data employee failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := employee.FormatEmployee(newEmployee)

	response := helper.APIResponse("Data employee has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *employeeHandler) ListEmployees(c *gin.Context) {
	var input employee.SearchEmployeeInput

	err := c.ShouldBindQuery(&input)
	if err != nil {
		response := helper.APIResponse("Error to get param employees", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	employees, err := h.employeeService.ListEmployee(input)
	if err != nil {
		response := helper.APIResponse("Error to get employees", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of employee", http.StatusOK, "success", employee.FormatEmployees(employees))
	c.JSON(http.StatusOK, response)
}

func (h *employeeHandler) GetEmployee(c *gin.Context) {
	var input employee.GetEmployeeDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get all detail of employee", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	employeeDetail, err := h.employeeService.GetEmployeeById(input)
	if err != nil {
		response := helper.APIResponse("Failed to get all detail of employee", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := employee.FormatEmployee(employeeDetail)

	response := helper.APIResponse("Employee detail", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *employeeHandler) UpdateEmployee(c *gin.Context) {
	var inputId employee.GetEmployeeDetailInput

	err := c.ShouldBindUri(&inputId)
	if err != nil {
		response := helper.APIResponse("Failed to update employee", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData employee.CreateEmployeeInput

	err = c.ShouldBind(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Update data employee failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	file, err := c.FormFile("images")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)

		return
	}

	path := fmt.Sprintf("images/%d-%s", inputId.Id, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)

		return
	}

	updateEmployee, err := h.employeeService.UpdateEmployee(inputId, inputData, path)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Update data employee failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Data employee has been updated", http.StatusOK, "success", employee.FormatEmployee(updateEmployee))
	c.JSON(http.StatusOK, response)
}

func (h *employeeHandler) DeleteEmployee(c *gin.Context) {
	var input employee.GetEmployeeDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to delete employee", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(employee.Employee)
	employeeId := currentUser.EmployeeId

	employeeDetail, err := h.employeeService.DeleteEmployee(input, employeeId)
	if err != nil {
		response := helper.APIResponse("Failed to delete employee", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := employee.FormatEmployee(employeeDetail)

	response := helper.APIResponse("Delete data employee success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
