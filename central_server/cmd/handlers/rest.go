package handlers

import (
	"errors"
	"net/http"

	"github.com/Vivi-Hoyos2710/vhoyoss-st0263/central_server/internal/auth"
	"github.com/Vivi-Hoyos2710/vhoyoss-st0263/central_server/internal/directory"
	"github.com/Vivi-Hoyos2710/vhoyoss-st0263/central_server/pkg/web"
	"github.com/gin-gonic/gin"
)

type ApiRest struct {
	authService      auth.Service
	directoryService directory.ServiceDirectory
}

func NewApiRest(auth auth.Service, dir directory.ServiceDirectory) *ApiRest {
	return &ApiRest{
		authService:      auth,
		directoryService: dir,
	}
}
func (a *ApiRest) GetIndexTable(c *gin.Context) {
	indexTable := a.directoryService.GetIndexTable()
	web.Success(c, http.StatusOK, indexTable)
}
func (a *ApiRest) SendIndex(c *gin.Context) {
	var indexInfo directory.Index
	if err := c.ShouldBindJSON(&indexInfo); err != nil {
		web.Error(c, http.StatusBadRequest, "invalid request body")
		return
	}
	err := a.directoryService.SendIndex(indexInfo)
	if err != nil {
		handleErrors(c, err)
		return
	}
	web.Success(c, http.StatusOK, "index sent successfully")

}

func (a *ApiRest) Login(c *gin.Context) {
	var peer auth.Peer
	if err := c.ShouldBindJSON(&peer); err != nil {
		web.Error(c, http.StatusBadRequest, "invalid request body")
		return
	}
	loggedPeer, _ := a.authService.Login(peer)
	
	web.SuccessLogin(c, http.StatusOK, loggedPeer)

}
func (a *ApiRest) Logout(c *gin.Context) {
	var peer auth.PeerLogOut
	if err := c.ShouldBindJSON(&peer); err != nil {
		web.Error(c, http.StatusBadRequest, "invalid request body")
		return
	}
	err := a.authService.Logout(peer)
	if err != nil {
		handleErrors(c, err)
		return
	}
	web.Success(c, http.StatusOK, "logged out successfully")

}
//Query Call to the directory service to get the location of the file, for fucntion DownloadFile
func (a *ApiRest) Query(c *gin.Context) {
	filename := c.Query("file")
	if filename == "" {
		web.Error(c, http.StatusBadRequest, "invalid request param")
		return
	}
	userNames, err := a.directoryService.Query(filename)
	if err != nil {
		handleErrors(c, err)
		return
	}
	user, err := a.authService.SelectRandomPeer(userNames)
	if err != nil {
		handleErrors(c, err)
		return

	}
	locationPath := user.UserURL
	web.SuccessQuery(c, http.StatusOK, locationPath)

}
func (a *ApiRest) AssignPeerUploading(c *gin.Context) {
	uploadingFile := c.Query("filename")
	if uploadingFile == "" {
		web.Error(c, http.StatusBadRequest, "invalid request param")
		return
	}
	username:= c.Query("user")
	location, err := a.authService.AssignPeer(username,uploadingFile)
	if err != nil {
		handleErrors(c, err)
		return
	}
	web.Success(c, http.StatusOK, location)

}
func handleErrors(c *gin.Context, err error) {
	switch {
	case errors.Is(err, directory.ErrNotFound):
		web.Error(c, http.StatusNotFound, err.Error())
	case errors.Is(err, auth.ErrNotFound):
		web.Error(c, http.StatusNotFound, err.Error())
	case errors.Is(err, directory.ErrInvalidFormat):
		web.Error(c, http.StatusBadRequest, err.Error())
	case errors.Is(err, auth.ErrExists):
		web.Error(c, http.StatusConflict, err.Error())
	case errors.Is(err, auth.ErrNoPeersAvailable):
		web.Error(c, http.StatusNotFound, err.Error())
	default:
		web.Error(c, http.StatusInternalServerError, err.Error())

	}

}
