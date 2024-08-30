/**
* @Author: cr-mao
* @Desc: error demo
**/
package controllers

import (
	"github.com/gin-gonic/gin"

	"goframe/app"
	"goframe/app/http/response"
)

type ErrorController struct{}

func (c *ErrorController) Demo1(ctx *gin.Context) {
	errService := app.ErrorService()
	err := errService.Demo1()
	response.WriteResponse(ctx, err, nil)
}
