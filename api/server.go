package api

import (
	"fmt"
	"strconv"

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

func (server *Server) setupRouter() {
	router := gin.Default()

	// 認証不要
	router.POST("/users", server.createUser)      //signup
	router.POST("/users/login", server.loginUser) //login

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	// Users
	authRoutes.GET("/users", server.listUser)          //users
	authRoutes.GET("/users/:id", server.getUser)       //user show page
	authRoutes.PUT("/users/:id", server.updateUser)    //update user
	authRoutes.DELETE("/users/:id", server.deleteUser) //delete user

	server.router = router
}

func validID(id string) (int, error) {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return 0, fmt.Errorf("id could not convert to integer")
	}
	if intID < 1 {
		return 0, fmt.Errorf("id cannot be under 0")
	}
	return intID, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
