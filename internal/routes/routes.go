package routes

import (
	"github.com/NishantRaut777/banking-api/internal/auth"
	"github.com/NishantRaut777/banking-api/internal/config"
	"github.com/NishantRaut777/banking-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine, cfg *config.Config) {
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/health", health)

			repo := auth.NewRepository()
			service := auth.NewService(repo)
			handler := auth.NewHandler(service)

			authGroup := v1.Group("/auth")
			{
				authGroup.POST("/signup", handler.Signup)
				authGroup.POST("/login", handler.Login)
			}

			userGroup := v1.Group("/users")
			userGroup.Use(middleware.AuthMiddleware([]byte(cfg.JWTSecret)))
			{
				userGroup.GET("/me", handler.Me)
			}
		}
	}
}

func health(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "ok",
		"service": "banking-api",
	})
}
