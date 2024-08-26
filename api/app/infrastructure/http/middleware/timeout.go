// Package middleware proporciona funciones de middleware para agregar un timeout a las solicitudes en un servidor Gin.
package middleware

import (
	"net/http"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

// timeoutMiddleware es un middleware de Gin que establece un timeout para las solicitudes entrantes.
// Por defecto, el timeout se establece en 40 minutos.
// Si una solicitud excede el timeout especificado, se enviará una respuesta con un código de estado 408 (Request Timeout).
func timeoutMiddleware() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(5*time.Second), // Establecer el timeout en 40 minutos
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(func(c *gin.Context) {
			c.String(http.StatusRequestTimeout, "timeout")
		}),
	)
}
