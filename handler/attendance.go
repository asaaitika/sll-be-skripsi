package handler

import (
	"net/http"
	"sll-be-skripsi/attendance"
	"sll-be-skripsi/auth"
	"sll-be-skripsi/employee"
	"sll-be-skripsi/helper"

	"github.com/gin-gonic/gin"
)

type attendanceHandler struct {
	attendanceService attendance.Service
	authService       auth.Service
}

func NewAttendanceHandler(attendanceService attendance.Service, authService auth.Service) *attendanceHandler {
	return &attendanceHandler{attendanceService, authService}
}

func (h *attendanceHandler) CreateClockInAttendance(c *gin.Context) {
	var input attendance.CreateAttendanceInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed clock in attendance", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(employee.Employee)
	employeeId := currentUser.EmployeeId

	clockInAttendance, err := h.attendanceService.ClockInAttendance(employeeId, input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Failed clock in attendance", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := attendance.FormatAttendance(clockInAttendance)

	response := helper.APIResponse("Success Clock In at "+clockInAttendance.CheckinTime, http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *attendanceHandler) ListAttendanceLog(c *gin.Context) {
	var input attendance.SearchAttendanceLogInput

	err := c.ShouldBindQuery(&input)
	if err != nil {
		response := helper.APIResponse("Error to get param attendance log", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(employee.Employee)
	employeeId := currentUser.EmployeeId

	attendances, err := h.attendanceService.ListAttendanceLog(employeeId, input)
	if err != nil {
		response := helper.APIResponse("Error to get attendance log", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of attendance log", http.StatusOK, "success", attendance.FormatAttendances(attendances))
	c.JSON(http.StatusOK, response)
}

func (h *attendanceHandler) UpdateClockOutAttendance(c *gin.Context) {
	var inputId attendance.GetAttendanceDetailInput

	err := c.ShouldBindUri(&inputId)
	if err != nil {
		response := helper.APIResponse("Failed to get param attendance", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData attendance.UpdateAttendanceInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Update clock out attendance failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	clockOutAttendance, err := h.attendanceService.ClockOutAttendance(inputId, inputData)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Update clock out attendance failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success Clock Out at "+clockOutAttendance.CheckoutTime, http.StatusOK, "success", attendance.FormatAttendance(clockOutAttendance))
	c.JSON(http.StatusOK, response)
}
