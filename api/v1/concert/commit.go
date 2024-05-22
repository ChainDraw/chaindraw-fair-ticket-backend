/*
 * @author: ASWLaunchs
 * @date: 2024/5/20 06:03
 * @description:
 */

package concert

import (
	"chaindraw-fair-ticket-backend/global"
	commonreq "chaindraw-fair-ticket-backend/model/common/request"
	commonresp "chaindraw-fair-ticket-backend/model/common/response"
	"chaindraw-fair-ticket-backend/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Summary Add a new concert
// @Description Add a new concert with the given details
// @Tags Concert
// @Accept  json
// @Produce  json
// @Param   concertAddReq  body    commonreq.ConcertAddReq  true  "Concert Add Request"
// @Success 200 {object} commonresp.CommitResp
// @Failure 400 {object} commonresp.CommitResp
// @Router /concert_add [post]
func ConcertAdd(ctx *gin.Context) {
	req := &commonreq.ConcertAddReq{}
	resp := &commonresp.CommitResp{}

	err := ctx.ShouldBindJSON(req)
	if err != nil {
		global.LOGGER.Info("ConcertAdd failed,bind json params failed", zap.Any("req", req))
		commonresp.FailWithMessage(ctx, "参数绑定失败")
		return
	}

	err = service.ConcertAdd(req)
	if err != nil {
		global.LOGGER.Error("ConcertAdd logic failed", zap.Any("req", req), zap.Error(err))
		commonresp.FailWithMessage(ctx, "抽票失败")
	}

	commonresp.OkWithData(ctx, resp)
}

// @Summary  update concert
// @Description update concert
// @Tags Concert
// @Accept  json
// @Produce  json
// @Param   ConcertId      query    int     true  "Comma-separated list of concert IDs"
// @Param   ConcertStatus    query    int        false "演唱会状态（0: 未开始, 1: 已过期, 2: 已取消）"
// @Param   ReviewStatus     query    int        false "审核状态（0: 未审核, 1: 审核通过, 2: 审核失败）"
// @Success 200 {object} commonresp.CommitResp
// @Failure 400 {object} commonresp.CommitResp
// @Router /concert_add [post]
func ConcertStatusUpdate(ctx *gin.Context) {
	req := &commonreq.ConcertAddReq{}
	resp := &commonresp.CommitResp{}

	concertId := ctx.Query("concert_id")
	reviewStatus := ctx.Query("review_status")
	concertStatus := ctx.Query("concert_status")

	err := service.ConcertStatusUpdate(concertId, reviewStatus, concertStatus)
	if err != nil {
		global.LOGGER.Error("ConcertAdd logic failed", zap.Any("req", req), zap.Error(err))
		commonresp.FailWithMessage(ctx, "抽票失败")
	}

	commonresp.OkWithData(ctx, resp)
}
