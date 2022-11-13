package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/tomoropy/fishing-with-backend/db/sqlc"
	db "github.com/tomoropy/fishing-with-backend/db/sqlc"
	"github.com/tomoropy/fishing-with-backend/token"
	"github.com/tomoropy/fishing-with-backend/util"
)

type Server struct {
	config     util.Config
	db         db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, db db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannnot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		db:         db,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server)setupRouter(){
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	router.GET("/users/:id", server.getUser)
	router.GET("/users", server.listUser)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
