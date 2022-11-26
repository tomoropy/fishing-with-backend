package api

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
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

	// ここからCorsの設定
  router.Use(cors.New(cors.Config{
    // アクセスを許可したいアクセス元
    AllowOrigins: []string{
      "http://localhost:3000",
    },
    // アクセスを許可したいHTTPメソッド
    AllowMethods: []string{
        "POST",
				"PUT",
        "GET",
				"DELETE",
        "OPTIONS",
    },
    // 許可したいHTTPリクエストヘッダ
    AllowHeaders: []string{
        "Access-Control-Allow-Credentials",
        "Access-Control-Allow-Headers",
        "Content-Type",
        "Content-Length",
        "Accept-Encoding",
        "Authorization",
    },
    // cookieなどの情報を必要とするかどうか
    AllowCredentials: true,
    // preflightリクエストの結果をキャッシュする時間
    MaxAge: 24 * time.Hour,
  }))

	// 認証不要
	router.POST("/users", server.createUser)      //signup
	router.POST("/users/login", server.loginUser) //login
	router.GET("/users", server.listUser)         //users
	router.GET("/users/:id", server.getUser)      //user show page

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	// 要 認証
	// Users
	authRoutes.GET("/users/me", server.myInfo)     //auth user
	authRoutes.PUT("/users", server.updateUser)    //update user
	authRoutes.DELETE("/users", server.deleteUser) //delete user

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
