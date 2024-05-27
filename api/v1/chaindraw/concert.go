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
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Summary Get list of concerts
// @Description Get list of concerts by IDs with pagination
// @Tags Concert
// @Accept  json
// @Produce  json
// @Param   ids      query    string     false  "Comma-separated list of concert IDs"
// @Param   page     query    int        false "Page number"
// @Param   page_size query   int        false "Number of items per page"
// @Success 200 {object} commonresp.ConcertListResponse
// @Failure 400 {object} commonresp.ConcertListResponse
// @Router /concert/concert_list [get]
func ConcertList(ctx *gin.Context) {
	resp := &commonresp.ConcertListResponse{}
	pageStr := ctx.DefaultQuery("page", "1")
	pageSizeStr := ctx.DefaultQuery("page_size", "10")
	ids := strings.Split(ctx.Query("ids"), ",")
	page, _ := strconv.Atoi(pageStr)         // 当前页码
	pageSize, _ := strconv.Atoi(pageSizeStr) // 每页记录数

	concerts, err := service.ConcertList(ids, page, pageSize)

	resp.Result = concerts
	if err != nil {
		global.LOGGER.Error("query concerts logic failed", zap.Any("ids", ids), zap.Error(err))
		ctx.JSON(http.StatusBadRequest, resp)
	}

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Review concerts
// @Description Review concerts
// @Tags Concert
// @Accept  json
// @Produce  json
// @Param   concertReviewReq  body    commonreq.ConcertReViewReq  true  "Concert Review Request"
// @Success 200 {object} commonresp.ConcertListResponse
// @Failure 400 {object} commonresp.ConcertListResponse
// @Router /concert/review [post]
func ReviewConcert(ctx *gin.Context) {
	req := &commonreq.ConcertReViewReq{}
	resp := &commonresp.ConcertListResponse{}

	err := ctx.ShouldBindJSON(req)
	if err != nil {
		global.LOGGER.Info("ReviewConcert failed, bind json params failed", zap.Any("req", req))
		commonresp.FailWithMessage(ctx, "参数绑定失败")
		return
	}

	concerts, err := service.ConcertReview(req)
	resp.Result = concerts
	if err != nil {
		global.LOGGER.Error("review concerts logic failed", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, resp)
	}

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Cancel concerts
// @Description Cancel concerts
// @Tags Concert
// @Accept  json
// @Produce  json
// @Param   concertCancelReq  body    commonreq.ConcertCancelReq  true  "Concert Cancel Request"
// @Success 200 {object} commonresp.ConcertListResponse
// @Failure 400 {object} commonresp.ConcertListResponse
// @Router /concert/cancel [post]
func CancelConcert(ctx *gin.Context) {
	req := &commonreq.ConcertCancelReq{}
	resp := &commonresp.ConcertCancellationResponse{}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		global.LOGGER.Info("CancelConcert failed, bind json params failed", zap.Any("req", req))
		commonresp.FailWithMessage(ctx, "参数绑定失败")
		return
	}

	concerts, err := service.ConcertCancel(req)
	resp = &concerts
	if err != nil {
		global.LOGGER.Error("review concerts logic failed", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, resp)
	}

	ctx.JSON(http.StatusOK, resp)
}
