/*
 * @author: biturd
 * @date: 2024/5/12 09:00
 * @description:
 */
package chaindraw

import (
	"chaindraw-fair-ticket-backend/global"
	commonresp "chaindraw-fair-ticket-backend/model/common/response"
	"chaindraw-fair-ticket-backend/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func ConcertList(ctx *gin.Context) {
	resp := &commonresp.ConcertListResponse{}
	ids := ctx.Query("ids")

	err := service.ConcertList(ids)
	if err != nil {
		global.LOGGER.Error("query concerts logic failed", zap.String("ids", ids), zap.Error(err))
		ctx.JSON(http.StatusBadRequest, resp)
	}

	ctx.JSON(http.StatusOK, resp)
}
