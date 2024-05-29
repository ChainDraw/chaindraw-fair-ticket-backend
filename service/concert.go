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

	"gorm.io/gorm"
)

func ConcertList(ids []string, page, pageSize int) ([]commonresp.Concert, error) {
	res := make([]commonresp.Concert, 0)
	concerts := make([]model.TbConcert, 0)
	tickets := make([]model.TbTicket, 0)
	tx := global.DB.Debug()
	offset := (page - 1) * pageSize

	//if frontend give ids, then query database according to ids.
	//else query database is default according to offset+size .
	var resultConcerts *gorm.DB
	if len(ids) > 0 {
		resultConcerts = tx.Where("concert_id IN ?", ids).Find(&concerts)
	} else {
		//query some concert data according to limit and page size.
		resultConcerts = tx.Order("concert_date DESC").Limit(pageSize).Offset(offset).Find(&concerts)
	}

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

		// build TicketType list
		var ticketTypeList []commonresp.TicketType
		for _, ticket := range tickets {
			formattedPrice := fmt.Sprintf("%.2f", ticket.Price)
			priceFloat, _ := strconv.ParseFloat(formattedPrice, 64)

			if ticket.ConcertID == concert.ConcertID {
				ticketTypeList = append(ticketTypeList, commonresp.TicketType{
					TicketType:           ticket.TicketType,
					TypeName:             ticket.TypeName,
					Price:                priceFloat,
					MaxQuantityPerWallet: ticket.MaxQuantityPerWallet,
					CreateAt:             strconv.Itoa(int(ticket.CreateAt)),
					UpdateAt:             strconv.Itoa(int(ticket.UpdateAt)),
				})
			}
		}
		res = append(res, commonresp.Concert{
			ConcertID:     concert.ConcertID,
			ConcertName:   concert.ConcertName,
			ConcertDate:   formattedTime,
			ReviewStatus:  int(concert.ReviewStatus),
			ConcertStatus: int(concert.ConcertStatus),
			ConcertImg:    concert.ConcertImgURL,
			TicketTypes:   ticketTypeList,
			Status:        "scheduled",
		})
	}
	return res, nil
}

func ConcertReview(concertViewRecord *commonreq.ConcertReViewReq) ([]commonresp.Concert, error) {
	res := make([]commonresp.Concert, 0)
	result := global.DB.Debug().Model(&model.TbConcert{}).Where("concert_id = ?", concertViewRecord.ConcertID).Update("review_status", concertViewRecord.Pass)

	if result.RowsAffected == 0 {
		return res, result.Error
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
			Code:      200,
			Status:    "success",
			Msg:       "Concert was cancelled, tickets and deposits refunded",
			RequestID: concertCancelReq.ConcertID,
			Result: commonresp.CancelResult{
				ConcertID: concertCancelReq.ConcertID,
			},
		}
	} else {
		res = commonresp.ConcertCancellationResponse{
			Code:      200,
			Status:    "failed",
			Msg:       "Failed to cancel ticket",
			RequestID: concertCancelReq.ConcertID,
			Result: commonresp.CancelResult{
				ConcertID: concertCancelReq.ConcertID,
			},
		}
	}

	return res, nil
}
