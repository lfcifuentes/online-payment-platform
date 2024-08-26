package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lfcifuentes/online-payment-platform/api/app/internal/adapters/pgsql"
	"github.com/lfcifuentes/online-payment-platform/api/app/modules/auth/data/repositories"
	"github.com/lfcifuentes/online-payment-platform/api/app/pkg"
)

// JWTMiddleware Middleware para validar el token
func JWTMiddleware(db *pgsql.DBAdapter) gin.HandlerFunc {
	authRepository := repositories.NewAuthRepository(db)
	return func(c *gin.Context) {
		// check header token
		tokenHeader := c.GetHeader("Authorization")
		// Verificar si el token de API está presente
		if tokenHeader == "" {
			// Si el token de API está ausente, responder con un código de estado 422 y un mensaje de error "Unauthorized"
			c.AbortWithStatusJSON(
				http.StatusUnprocessableEntity,
				gin.H{"message": "Unauthorized"},
			)
			return
		}
		jwtApi := pkg.NewApiJWT()
		claims, token, err := jwtApi.ValidateJWT(tokenHeader)

		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnprocessableEntity,
				gin.H{"message": "Unauthorized"},
			)
			return
		}
		// Verificar si el token es válido
		tokenStatus, err := authRepository.ValidateToken(token)
		if err != nil || !tokenStatus {
			c.AbortWithStatusJSON(
				http.StatusUnprocessableEntity,
				gin.H{"message": "Unauthorized"},
			)
			return
		}
		// id del usuario
		userID := claims.User.Id
		c.Set("userID", int64(userID))

		// Setear el token y los claims en el contexto
		c.Set("token", token)
		c.Set("claims", claims)

		// Add your code here
		c.Next()
	}
}
