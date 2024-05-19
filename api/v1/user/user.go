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
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/spruceid/siwe-go"
	"go.uber.org/zap"
	"net/http"
	"time"
)

var Store = sessions.NewCookieStore([]byte("siwe-quickstart-secret"))

func getSession(c *gin.Context) *sessions.Session {
	session, _ := Store.Get(c.Request, "siwe-quickstart")
	return session
}

// @Summary Generate nonce
// @Description Generate nonce for session
// @Produce json
// @Success 200 {string} string "Nonce generated successfully"
// @Router /user/nonce [get]
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

type RequestBody struct {
	Message   string `json:"message"`
	Signature string `json:"signature"`
}

// @Summary Verify signature
// @Description Verify signature with message and nonce
// @Produce json
// @Accept json
// @Param request body RequestBody true "Request body"
// @Success 200 {boolean} boolean "Verification success"
// @Failure 400 {string} string "Verification failed"
// @Router /user/verify [post]
func Verify(c *gin.Context) {
	session := getSession(c)
	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		commonresp.FailWithMessage(c, fmt.Sprintf("Invalid request body", err.Error()))
		return
	}

	siweObj, err := siwe.ParseMessage(requestBody.Message)
	if err != nil {
		global.LOGGER.Info("siwe ParseMessage failed")
		commonresp.FailWithMessage(c, err.Error())
		return
	}

	nonce, ok := session.Values["nonce"].(string)
	if !ok {
		global.LOGGER.Info("session get nonce failed")
		commonresp.FailWithMessage(c, "session get nonce failed")
		return
	}
	_, err = siweObj.Verify(requestBody.Signature, nil, &nonce, nil)
	if err != nil {
		global.LOGGER.Info("siweObj verify failed")
		commonresp.FailWithMessage(c, "siweObj verify failed")
		return
	}

	msgJSON := commonreq.SiweMessage{
		Domain:         siweObj.GetDomain(),
		Address:        siweObj.GetAddress(),
		Uri:            siweObj.GetURI(),
		Version:        siweObj.GetVersion(),
		Statement:      siweObj.GetStatement(),
		Nonce:          siweObj.GetNonce(),
		ChainID:        siweObj.GetChainID(),
		IssuedAt:       siweObj.GetIssuedAt(),
		ExpirationTime: siweObj.GetExpirationTime(),
		NotBefore:      siweObj.GetNotBefore(),
		RequestID:      siweObj.GetRequestID(),
		Resources:      siweObj.GetResources(),
	}

	jsonString, err := json.Marshal(msgJSON)
	if err != nil {
		global.LOGGER.Error("siweObj transfer JSON failed", zap.Error(err))
		return
	}
	session.Values["siwe"] = string(jsonString)

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

// @Summary Get personal information
// @Description Get personal information
// @Produce json
// @Success 200 {string} string "Personal information retrieved successfully"
// @Failure 401 {string} string "Unauthorized"
// @Router /user/personal_information [get]
func PersonalInformation(c *gin.Context) {
	session := getSession(c)
	siweData := session.Values["siwe"]
	if siweData == nil {
		commonresp.FailWithMessage(c, "You have to first sign_in")
		return
	}
	siweDatastr := siweData.(string)
	siweResp := &commonreq.SiweMessage{}
	err := json.Unmarshal([]byte(siweDatastr), siweResp)
	if err != nil {
		global.LOGGER.Error("siwe data unmarshal failed", zap.Error(err))
		commonresp.FailWithMessage(c, "siwe data unmarshal failed")
		return
	}
	commonresp.OkWithData(c, siweResp)
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

func SessionDemo(ctx *gin.Context) {
	session := getSession(ctx)
	ty := ctx.Query("type")
	if ty == "0" { // 存
		// 存储一个值到会话中
		session.Values["key"] = "value"
		session.Save(ctx.Request, ctx.Writer)
	} else { // 取
		// 从会话中检索一个值
		value := session.Values["key"].(string)
		fmt.Printf("Value: %s", value)
	}
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
