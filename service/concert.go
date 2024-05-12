/*
 * @author: biturd
 * @date: 2024/5/12 09:11
 * @description:
 */
package service

import (
	"chaindraw-fair-ticket-backend/global"
	"chaindraw-fair-ticket-backend/model"
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
