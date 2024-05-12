/*
 * @author: biturd
 * @date: 2024/5/8 01:32
 * @description:
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
)

func LotteryRecordAdd(ctx *gin.Context) {
	req := &commonreq.LotteryRecordReq{}
	resp := &commonresp.LotteryRecordResponse{}

	err := ctx.ShouldBindJSON(req)
	if err != nil {
		global.LOGGER.Info("LotteryRecordAdd failed,bind json params failed", zap.Any("req", req))
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	err = service.LotteryRecordAdd(req)
	if err != nil {
		global.LOGGER.Error("LotteryRecordAdd logic failed", zap.Any("req", req), zap.Error(err))
		ctx.JSON(http.StatusBadRequest, resp)
	}

	ctx.JSON(http.StatusOK, resp)
}
