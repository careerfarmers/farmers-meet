package main

import (
    "github.com/gin-gonic/gin"
    "farmers-meet.com/apiProject/routes"
)

func main() {
    r := gin.Default()

    // Initialize routes
    routes.SetupRoutes(r)

    r.Run(":8080")
}
