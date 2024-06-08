package routes

import (
	"net/http"

	"github.com/employee/api/constsval"
	"github.com/employee/api/handler"
	"github.com/employee/api/response"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		response.JSONError(c, http.StatusNotFound, constsval.ERROR_NOT_FOUND)
	})

	apiV1 := router.Group("/api/v1")

	apiV1.POST("/employee/", handler.CreateEmployee)
	apiV1.GET("/employee/:empId", handler.GetIndividualEmployee)
	apiV1.PUT("/employee/:empId", handler.UpdateEmployee)
	apiV1.DELETE("/employee/:empId", handler.DeleteEmployee)
	apiV1.GET("/employee", handler.GetAllEmployees)

}
