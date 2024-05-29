/*
 * @author: biturd
 * @date: 2024/5/12 09:02
 * @description:
 */
package service

import (
	"chaindraw-fair-ticket-backend/global"
	"chaindraw-fair-ticket-backend/model"
	commonreq "chaindraw-fair-ticket-backend/model/common/request"
)

func LotteryRecordAdd(lotteryRecord *commonreq.LotteryRecordReq) (err error) {
	record := &model.TbTicket{
		ConcertID: lotteryRecord.ConcertID,
		TypeName:  lotteryRecord.TypeName,
		//TicketType: lotteryRecord.TicketType,
		//Quantity:   lotteryRecord.Quantity,
		//WalletAddress: lotteryRecord.WalletAddress,
	}
	global.DB.Save(record)
	err = global.DB.Error
	if err != nil {
		return
	}
	return
}
