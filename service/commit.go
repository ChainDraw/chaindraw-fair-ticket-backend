/*
 * @author: ASWLaunchs
 * @date: 2024/5/12 09:02
 * @description:
 */
package service

import (
	commonreq "chaindraw-fair-ticket-backend/model/common/request"
)

func ConcertAdd(concert *commonreq.ConcertAddReq) (err error) {
	//2024-05-20 暂时别用，演唱会主办方提交信息Req和数据库TbConcert不同，Req参考的是文档 by ASWLauncher
	// record := &model.TbConcert{
	// 	ConcertID:        concert.ConcertID,
	// 	ConcertName:      concert.ConcertName,
	// 	Address:          concert.Address,
	// 	Remark:           concert.Remark,
	// 	LotteryStartDate: concert.LotteryStartDate,
	// 	LotteryEndDate:   concert.LotteryEndDate,
	// 	ConcertDate:      concert.ConcertDate,
	// 	ConcertImg:       concert.ConcertImg,
	// 	ConcertStatus:    concert.ConcertStatus,
	// 	ReviewStatus:     concert.ReviewStatus,
	// 	Tickets:          concert.Tickets,
	// }
	// global.DB.Save(record)
	// err = global.DB.Error
	// if err != nil {
	// 	return
	// }
	return
}
