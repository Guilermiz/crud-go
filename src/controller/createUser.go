package controller

import (
	"github.com/Guilermiz/crud-go/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func CreateUser(context *gin.Context) {

	err := rest_err.NewBadRequestError("Bad Request", nil)
	context.JSON(err.Code, err)

}
