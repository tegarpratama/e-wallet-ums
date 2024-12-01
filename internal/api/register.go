package api

import (
	"ewallet-ums/constant"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterHandler struct {
	RegisterService interfaces.RegisterService
}

func (api *RegisterHandler) Register(c *gin.Context) {
	var (
		log = helpers.Logger
	)
	req := models.User{}

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("failed to parse request: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constant.ErrFailedBadRequest, nil)
		return
	}

	if err := req.Validate(); err != nil {
		log.Error("failed to validate request: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constant.ErrFailedBadRequest, nil)
		return
	}

	resp, err := api.RegisterService.Register(c.Request.Context(), req)
	if err != nil {
		log.Error("failed to register new user: ", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constant.ErrServerError, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusCreated, constant.SuccessMessage, resp)
}
