package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/tomoropy/fishing-with-backend/db/sqlc"
	"github.com/tomoropy/fishing-with-backend/util"
)

type createUserRequest struct {
	Name        string `json:"name" binding:"required,alphanum,max=10"`
	ProfileText string `json:"profileText" binding:"required,max=100"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=6"`
}

type createUserResponse struct {
	ID           int32          `json:"id"`
	Name         string         `json:"name"`
	ProfileText  string         `json:"profileText"`
	ProfileImage sql.NullString `json:"profileImage"`
	HeaderImage  sql.NullString `json:"headerImage"`
	CreatedAt    sql.NullTime   `json:"createdAt"`
	Email        string         `json:"email"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Name:           req.Name,
		ProfileText:    req.ProfileText,
		Email:          req.Email,
		HashedPassword: hashedPassword,
	}

	user, err := server.db.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := createUserResponse{
		Name:         user.Name,
		ProfileText:  user.ProfileText,
		Email:        user.Email,
		ProfileImage: user.ProfileImage,
		HeaderImage:  user.HeaderImage,
		CreatedAt:    user.CreatedAt,
	}

	ctx.JSON(http.StatusOK, rsp)
}

type getUserRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.db.GetUser(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type listUserRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listUser(ctx *gin.Context) {
	var req listUserRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListUserParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	users, err := server.db.ListUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		log.Fatalln(err)
		return
	}

	ctx.JSON(http.StatusOK, users)
}
