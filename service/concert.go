/*
 * @author: biturd
 * @date: 2024/5/12 09:11
 * @description:
 */
package service

import (
	"chaindraw-fair-ticket-backend/global"
	"chaindraw-fair-ticket-backend/model"
	commonreq "chaindraw-fair-ticket-backend/model/common/request"
	commonresp "chaindraw-fair-ticket-backend/model/common/response"
)

func ConcertList(ids []string) ([]commonresp.Concert, error) {
	res := make([]commonresp.Concert, 0)
	concerts := make([]model.TbConcert, 0)
	global.DB.Where("id IN ?", ids).Find(concerts)
	if global.DB.RowsAffected == 0 {
		return res, global.DB.Error
	}

	for _, concert := range concerts {
		res = append(res, commonresp.Concert{
			ConcertID: concert.ConcertID,
		})
	}
	return res, nil
}

func ConcertReview(concertViewRecord *commonreq.ConcertReViewReq) ([]commonresp.Concert, error) {
	res := make([]commonresp.Concert, 0)
	global.DB.Model(&model.TbConcert{}).Where("concert_id = ?", concertViewRecord.ConcertID).Update("concert_status", concertViewRecord.Pass)

	if global.DB.RowsAffected == 0 {
		return res, global.DB.Error
	}

	res = append(res, commonresp.Concert{
		ConcertID: concertViewRecord.ConcertID,
	})
	return res, nil
}
