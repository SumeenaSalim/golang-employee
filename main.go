package main

import (
	"fmt"

	"github.com/employee/api/routes"
	"github.com/employee/dbconnect"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Initializing...")

	dbconnect.ConnectDatabase()
	defer dbconnect.DbClient.Close()

	router := gin.Default()

	routes.InitializeRoutes(router)
	router.Run(":8080")
}
