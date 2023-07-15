package handler

import (
	"fmt"
	"net/http"
	"sll-be-skripsi/auth"
	"sll-be-skripsi/employee"
	"sll-be-skripsi/helper"
	"sll-be-skripsi/timeoff"

	"github.com/gin-gonic/gin"
)

type timeoffHandler struct {
	timeoffService timeoff.Service
	authService    auth.Service
}

func NewTimeOffHandler(timeoffService timeoff.Service, authService auth.Service) *timeoffHandler {
	return &timeoffHandler{timeoffService, authService}
}

func (h *timeoffHandler) RequestTimeOff(c *gin.Context) {
	var input timeoff.CreateRequestTimeOffInput

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Request time off failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(employee.Employee)
	employeeId := currentUser.EmployeeId
	input.EmployeeId = currentUser.EmployeeId

	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload file", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)

		return
	}

	path := fmt.Sprintf("files/%d-%s", employeeId, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload file", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)

		return
	}

	newTimeOff, err := h.timeoffService.CreateRequestTimeOff(input, path)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Request time off failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := timeoff.FormatTimeOff(newTimeOff)

	response := helper.APIResponse("Data request has been saved", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *timeoffHandler) ListTimeOff(c *gin.Context) {
	var input timeoff.SearchTimeOffInput

	err := c.ShouldBindQuery(&input)
	if err != nil {
		response := helper.APIResponse("Error to get param time off", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	timeoffs, err := h.timeoffService.ListTimeOff(input)
	if err != nil {
		response := helper.APIResponse("Error to get time off", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of Time Off", http.StatusOK, "success", timeoff.FormatTimeOffs(timeoffs))
	c.JSON(http.StatusOK, response)
}

func (h *timeoffHandler) ListRequestTimeOff(c *gin.Context) {
	var input timeoff.SearchRequestTimeOffInput

	err := c.ShouldBindQuery(&input)
	if err != nil {
		response := helper.APIResponse("Error to get param request time off", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	timeoffs, err := h.timeoffService.ListRequestTimeOff(input)
	if err != nil {
		response := helper.APIResponse("Error to get request time off", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of Request Time Off", http.StatusOK, "success", timeoff.FormatTimeOffs(timeoffs))
	c.JSON(http.StatusOK, response)
}

func (h *timeoffHandler) GetTimeOff(c *gin.Context) {
	var input timeoff.GetTimeOffDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get all detail of time off", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	timeoffDetail, err := h.timeoffService.GetTimeOffById(input)
	if err != nil {
		response := helper.APIResponse("Failed to get all detail of time off", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := timeoff.FormatTimeOff(timeoffDetail)

	response := helper.APIResponse("Time off detail", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *timeoffHandler) UpdateRequestTimeOff(c *gin.Context) {
	var inputId timeoff.GetTimeOffDetailInput

	err := c.ShouldBindUri(&inputId)
	if err != nil {
		response := helper.APIResponse("Failed to get param time off", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData timeoff.UpdateStatusTimeOffInput

	err = c.ShouldBind(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Update status time off failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updateTimeOff, err := h.timeoffService.UpdateRequestTimeOff(inputId, inputData)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Update status time off failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Data time off has been updated", http.StatusOK, "success", timeoff.FormatTimeOff(updateTimeOff))
	c.JSON(http.StatusOK, response)
}

func (h *timeoffHandler) DeleteTimeOff(c *gin.Context) {
	var input timeoff.GetTimeOffDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to delete time off", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	timeoffDetail, err := h.timeoffService.DeleteTimeOff(input)
	if err != nil {
		response := helper.APIResponse("Failed to delete time off", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := timeoff.FormatTimeOff(timeoffDetail)

	response := helper.APIResponse("Delete data time off success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
