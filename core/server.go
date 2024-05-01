package core

import (
	"chaindraw-fair-ticket-backend/global"
	"chaindraw-fair-ticket-backend/router"

	"github.com/fvbock/endless"
)

func RunServer() {
	Router := router.Router()

	if err := endless.ListenAndServe(global.APP_CONFIG.Server.Port, Router); err != nil {
		global.LOGGER.Error(err.Error())
		return
	}
}
