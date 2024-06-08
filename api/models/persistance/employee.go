package persistance

import (
	"fmt"

	"github.com/edgedb/edgedb-go"
	"github.com/employee/api/models/objects"
	"github.com/employee/dbconnect"
	"github.com/gin-gonic/gin"
)

func CreateEmployee(c *gin.Context, params map[string]interface{}) (objects.EmployeeData, error) {
	var emp objects.EmployeeData

	query := `
		INSERT Employees::Employees {
			name := <str>$name,
			position := <str>$position,
			salary := <float64>$salary
		};
	`
	err := dbconnect.DbClient.QuerySingle(c.Request.Context(), query, &emp, params)
	return emp, err
}

func GetIndividualEmployee(c *gin.Context, empID string) (objects.EmployeeData, error) {
	employeeID, _ := edgedb.ParseUUID(empID)
	params := map[string]interface{}{
		"employeeID": employeeID,
	}
	var result objects.EmployeeData

	query := `
		SELECT Employees::Employees {id, name, position, salary} FILTER .id = <uuid>$employeeID;
	`

	err := dbconnect.DbClient.QuerySingle(c.Request.Context(), query, &result, params)
	return result, err
}

func UpdateEmployee(c *gin.Context, params map[string]interface{}) error {
	var emp objects.EmployeeData
	fmt.Print("\n\n*********************", params)

	query := `
		UPDATE Employees::Employees 
		FILTER .id = <uuid>$employeeID
		SET {
			name := <str>$name,
			position := <str>$position,
			salary := <float64>$salary
		};
	`

	err := dbconnect.DbClient.QuerySingle(c.Request.Context(), query, &emp, params)
	return err
}

func DeleteEmployee(c *gin.Context, empID string) error {
	employeeID, _ := edgedb.ParseUUID(empID)
	params := map[string]interface{}{
		"employeeID": employeeID,
	}
	var result objects.EmployeeData

	query := `
		DELETE Employees::Employees {id} FILTER .id = <uuid>$employeeID;
	`

	err := dbconnect.DbClient.QuerySingle(c.Request.Context(), query, &result, params)
	return err
}

func GetAllEmployeesCount(c *gin.Context, params map[string]interface{}) (int64, error) {
	var query string
	var result int64
	query = fmt.Sprintf("SELECT count (( %s ))", getEmployeesQuery())

	err := dbconnect.DbClient.QuerySingle(c.Request.Context(), query, &result, params)
	return result, err
}

func getEmployeesQuery() string {
	query := `
		SELECT Employees::Employees {
			id,
			name,
			position,
			salary
		}
	`
	return query
}

func GetAllEmployees(c *gin.Context, params map[string]interface{}) ([]objects.EmployeeData, error) {
	type EmployeeList []objects.EmployeeData
	var result EmployeeList
	query := fmt.Sprintf("%s ORDER BY .created_at DESC offset <int64>$offset limit <int64>$limit", getEmployeesQuery())

	err := dbconnect.DbClient.Query(c.Request.Context(), query, &result, params)
	return result, err
}
