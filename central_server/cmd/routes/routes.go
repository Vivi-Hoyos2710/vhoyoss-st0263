package routes

import (
	"github.com/Vivi-Hoyos2710/vhoyoss-st0263/central_server/cmd/handlers"
	"github.com/Vivi-Hoyos2710/vhoyoss-st0263/central_server/internal/auth"
	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	engine      *gin.Engine
	routerGroup *gin.RouterGroup
}

func NewRouter(eng *gin.Engine) Router {
	return &router{engine: eng}
}

func (r *router) MapRoutes() {
	r.buildHealthCheckRoutes()
	r.setGroup()
	r.buildAuthRoutes()

}

func (r *router) setGroup() {
	r.routerGroup = r.engine.Group("/api/v1")
}

func (r *router) buildHealthCheckRoutes() {
	r.engine.GET("/ping", handlers.HealthCheckHandler)
}
func (r *router) buildAuthRoutes() {
	peersMap := make(map[int]auth.Peer)
	repo := auth.NewDefaultRepo(peersMap)
	authService := auth.NewServiceClient(repo)
	handler := handlers.NewApiRest(authService)
	r.routerGroup.POST("/login", handler.Login)
}
