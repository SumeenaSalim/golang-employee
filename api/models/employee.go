package models

import (
	"github.com/edgedb/edgedb-go"
	"github.com/employee/api/models/objects"
	"github.com/employee/api/models/persistance"
	"github.com/gin-gonic/gin"
)

func CreateEmployee(c *gin.Context, employee objects.Employee) (objects.EmployeeData, error) {
	params := map[string]interface{}{
		"name":     employee.Name,
		"position": employee.Position,
		"salary":   employee.Salary,
	}
	result, err := persistance.CreateEmployee(c, params)
	if err != nil {
		return objects.EmployeeData{}, err
	}
	return result, nil
}

func GetEmployee(c *gin.Context, empID string) (objects.EmployeeData, error) {
	result, err := persistance.GetIndividualEmployee(c, empID)
	if err != nil {
		return objects.EmployeeData{}, err
	}

	return result, nil
}

func UpdateEmployee(c *gin.Context, empID string, emp objects.Employee) (objects.EmployeeData, error) {
	employeeID, _ := edgedb.ParseUUID(empID)
	params := map[string]interface{}{
		"employeeID": employeeID,
		"name": emp.Name,
		"position": emp.Position,
		"salary": emp.Salary,
	}

	result, err := persistance.UpdateEmployee(c, params)
	if err != nil {
		return objects.EmployeeData{}, err
	}

	return result, nil
}

func DeleteEmployee(c *gin.Context, empID string) error {
	err := persistance.DeleteEmployee(c, empID)
	if err != nil {
		return err
	}

	return nil
}