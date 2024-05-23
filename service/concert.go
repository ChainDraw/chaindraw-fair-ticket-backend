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
	"time"
)

func ConcertList(ids []string, page, pageSize int) ([]commonresp.Concert, error) {
	res := make([]commonresp.Concert, 0)
	concerts := make([]model.TbConcert, 0)
	tx := global.DB.Debug()
	offset := (page - 1) * pageSize
	if len(ids) == 0 {
		tx.Where("id IN (?)", ids)
		tx.Model(&model.TbConcert{}).Where("id IN ?", ids)
	}
	tx.Order("concert_date DESC").Limit(pageSize).Offset(offset).Find(&concerts)

	tx.Order("concert_date DESC").Limit(pageSize).Offset(offset).Find(&concerts)
	if global.DB.RowsAffected == 0 {
		return res, global.DB.Error
	}

	for _, concert := range concerts {
		t := time.Unix(concert.ConcertDate, 0)
		formattedTime := t.Format("2006-01-02 15:04:05")
		res = append(res, commonresp.Concert{
			ConcertID:   concert.ConcertID,
			ConcertName: concert.ConcertName,
			ConcertDate: formattedTime,
			ConcertImg:  concert.ConcertImgURL,
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
