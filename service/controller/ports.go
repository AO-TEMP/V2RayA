package controller

import (
	"v2rayA/common"
	"v2rayA/persistence/configure"
	"v2rayA/service"
	"github.com/gin-gonic/gin"
)

func PutPorts(ctx *gin.Context) {
	var data configure.Ports
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		common.ResponseError(ctx, logError(nil, "bad request"))
		return
	}
	err = service.SetPorts(&data)
	if err != nil {
		common.ResponseError(ctx, logError(err))
		return
	}
	common.ResponseSuccess(ctx, nil)
}

func GetPorts(ctx *gin.Context) {
	common.ResponseSuccess(ctx, service.GetPortsDefault())
}
