// Package middleware proporciona funciones de middleware para configurar y gestionar el CORS (Cross-Origin Resource Sharing).
package middleware

import (
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// appCORS habilita la configuración de CORS para permitir solicitudes desde diferentes orígenes.
// Utiliza la lista de dominios aceptados definida en la variable de entorno ACCEPTED_DOMAINS.
// Habilita los métodos GET, POST, PUT, HEAD y OPTIONS.
// Los encabezados permitidos incluyen X-Requested-With, Content-Type, Authorization y X-Wacohunt-Api-Key.
// Permite el uso de credenciales en las solicitudes.
func serverCORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: strings.Split(os.Getenv("ACCEPTED_DOMAINS"), ","),
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders: []string{
			"Content-Type",
			"Authorization",
		},
		AllowCredentials: false, // Permitir el uso de credenciales en las solicitudes
	})
}
