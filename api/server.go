package api

import (
	"github.com/gin-gonic/gin"
	_ "github.com/tomoropy/fishing-with-backend/db/sqlc"
	db "github.com/tomoropy/fishing-with-backend/db/sqlc"
)

type Server struct {
	db db.Store
	router *gin.Engine
}

func NewServer(db db.Store) *Server {
	server := &Server{db: db}
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.GET("/users/:id", server.getUser)
	router.GET("/users", server.listUser)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
