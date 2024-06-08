package handler

import (
	"fmt"
	"net/http"

	"github.com/employee/api/constsval"
	"github.com/employee/api/models"
	"github.com/employee/api/models/objects"
	"github.com/employee/api/models/persistance"
	"github.com/employee/api/response"
	"github.com/employee/api/utils"
	"github.com/gin-gonic/gin"
)

func CreateEmployee(c *gin.Context) {
	var emp objects.Employee

	if err := utils.BindAndValidateRequestBody(c, &emp); err != nil {
		response.JSONError(c, http.StatusBadRequest, constsval.INVALID_REQUEST_BODY, err)
		return
	}

	err := emp.ValidateEmployee()
	if len(err) > 0 {
		response.JSONError(c, http.StatusBadRequest, constsval.INVALID_REQUEST_BODY, err)
		return
	}

	result, EmpErr := models.CreateEmployee(c, emp)
	if EmpErr != nil {
		response.JSONError(c, http.StatusBadRequest, EmpErr.Error(), response.ErrorDetail{Error: EmpErr.Error()})
		return
	}
	empId, _ := result.ID.Get()
	message := fmt.Sprintf("%s successfully created", empId)

	response.JSONSuccess(c, http.StatusCreated, "", message)
}

func GetIndividualEmployee(c *gin.Context) {
	empId := c.Param("empId")
	result, err := models.GetEmployee(c, empId)
	if err != nil {
		response.JSONError(c, http.StatusBadRequest, constsval.ID_NOT_FOUND, response.ErrorDetail{Error: constsval.ID_NOT_FOUND})
		return
	}

	response.JSONSuccess(c, http.StatusOK, result)
}

func UpdateEmployee(c *gin.Context) {
	empId := c.Param("empId")
	var emp objects.Employee
	if err := utils.BindAndValidateRequestBody(c, &emp); err != nil {
		response.JSONError(c, http.StatusBadRequest, constsval.INVALID_REQUEST_BODY, err)
		return
	}

	err := emp.ValidateEmployee()
	if len(err) > 0 {
		response.JSONError(c, http.StatusBadRequest, constsval.INVALID_REQUEST_BODY, err)
		return
	}

	EmpErr := models.UpdateEmployee(c, empId, emp)
	if err != nil {
		response.JSONError(c, http.StatusBadRequest, EmpErr.Error(), response.ErrorDetail{Error: EmpErr.Error()})
		return
	}

	message := fmt.Sprintf("%s successfully updated", empId)
	response.JSONSuccess(c, http.StatusOK, "", message)
}

func DeleteEmployee(c *gin.Context) {
	empId := c.Param("empId")
	err := models.DeleteEmployee(c, empId)
	if err != nil {
		response.JSONError(c, http.StatusBadRequest, err.Error(), response.ErrorDetail{Error: err.Error()})
		return
	}

	message := fmt.Sprintf("%s successfully deleted", empId)
	response.JSONSuccess(c, http.StatusOK, "", message)
}

func GetAllEmployees(c *gin.Context) {
	limit, offset, _ := utils.GetPaginationData(c)
	params := map[string]interface{}{
		"limit":  limit,
		"offset": offset,
	}
	count, _ := persistance.GetAllEmployeesCount(c, params)
	result, err := persistance.GetAllEmployees(c, params)

	if err != nil {
		response.JSONError(c, http.StatusBadRequest, err.Error(), response.ErrorDetail{Error: err.Error()})
		return
	}

	type pagination struct {
		Limit  int64 `edgedb:"limit" json:"limit"`
		Offset int64 `edgedb:"offset" json:"offset"`
		Count  int64 `edgedb:"count" json:"count"`
	}

	res := struct {
		EmployeeData []objects.EmployeeData `json:"employee"`
		Pagination   pagination             `json:"pagination"`
	}{
		EmployeeData: result,
		Pagination: pagination{
			Limit:  limit,
			Offset: offset,
			Count:  count,
		},
	}
	response.JSONSuccess(c, http.StatusOK, res)

}
