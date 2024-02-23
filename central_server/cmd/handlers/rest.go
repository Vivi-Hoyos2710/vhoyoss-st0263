package handlers

import (
	"errors"
	"github.com/Vivi-Hoyos2710/vhoyoss-st0263/central_server/internal/auth"
	"github.com/Vivi-Hoyos2710/vhoyoss-st0263/central_server/pkg/web"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiRest struct {
	service auth.Service
}

func NewApiRest(service auth.Service) *ApiRest {
	return &ApiRest{service: service}
}

func (a *ApiRest) SendIndex(c *gin.Context) {

}

func (a *ApiRest) Login(c *gin.Context) {
	var peer auth.Peer
	if err := c.ShouldBindJSON(&peer); err != nil {
		web.Error(c, http.StatusBadRequest, "invalid request body")
		return
	}
	loggedPeer, err := a.service.Login(peer)
	if err != nil {
		handleErrors(c, err)
		return
	}
	web.SuccessLogin(c, http.StatusOK, loggedPeer)

}
func handleErrors(c *gin.Context, err error) {
	switch {
	case errors.Is(err, auth.ErrExists):
		web.Error(c, http.StatusConflict, err.Error())
	default:
		web.Error(c, http.StatusInternalServerError, err.Error())

	}

}
