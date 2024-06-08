package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"github.com/employee/api/response"
	"strings"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func BindAndValidateRequestBody(c *gin.Context, target interface{}) response.ErrorDetails {
	var errors response.ErrorDetails

	if bindErr := c.ShouldBindJSON(target); bindErr != nil {
		switch e := bindErr.(type) {

		case validator.ValidationErrors:
			handleValidationErrors(e, &errors)

		case *json.UnmarshalTypeError:
			handleJSONTypeError(e, &errors)

		case *json.SyntaxError:
			handleJSONSyntaxError(e, &errors)

		default:
			if bindErr == io.EOF {
				bindErr = fmt.Errorf("request body is missing")
			}
			handleUnknownErrors(bindErr, &errors)

		}
	}
	return errors
}

func handleUnknownErrors(e error, errs *response.ErrorDetails) {
	log.Printf("DEBUG: Consider adding a case to handle the error type '%T' in the function BindAndValidateRequestBody", e)
	*errs = []response.ErrorDetail{{
		Error: e.Error(),
	}}
}

func handleValidationErrors(e validator.ValidationErrors, errs *response.ErrorDetails) {
	for _, ve := range e {
		jsonField := parseJSONNameSpace(ve)
		err := response.ErrorDetail{
			Field: jsonField,
			Error: generateValidationErrorMessage(ve, jsonField),
		}
		*errs = append(*errs, err)
	}
}

func handleJSONTypeError(e *json.UnmarshalTypeError, errs *response.ErrorDetails) {
	*errs = []response.ErrorDetail{{
		Field: e.Field,
		Error: fmt.Sprintf("expected %v; got %v", e.Type, e.Value),
	}}
}

func handleJSONSyntaxError(e *json.SyntaxError, errs *response.ErrorDetails) {
	*errs = []response.ErrorDetail{{
		Error: e.Error(),
	}}
}

func parseJSONNameSpace(err validator.FieldError) string {
	jsonNameSpace := err.Namespace()
	firstString, secondString, _ := strings.Cut(jsonNameSpace, ".")

	if unicode.IsUpper(rune(firstString[0])) {
		return secondString
	}
	return jsonNameSpace
}

func generateValidationErrorMessage(err validator.FieldError, fieldName string) string {
	switch err.Tag() {

	case "required":
		return "required"

	case "min":
		return fmt.Sprintf("minimum %v characters required", err.Param())

	case "max":
		return fmt.Sprintf("maximum %v characters allowed", err.Param())

	default:
		log.Printf("DEBUG: Add a custom error message for JSON binding tag '%v' used in '%v'", err.Tag(), fieldName)
		return fmt.Sprintf("%v | %v", err.Tag(), err.Param())

	}
}
