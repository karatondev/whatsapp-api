package rest

import (
	"net/http"
	"strings"
	"whatsapp-api/internal/provider"
	"whatsapp-api/internal/service"
	"whatsapp-api/model"
	"whatsapp-api/model/constant"
	"whatsapp-api/util"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	log     provider.ILogger
	service service.MessagesApi
}

func NewRest(log provider.ILogger, service service.MessagesApi) *Rest {
	return &Rest{log: log, service: service}
}

func (rs *Rest) CreateServer(address string) (*http.Server, error) {
	gin.SetMode(util.Configuration.Server.Mode)

	r := gin.New()
	r.GET("/ping", rs.checkConnectivity)
	r.POST(util.Configuration.Server.Path.Messages, rs.pushMessage)
	r.GET(util.Configuration.Server.Path.Contacts, rs.getContacts)
	r.GET(util.Configuration.Server.Path.Groups, rs.getGroups)

	server := &http.Server{
		Addr:    address,
		Handler: r,
	}

	return server, nil
}

func (rs *Rest) checkConnectivity(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (rs *Rest) pushMessage(ctx *gin.Context) {

	msg := model.MessageRequest{}
	if err := ctx.ShouldBindJSON(&msg); err != nil {
		errResp := model.ErrorResponse{
			Error: model.Error{
				Code:    constant.ValidationError,
				Message: err.Error(),
			},
		}
		ctx.JSON(http.StatusInternalServerError, errResp)
		return
	}

	message, err := rs.service.PushMessage(ctx, &msg)
	if err != nil {
		code, errResp := ErrorResponseMap(err)
		ctx.JSON(code, errResp)
		return
	}

	ctx.JSON(http.StatusOK, message)
}

func (rs *Rest) getContacts(ctx *gin.Context) {

	var deviceId = strings.ToLower(ctx.Param("id"))
	contacts, err := rs.service.GetContacts(ctx, deviceId)
	if err != nil {
		code, errResp := ErrorResponseMap(err)
		ctx.JSON(code, errResp)
		return
	}

	ctx.JSON(http.StatusOK, contacts)
}

func (rs *Rest) getGroups(ctx *gin.Context) {

	var deviceId = strings.ToLower(ctx.Param("id"))
	groups, err := rs.service.GetGroups(ctx, deviceId)
	if err != nil {
		code, errResp := ErrorResponseMap(err)
		ctx.JSON(code, errResp)
		return
	}

	ctx.JSON(http.StatusOK, groups)
}
