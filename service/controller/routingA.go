package controller

import (
	"v2rayA/common"
	"v2rayA/core/routingA"
	"v2rayA/persistence/configure"
	"github.com/gin-gonic/gin"
)

func GetRoutingA(ctx *gin.Context) {
	common.ResponseSuccess(ctx, gin.H{
		"routingA": configure.GetRoutingA(),
	})
}
func PutRoutingA(ctx *gin.Context) {
	var data struct {
		RoutingA string `json:"routingA"`
	}
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		common.ResponseError(ctx, logError(err, "bad request"))
		return
	}
	_, err = routingA.Parse(data.RoutingA)
	if err != nil {
		common.ResponseError(ctx, logError(err))
		return
	}
	err = configure.SetRoutingA(data.RoutingA)
	if err != nil {
		common.ResponseError(ctx, logError(err))
		return
	}
	common.ResponseSuccess(ctx, nil)
}
