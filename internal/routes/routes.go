package routes

import (
	"github.com/NishantRaut777/banking-api/internal/auth"
	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/health", health)

			authGroup := v1.Group("/auth")
			{
				repo := auth.NewRepository()
				service := auth.NewService(repo)
				handler := auth.NewHandler(service)

				authGroup.POST("/signup", handler.Signup)
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
