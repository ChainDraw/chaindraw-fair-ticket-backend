/*
 * @author: asta
 * @date: 2024/5/13 23:27
 * @description: concert controller
 *
 */
package chaindraw

import (
	"chaindraw-fair-ticket-backend/global"
	commonreq "chaindraw-fair-ticket-backend/model/common/request"
	commonresp "chaindraw-fair-ticket-backend/model/common/response"
	"chaindraw-fair-ticket-backend/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

func ConcertList(ctx *gin.Context) {
	resp := &commonresp.ConcertListResponse{}
	ids := strings.Split(ctx.Query("ids"), ",")
	concerts, err := service.ConcertList(ids)

	resp.Result = concerts
	if err != nil {
		global.LOGGER.Error("query concerts logic failed", zap.Any("ids", ids), zap.Error(err))
		ctx.JSON(http.StatusBadRequest, resp)
	}

	ctx.JSON(http.StatusOK, resp)
}

func ReviewConcert(ctx *gin.Context) {
	req := &commonreq.ConcertReViewReq{}
	resp := &commonresp.ConcertListResponse{}
	concerts, err := service.ConcertReview(req)

	resp.Result = concerts
	if err != nil {
		global.LOGGER.Error("review concerts logic failed", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, resp)
	}

	ctx.JSON(http.StatusOK, resp)
}
