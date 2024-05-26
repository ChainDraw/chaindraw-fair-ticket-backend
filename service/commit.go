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
	"strconv"
	"time"
)

func ConcertAdd(concert *commonreq.ConcertAddReq) (err error) {
	//2024-05-20 待解决，演唱会主办方提交信息Req涉及到tb_concert和tb_ticket共2个表的数据存储，需要考虑到ticket结构体数组的问题，Req参考的是文档 by ASWLauncher
	//日期转换
	lotteryStartDate, _ := time.Parse("2006-01-02 15:04:05", concert.LotteryStartDate)
	lotteryEndDate, _ := time.Parse("2006-01-02 15:04:05", concert.LotteryEndDate)
	concertDate, _ := time.Parse("2006-01-02 15:04:05", concert.ConcertDate)

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

	global.DB.Save(&record)
	err = global.DB.Error
	if err != nil {
		return
	}

	// 保存演唱会门票
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
			return
		}

	}
	return
}
func ConcertStatusUpdate(concertId, reviewStatusStr, concertStatusStr string) (err error) {
	//2024-05-20 待解决，演唱会主办方提交信息Req涉及到tb_concert和tb_ticket共2个表的数据存储，需要考虑到ticket结构体数组的问题，Req参考的是文档 by ASWLauncher
	reviewStatus, _ := strconv.Atoi(reviewStatusStr)
	concertStatus, _ := strconv.Atoi(concertStatusStr)
	global.DB.Where("concert_id = ?", concertId).Updates(model.TbConcert{ReviewStatus: int64(reviewStatus), ConcertStatus: int64(concertStatus)})
	if global.DB.Error == nil {
		return err
	}

	return
}
