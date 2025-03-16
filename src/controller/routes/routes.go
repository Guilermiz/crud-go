package routes

import (
	"github.com/Guilermiz/crud-go/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(routes *gin.RouterGroup) {

	routes.GET("/getUserById/:userId", controller.FindUserById)
	routes.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	routes.POST("/createUser", controller.CreateUser)
	routes.PUT("updateUser", controller.UpdateUser)
	routes.DELETE("/deleteUser/:userId", controller.DeleteUser)

}
