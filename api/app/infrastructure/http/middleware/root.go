package middleware

import "github.com/gin-gonic/gin"

// StartAllMiddlewares configura y aplica una serie de middlewares comunes en un motor Gin.
// Los middlewares aplicados incluyen el registro de solicitudes, la recuperación de panics,
// el timeout de solicitudes, la configuración CORS, la obtención de la IP real del cliente y
// la restricción de tipos de contenido.
func StartAllMiddlewares(r *gin.Engine) *gin.Engine {
	// Usar los middleware básicos
	r.Use(
		gin.Logger(),        // Middleware para registrar las solicitudes HTTP
		gin.Recovery(),      // Middleware para recuperarse de panics y devolver un error 500
		timeoutMiddleware(), // Middleware para establecer un timeout en las solicitudes
		serverCORS(),        // Middleware para habilitar la configuración CORS
	)

	return r
}
