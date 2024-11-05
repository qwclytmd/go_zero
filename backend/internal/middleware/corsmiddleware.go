package middleware

import (
	"github.com/gin-contrib/cors"
	"net/http"
	"time"
)

type CorsMiddleware struct {
}

func NewCorsMiddleware() *CorsMiddleware {
	return &CorsMiddleware{}
}

func (m *CorsMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Content-Type", "Origin", "Accept", "*"},
			MaxAge:           12 * time.Hour,
			AllowCredentials: true,
		})

		// 放行所有OPTIONS方法
		if r.Method == "OPTIONS" {
			return
		}
		next(w, r)
	}
}
