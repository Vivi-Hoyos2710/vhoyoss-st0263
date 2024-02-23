package main

import (
	"github.com/Vivi-Hoyos2710/vhoyoss-st0263/central_server/cmd/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	eng := gin.Default()
	router := routes.NewRouter(eng)
	router.MapRoutes()
	if err := eng.Run(); err != nil {
		panic(err)
	}
}
