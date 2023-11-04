package routes

import (
    "github.com/gin-gonic/gin"
    "farmers-meet.com/apiProject/controllers"
)

func SetupRoutes(r *gin.Engine) {
    // Define routes here
    r.GET("/api/hello", controllers.GetHello)
}
