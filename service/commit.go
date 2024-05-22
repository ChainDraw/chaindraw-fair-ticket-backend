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

	record := &model.TbConcert{
		ConcertID:   concert.ConcertID,
		ConcertName: concert.ConcertName,
		//Address:       concert.Address,
		//Remark:        concert.Remark,
		ConcertDate:   lotteryStartDate.UnixMilli(),
		ConcertImgURL: concert.ConcertImg,
		ConcertStatus: concert.ConcertStatus,
		Status:        concert.ReviewStatus,
		//Tickets:       concert.Tickets,
	}
	global.DB.Save(record)
	err = global.DB.Error
	if err != nil {
		return
	}
	return
}

func ConcertStatusUpdate(concertId, reviewStatusStr, concertStatusStr string) (err error) {
	//2024-05-20 待解决，演唱会主办方提交信息Req涉及到tb_concert和tb_ticket共2个表的数据存储，需要考虑到ticket结构体数组的问题，Req参考的是文档 by ASWLauncher
	reviewStatus,_ := strconv.Atoi(reviewStatusStr)
	concertStatus,_ := strconv.Atoi(concertStatusStr)
	global.DB.Where("concert_id = ?", concertId).Updates(model.TbConcert{Status: reviewStatus, ConcertStatus: concertStatus})
	if global.DB.Error == nil {
		return err
	}

	return
}
