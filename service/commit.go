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
)

func ConcertAdd(concert *commonreq.ConcertAddReq) (err error) {
	//2024-05-20 待解决，演唱会主办方提交信息Req涉及到tb_concert和tb_ticket共2个表的数据存储，需要考虑到ticket结构体数组的问题，Req参考的是文档 by ASWLauncher
	record := &model.TbConcert{
		ConcertID:   concert.ConcertID,
		ConcertName: concert.ConcertName,
		//Address:       concert.Address,
		//Remark:        concert.Remark,
		//ConcertDate:   concert.ConcertDate,
		ConcertImgURL: concert.ConcertImg,
		ConcertStatus: int64(concert.ConcertStatus),
		//ReviewStatus:  concert.ReviewStatus,
		//Tickets:       concert.Tickets,
	}
	global.DB.Save(record)
	err = global.DB.Error
	if err != nil {
		return
	}
	return
}
