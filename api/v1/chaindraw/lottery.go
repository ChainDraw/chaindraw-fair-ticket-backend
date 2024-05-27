/*
 * @author: biturd
 * @date: 2024/5/8 01:32
 * @description:
 */
package chaindraw

import (
	"chaindraw-fair-ticket-backend/global"
	"chaindraw-fair-ticket-backend/model"
	commonreq "chaindraw-fair-ticket-backend/model/common/request"
	commonresp "chaindraw-fair-ticket-backend/model/common/response"
	event "chaindraw-fair-ticket-backend/model/event"
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

// @Summary ticket list
// @Description GETTicketlist
// @Tags Lottery
// @Param   lottery_address  query    string  true  "Lottery Address"
// @Success 200 {object} commonresp.TicketListResponse
// @Failure 400 {object} commonresp.TicketListResponse
// @Router /lottery/ticketList [get]
func TicketListGet(ctx *gin.Context) {
	lotteryAddress := ctx.Query("lottery_address")
	resp := &commonresp.TicketListResponse{}
	// Step 1: Query event_nft_listed table
	tickets := make([]event.EventNftListed, 0)
	result := global.DB.Where("lottery_address = ?", lotteryAddress).Find(&tickets)
	if result.Error != nil {
		commonresp.FailWithMessage(ctx, "查询 event_nft_listed 表时出错")
		return
	}

	if len(tickets) == 0 {
		commonresp.FailWithMessage(ctx, "未找到相关门票")
		return
	}
	// Step 2: Query event_escrow_created table using the first ticket's lottery_address
	var escrow event.EventEscrowCreated
	result = global.DB.Where("escrow_address = ?", lotteryAddress).First(&escrow)
	if result.Error != nil {
		commonresp.FailWithMessage(ctx, "查询 event_escrow_created 表时出错")
		return
	}

	// Step 3: Query tb_concert table using concert_id from the event_escrow_created
	var concert *model.TbConcert
	result = global.DB.Where("concert_id = ?", escrow.ConcertId).First(&concert)
	if result.Error != nil {
		commonresp.FailWithMessage(ctx, "查询 tb_concert 表时出错")
		return
	}

	// Step 4: Query tb_ticket table using ticket_type from the event_escrow_created
	var ticketType model.TbTicket
	result = global.DB.Where("ticket_type = ?", escrow.TicketType).First(&ticketType)
	if result.Error != nil {
		commonresp.FailWithMessage(ctx, "查询 tb_ticket 表时出错")
		return
	}

	// Build response
	resp.Result = make([]commonresp.SoldTicket, 0)
	for _, ticket := range tickets {
		resp.Result = append(resp.Result, commonresp.SoldTicket{
			ConcertName: concert.ConcertName,
			TypeName:    ticketType.TypeName,
			Url:         ticketType.TicketImg,
			Price:       ticketType.Price,
			Seller:      ticket.Seller,
			TokenID:     ticket.TokenId,
		})
	}

	commonresp.OkWithData(ctx, resp)

}
