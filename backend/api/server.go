package api

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	db "github.com/m3phist/gobank/backend/db/sqlc"
)

type Server struct {
	queries *db.Queries
	router  *gin.Engine
}

func NewServer(port int) {
	g := gin.Default()

	g.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to GoBank API!",
		})
	})

	err := g.Run(fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
