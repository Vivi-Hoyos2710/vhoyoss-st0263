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
	AuthService      auth.Service
	directoryService directory.ServiceDirectory
}

func NewApiRest(auth auth.Service, dir directory.ServiceDirectory) *ApiRest {
	return &ApiRest{
		AuthService:      auth,
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
	peer, ok := c.Get("peer")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Peer information not found"})
		return
	}
	peerString := peer.(auth.Peer).Username
	indexInfo.Username = peerString
	err := a.directoryService.SendIndex(indexInfo)
	if err != nil {
		handleErrors(c, err)
		return
	}
	web.Success(c, http.StatusOK, "index sent successfully")

}

// Login Call to the auth service to get the token, for fucntion Login
func (a *ApiRest) Login(c *gin.Context) {
	var peer auth.Peer
	if err := c.ShouldBindJSON(&peer); err != nil {
		web.Error(c, http.StatusBadRequest, "invalid request body")
		return
	}
	loggedPeer, err := a.AuthService.Login(peer)
	if err != nil {
		handleErrors(c, err)
		return
	}

	web.SuccessLogin(c, http.StatusOK, loggedPeer)

}
func (a *ApiRest) Logout(c *gin.Context) {
	peer, ok := c.Get("peer")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Peer information not found"})
		return
	}
	authUser := peer.(auth.Peer)
	err := a.AuthService.Logout(authUser)
	if err != nil {
		handleErrors(c, err)
		return
	}
	web.Success(c, http.StatusOK, "logged out successfully")

}

// Query Call to the directory service to get the location of the file, for fucntion DownloadFile
func (a *ApiRest) Query(c *gin.Context) {
	filename := c.Query("file")
	if filename == "" {
		web.Error(c, http.StatusBadRequest, "invalid request param")
		return
	}
	peer, ok := c.Get("peer")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Peer information not found"})
		return
	}
	excludedPeerr := peer.(auth.Peer)
	peerUserNames, err := a.directoryService.Query(filename)
	if err != nil {
		handleErrors(c, err)
		return
	}
	locationPath, err := a.AuthService.AssignPeer(excludedPeerr.Username, peerUserNames)
	if err != nil {
		handleErrors(c, err)
		return

	}
	web.SuccessQuery(c, http.StatusOK, locationPath)

}
func (a *ApiRest) AssignPeerUploading(c *gin.Context) {

	peer, ok := c.Get("peer")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Peer information not found"})
		return
	}
	excludedPeerr := peer.(auth.Peer)
	location, err := a.AuthService.AssignPeer(excludedPeerr.Username, nil)
	if err != nil {
		handleErrors(c, err)
		return
	}
	web.SuccessQuery(c, http.StatusOK, location)

}
func handleErrors(c *gin.Context, err error) {
	switch {
	case errors.Is(err, directory.ErrNotFound):
		web.Error(c, http.StatusNotFound, err.Error())
	case errors.Is(err, auth.ErrNotFound):
		web.Error(c, http.StatusNotFound, err.Error())
	case errors.Is(err, directory.ErrInvalidFormat):
		web.Error(c, http.StatusBadRequest, err.Error())
	case errors.Is(err, auth.ErrNoPeersAvailable):
		web.Error(c, http.StatusNotFound, err.Error())
	case errors.Is(err, auth.ErrInvalidPassword):
		web.Error(c, http.StatusUnauthorized, err.Error())

	default:
		web.Error(c, http.StatusInternalServerError, err.Error())

	}

}
