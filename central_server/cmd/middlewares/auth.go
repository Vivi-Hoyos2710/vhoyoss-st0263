package middlewares

import (
	"fmt"
	"net/http"

	"github.com/Vivi-Hoyos2710/vhoyoss-st0263/central_server/internal/auth"
	"github.com/Vivi-Hoyos2710/vhoyoss-st0263/central_server/pkg/web"
	"github.com/gin-gonic/gin"
)

func CustomMiddleware(s *auth.ServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		Token := c.Request.Header.Get("Authorization")
		claim, err := s.ValidatePeerToken(Token)
		if err != nil {
			fmt.Println(err)
			web.Error(c, http.StatusUnauthorized, "Unauthorized")
			c.Abort()
		} else {
			c.Set("peer", claim.Peer)
			c.Next()
		}

	}
}
