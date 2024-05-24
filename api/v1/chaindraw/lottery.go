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
	"fmt"
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

// @Summary Lottery list
// @Description GETLotterylist
// @Tags Lottery
// @Accept  json
// @Produce  json
// @Param   LotteryListReq  body    commonreq.LotteryListReq  true  "LotteryListRequest"
// @Success 200 {object} commonresp.LotteryListResponse
// @Failure 400 {object} commonresp.LotteryListResponse
// @Router /lottery/list [get]
func LotteryListGet(ctx *gin.Context) {
	req := &commonreq.LotteryListReq{}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		fmt.Println(err)
	}
	lotterys := make([]model.EventEscrowCreated, 0)
	global.DB.Where("concert_id = ?", req.ConcertID).Where("ticket_type = ?", req.TicketType).Find(&lotterys)
	resp := &commonresp.LotteryListResponse{}
	for _, lottery := range lotterys {
		resp.Result.LotteryList = append(resp.Result.LotteryList, lottery.EscrowAddress)
	}
	commonresp.OkWithData(ctx, resp)
}
