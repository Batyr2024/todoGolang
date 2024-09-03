package middleware

import (
	"github.com/rs/cors"
)

var Cors = cors.New(cors.Options{
	AllowedOrigins: []string{"http://example.com", "https://example.com"},
	AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	AllowedHeaders: []string{"Authorization", "Content-Type"},
	AllowCredentials: true,
   })