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
	"fmt"
	"strconv"
	"time"
)

func ConcertList(ids []string, page, pageSize int) ([]commonresp.Concert, error) {
	res := make([]commonresp.Concert, 0)
	concerts := make([]model.TbConcert, 0)
	tickets := make([]model.TbTicket, 0)
	tx := global.DB.Debug()
	offset := (page - 1) * pageSize
	if len(ids) == 0 {
		tx.Model(&model.TbConcert{}).Where("id IN ?", ids)
	}

	//query some concert data according to limit and page size.
	resultConcerts := tx.Order("concert_date DESC").Limit(pageSize).Offset(offset).Find(&concerts)

	if resultConcerts.RowsAffected == 0 {
		return res, global.DB.Error
	} else {
		concertIds := make([]string, 0)
		for _, v := range concerts {
			concertIds = append(concertIds, v.ConcertID)
		}
		tx.Where("concert_id IN (?)", concertIds).Order("create_at DESC").Limit(pageSize).Offset(offset).Find(&tickets)
	}

	for _, concert := range concerts {
		t := time.Unix(concert.ConcertDate/1000, 0)
		formattedTime := t.Format(time.RFC3339)
		// 查询相应的门票类型信息
		var ticketTypes []model.TbTicket
		ticketResult := global.DB.Where("concert_id = ?", concert.ConcertID).Find(&ticketTypes)
		if ticketResult.Error != nil {
			return res, ticketResult.Error
		}

		// 构建 TicketType 列表
		var ticketTypeList []commonresp.TicketType
		for _, ticket := range ticketTypes {
			formattedPrice := fmt.Sprintf("%.2f", ticket.Price)
			priceFloat, _ := strconv.ParseFloat(formattedPrice, 64)
			// 现在 priceFloat 是保留两位小数的 float64 类型

			ticketTypeList = append(ticketTypeList, commonresp.TicketType{
				TicketType:           ticket.TicketType,
				TypeName:             ticket.TypeName,
				Price:                priceFloat,
				MaxQuantityPerWallet: ticket.MaxQuantityPerWallet,
				CreateAt:             ticket.CreateAt,
				UpdateAt:             ticket.UpdateAt,
			})
		}
		res = append(res, commonresp.Concert{
			ConcertID:   concert.ConcertID,
			ConcertName: concert.ConcertName,
			ConcertDate: formattedTime,
			ConcertImg:  concert.ConcertImgURL,
			TicketTypes: ticketTypeList,
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

func ConcertCancel(concertCancelReq *commonreq.ConcertCancelReq) (commonresp.ConcertCancellationResponse, error) {
	result := global.DB.Model(&model.TbConcert{}).Where("concert_id = ?", concertCancelReq.ConcertID).Updates(map[string]interface{}{
		"concert_status": 1,
		"cancel_reason":  concertCancelReq.CancelReason,
	})

	var res commonresp.ConcertCancellationResponse
	if result.RowsAffected > 0 {
		res = commonresp.ConcertCancellationResponse{
			Status:    "success",
			Msg:       "Concert was cancelled, tickets and deposits refunded",
			RequestID: concertCancelReq.ConcertID,
			Result: commonresp.CancelResult{
				ConcertID: concertCancelReq.ConcertID,
			},
		}
	} else {
		res = commonresp.ConcertCancellationResponse{
			Status:    "failed",
			Msg:       "failed to cancel ticket",
			RequestID: concertCancelReq.ConcertID,
			Result: commonresp.CancelResult{
				ConcertID: concertCancelReq.ConcertID,
			},
		}
	}

	return res, nil
}
