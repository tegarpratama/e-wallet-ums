package api

import (
	"ewallet-ums/constant"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LogoutHandler struct {
	LogoutService interfaces.ILogoutService
}

func (api *LogoutHandler) Logout(c *gin.Context) {
	log.Println("API EXECUTED")
	var (
		log = helpers.Logger
	)

	token := c.Request.Header.Get("Authorization")
	err := api.LogoutService.Logout(c, token)
	if err != nil {
		log.Error("failed on logout service: ", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constant.ErrServerError, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constant.SuccessMessage, nil)
}
