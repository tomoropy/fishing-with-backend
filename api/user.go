package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/tomoropy/fishing-with-backend/db/sqlc"
)

type createUserRequest struct {
	Name        string `json:"name" binding:"required,max=10"`
	ProfileText string `json:"profileText" binding:"required,max=100"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Name: req.Name,
		ProfileText: req.ProfileText,
	}

	user, err := server.db.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return 
	}
	
	ctx.JSON(http.StatusOK, user)
}
