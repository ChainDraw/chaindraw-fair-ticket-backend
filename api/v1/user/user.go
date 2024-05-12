/*
 * @author: biturd
 * @date: 2024/5/8 01:32
 * @description:
 */
package user

import (
	"chaindraw-fair-ticket-backend/global"
	commonreq "chaindraw-fair-ticket-backend/model/common/request"
	commonresp "chaindraw-fair-ticket-backend/model/common/response"
	"chaindraw-fair-ticket-backend/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/spruceid/siwe-go"
	"go.uber.org/zap"
	"net/http"
	"time"
)

var store = sessions.NewCookieStore([]byte("siwe-quickstart-secret"))

func getSession(c *gin.Context) *sessions.Session {
	session, _ := store.Get(c.Request, "siwe-quickstart")
	return session
}

func Nonce(c *gin.Context) {
	session := getSession(c)
	nonce := siwe.GenerateNonce()
	session.Values["nonce"] = nonce
	err := session.Save(c.Request, c.Writer)
	if err != nil {
		global.LOGGER.Info("session save failed")
		commonresp.FailWithMessage(c, "session save failed")
		return
	}
	commonresp.OkWithData(c, nonce)
}

func Verify(c *gin.Context) {
	session := getSession(c)
	message := c.PostForm("message")
	signature := c.PostForm("signature")

	siweObj, err := siwe.ParseMessage(message)
	if err != nil {
		global.LOGGER.Info("siwe ParseMessage failed")
		commonresp.FailWithMessage(c, err.Error())
		return
	}
	var domain = "example.com"
	nonce, ok := session.Values["nonce"].(string)
	if !ok {
		global.LOGGER.Info("session get nonce failed")
		commonresp.FailWithMessage(c, "session get nonce failed")
		return
	}
	pubKey, err := siweObj.Verify(signature, &domain, &nonce, nil)
	if err != nil {
		global.LOGGER.Info("siweObj verify failed")
		commonresp.FailWithMessage(c, "siweObj verify failed")
		return
	}
	session.Values["siwe"] = pubKey
	session.Options = &sessions.Options{MaxAge: 0, Path: "/", HttpOnly: true, Secure: false, SameSite: 0}

	expireTimeStr := *siweObj.GetExpirationTime()
	expireTime, err := time.Parse(time.RFC3339, expireTimeStr)
	if err != nil {
		global.LOGGER.Info("parse expireTimeStr failed")
		commonresp.FailWithMessage(c, "parse expireTimeStr failed")
		return
	}

	issuedAtStr := siweObj.GetIssuedAt()
	issuedAt, err := time.Parse(time.RFC3339, issuedAtStr)
	if err != nil {
		global.LOGGER.Info("parse issuedAtTimeStr failed")
		commonresp.FailWithMessage(c, "parse issuedAtTimeStr failed")
		return
	}

	maxAge := int(expireTime.Sub(issuedAt).Seconds())
	session.Options.MaxAge = maxAge

	err = session.Save(c.Request, c.Writer)
	if err != nil {
		global.LOGGER.Info("session get nonce failed")
		commonresp.FailWithMessage(c, "session get nonce failed")
		return
	}
	commonresp.OkWithData(c, true)
}

func PersonalInformation(c *gin.Context) {
	session := getSession(c)
	siweData := session.Values["siwe"]
	if siweData == nil {
		commonresp.FailWithMessage(c, "You have to first sign_in")
		return
	}

	commonresp.OkWithData(c, fmt.Sprintf("You are authenticated and your address is: %s", siweData.(string)))
}
func Login(ctx *gin.Context) {
	req := &commonreq.LoginReq{}
	resp := &commonresp.LoginResp{}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		global.LOGGER.Info("login failed,bind json params failed", zap.Any("req", req))
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	err = service.Login(req)
	if err != nil {
		global.LOGGER.Error("login logic failed", zap.Any("req", req), zap.Error(err))
		ctx.JSON(http.StatusBadRequest, resp)
	}

	ctx.JSON(http.StatusOK, resp)
}

func Register(ctx *gin.Context) {
	req := &commonreq.RegisterReq{}
	resp := &commonresp.RegisterResp{}

	err := ctx.ShouldBindJSON(req)
	if err != nil {
		global.LOGGER.Info("Register failed,bind json params failed", zap.Any("req", req))
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	err = service.Register(req)
	if err != nil {
		global.LOGGER.Error("Register logic failed", zap.Any("req", req), zap.Error(err))
		ctx.JSON(http.StatusBadRequest, resp)
	}

	ctx.JSON(http.StatusOK, resp)
}
