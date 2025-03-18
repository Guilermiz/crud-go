package controller

import (
	"fmt"
	"github.com/Guilermiz/crud-go/src/configuration/validation"
	"github.com/Guilermiz/crud-go/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(context *gin.Context) {

	var userRequest request.UserRequest

	if err := context.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)
		context.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)

}
