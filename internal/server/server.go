package server

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/NishantRaut777/banking-system-go/internal/config"
	"github.com/NishantRaut777/banking-system-go/internal/database"
	"github.com/NishantRaut777/banking-system-go/internal/routes"
	"github.com/NishantRaut777/banking-system-go/internal/utils"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

func Start() {
	cfg := config.LoadConfig()

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	if cfg.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	// connect to postgresql before starting server
	database.Connect(cfg.DatabaseURL)
	defer database.Close()

	utils.SetJWTSecret(cfg.JWTSecret)

	router := gin.New()
	router.MaxMultipartMemory = 10 << 20

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	routes.Register(router, cfg)

	router.GET("/swagger/*any",
		ginSwagger.WrapHandler(
			swaggerFiles.Handler,
			ginSwagger.URL("doc.json"),
			ginSwagger.DefaultModelsExpandDepth(-1),
		),
	)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "route not found",
		})
	})

	// HTTP Server
	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	// Start server
	go func() {
		logger.Info("Starting server", zap.String("port", cfg.Port))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatal("Server error", zap.Error(err))
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Forced shutdown", zap.Error(err))
	}

	logger.Info("Server stopped")
}
