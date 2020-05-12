package controller

import (
	"v2rayA/common"
	"v2rayA/core/touch"
	"v2rayA/persistence/configure"
	"v2rayA/service"
	"github.com/gin-gonic/gin"
)

/*修改Remarks*/
func PatchSubscription(ctx *gin.Context) {
	var data struct {
		Subscription touch.Subscription `json:"subscription"`
	}
	err := ctx.ShouldBindJSON(&data)
	s := data.Subscription
	index := s.ID - 1
	if err != nil || s.TYPE != configure.SubscriptionType || index < 0 || index >= configure.GetLenSubscriptions() {
		common.ResponseError(ctx, logError(nil, "bad request"))
		return
	}
	err = service.ModifySubscriptionRemark(s)
	if err != nil {
		common.ResponseError(ctx, logError(err))
		return
	}
	GetTouch(ctx)
}

/*更新订阅*/
func PutSubscription(ctx *gin.Context) {
	var data configure.Which
	err := ctx.ShouldBindJSON(&data)
	index := data.ID - 1
	if err != nil || data.TYPE != configure.SubscriptionType || index < 0 || index >= configure.GetLenSubscriptions() {
		common.ResponseError(ctx, logError(nil, "bad request"))
		return
	}
	err = service.UpdateSubscription(index, false)
	if err != nil {
		common.ResponseError(ctx, logError(err))
		return
	}
	GetTouch(ctx)
}
