package objects

import (
	"strings"

	"github.com/edgedb/edgedb-go"
	"github.com/employee/api/response"
)

type Employee struct {
	ID       edgedb.OptionalUUID `edgedb:"id" json:"id"`
	Name     string              `edgedb:"name" json:"bookName"`
	Position string              `edgedb:"position" json:"position"`
	Salary   float64             `edgedb:"salary" json:"salary"`
}

func (b Employee) ValidateEmployee() response.ErrorDetails {
	var valerr response.ErrorDetails
	if strings.TrimSpace(b.Name) == "" {
		valerr = append(valerr, response.ErrorDetail{Field: "name", Error: "required"})
	}

	if strings.TrimSpace(b.Position) == "" {
		valerr = append(valerr, response.ErrorDetail{Field: "position", Error: "required"})
	}

	if b.Salary < 1000 {
		valerr = append(valerr, response.ErrorDetail{Field: "salary", Error: "salary should be greater than 1000"})
	}

	return valerr
}

type EmployeeData struct {
	ID       edgedb.OptionalUUID    `edgedb:"id" json:"id"`
	Name     string                 `edgedb:"name" json:"name"`
	Position edgedb.OptionalStr     `edgedb:"position" json:"position"`
	Salary   edgedb.OptionalFloat64 `edgedb:"salary" json:"salary"`
}
