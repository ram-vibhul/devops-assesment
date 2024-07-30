package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) healthcheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "healthy",
	})
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
