package api

import (
    "github.com/gin-gonic/gin"
    "github.com/yourorg/saascache/internal/api/handlers"
)

// NewRouter builds the HTTP router with middleware and routes.
func NewRouter() *gin.Engine {
    r := gin.New()
    r.Use(gin.Logger(), gin.Recovery())

    v1 := r.Group("/v1")
    {
        v1.GET("/health", handlers.Health)
        v1.POST("/recommend", handlers.Recommend)
        v1.POST("/provision", handlers.Provision)
    }
    return r
}
