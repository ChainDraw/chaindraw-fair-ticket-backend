/*
 * @author: ASWLaunchs
 * @date: 2024/5/12 09:02
 * @description:
 */
package service

import (
	"chaindraw-fair-ticket-backend/global"
	"chaindraw-fair-ticket-backend/model"
	commonreq "chaindraw-fair-ticket-backend/model/common/request"
	commonresp "chaindraw-fair-ticket-backend/model/common/response"
	"strconv"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

func ConcertAdd(concert *commonreq.ConcertAddReq) (resp *commonresp.CommitResp, err error) {
	//datetime convertion method must use time.RFC3339 .
	lotteryStartDate, _ := time.Parse(time.RFC3339, concert.LotteryStartDate)
	lotteryEndDate, _ := time.Parse(time.RFC3339, concert.LotteryEndDate)
	concertDate, _ := time.Parse(time.RFC3339, concert.ConcertDate)

	//if concert id is not existed then set a uuid value to it.
	if len(strings.TrimSpace(concert.ConcertID)) == 0 {
		u1, _ := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
		concert.ConcertID = u1.String()
	}

	//store concert.
	record := &model.TbConcert{
		ConcertID:        concert.ConcertID,
		ConcertName:      concert.ConcertName,
		Address:          concert.Address,
		Remark:           concert.Remark,
		ConcertDate:      concertDate.UnixMilli(),
		ConcertImgURL:    concert.ConcertImg,
		ConcertStatus:    int64(concert.ConcertStatus),
		ReviewStatus:     int64(concert.ReviewStatus),
		LotteryStartDate: lotteryStartDate.UnixMilli(),
		LotteryEndDate:   lotteryEndDate.UnixMilli(),
	}

	//set resp
	resp = &commonresp.CommitResp{}
	resp.Code = 200
	resp.Status = "success"
	resp.Msg = "Concert commit success!"
	resp.RequestID = concert.ConcertID
	resp.Result = &commonresp.ResultData{
		ConcertID: concert.ConcertID,
	}

	global.DB.Save(&record)
	err = global.DB.Error
	if err != nil {
		return resp, global.DB.Error
	}

	// 保存演唱会门票
	//store ticket.
	for _, ticket := range concert.Tickets {
		ticketType, _ := strconv.Atoi(ticket.TicketType)
		price, _ := strconv.ParseFloat(ticket.Price, 64)

		canTrade := 0
		if ticket.Trade {
			canTrade = 1
		}
		var ticketRecord = &model.TbTicket{
			ConcertID:            concert.ConcertID,
			TicketType:           int64(ticketType),
			TypeName:             ticket.TypeName,
			Num:                  int64(ticket.Num),
			Price:                price,
			TicketImg:            ticket.TicketImg,
			Trade:                int64(canTrade),
			MaxQuantityPerWallet: int64(ticket.MaxQuantityPerWallet),
			CreateAt:             time.Now().Unix(),
			UpdateAt:             time.Now().Unix(),
		}
		global.DB.Save(&ticketRecord)
		if global.DB.Error != nil {
			return resp, global.DB.Error
		}
	}

	return resp, nil
}
func ConcertStatusUpdate(concertId, reviewStatusStr, concertStatusStr string) (err error) {
	reviewStatus, _ := strconv.Atoi(reviewStatusStr)
	concertStatus, _ := strconv.Atoi(concertStatusStr)
	global.DB.Where("concert_id = ?", concertId).Updates(model.TbConcert{ReviewStatus: int64(reviewStatus), ConcertStatus: int64(concertStatus)})
	if global.DB.Error == nil {
		return err
	}

	return
}
