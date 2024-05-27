/*
 * @author: biturd
 * @date: 2024/5/8 01:32
 * @description:
 */
package chaindraw

import (
	"chaindraw-fair-ticket-backend/global"
	commonreq "chaindraw-fair-ticket-backend/model/common/request"
	commonresp "chaindraw-fair-ticket-backend/model/common/response"
	model "chaindraw-fair-ticket-backend/model/event"
	"chaindraw-fair-ticket-backend/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LotteryRecordAdd(ctx *gin.Context) {
	req := &commonreq.LotteryRecordReq{}
	resp := &commonresp.LotteryRecordResponse{}

	err := ctx.ShouldBindJSON(req)
	if err != nil {
		global.LOGGER.Info("LotteryRecordAdd failed,bind json params failed", zap.Any("req", req))
		commonresp.FailWithMessage(ctx, "参数绑定失败")
		return
	}

	err = service.LotteryRecordAdd(req)
	if err != nil {
		global.LOGGER.Error("LotteryRecordAdd logic failed", zap.Any("req", req), zap.Error(err))
		commonresp.FailWithMessage(ctx, "抽票失败")
	}

	commonresp.OkWithData(ctx, resp)
}

// @Summary Get Lottery Address
// @Description GETLotterylist
// @Tags Lottery
// @Accept  json
// @Produce  json
// @Param   concert_id      query    string     false  " concert_id"
// @Param   ticket_type     query    string     false   "ticket_type"
// @Success 200 {object} commonresp.LotteryListResponse
// @Failure 400 {object} commonresp.LotteryListResponse
// @Router /lottery/list [get]
func LotteryListGet(ctx *gin.Context) {
	conertId, existC := ctx.GetQuery("concert_id")
	ticket_type, existT := ctx.GetQuery("ticket_type")
	if !existC || !existT {
		global.LOGGER.Info("LotteryListGet failed, get query params failed", zap.Any("req", gin.H{"concert_id": conertId,
			"ticket_type": ticket_type}))
		commonresp.FailWithMessage(ctx, "参数获取失败")
		return
	}
	lotterys := make([]model.EventEscrowCreated, 0)
	global.DB.Where("concert_id = ?", conertId).Where("ticket_type = ?", ticket_type).Find(&lotterys)
	resp := &commonresp.LotteryListResponse{}
	for _, lottery := range lotterys {
		resp.Result.LotteryList = append(resp.Result.LotteryList, lottery.EscrowAddress)
	}
	commonresp.OkWithData(ctx, resp)
}
