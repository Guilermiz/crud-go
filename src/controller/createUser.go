package controller

import (
	"github.com/Guilermiz/crud-go/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func CreateUser(context *gin.Context) {

	err := rest_err.NewUnprocessableEntityError("Erro na requisicao")
	context.JSON(err.Code, err)

}
