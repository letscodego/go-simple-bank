package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	db "github.com/letscodego/go-simple-bank/db/sqlc"
	"github.com/letscodego/go-simple-bank/util"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"` //https://github.com/go-playground/validator
	Password string `json:"password" binding:"required,min=6" `
	FullName string `json:"full_name" binding:"required" `
	Email    string `json:"email" binding:"required,email" `
}

type createUserResponse struct {
	Username          string    `json:"username"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		FullName:       req.FullName,
		Email:          req.Email,
		HashedPassword: hashPassword,
	}

	_, err = server.store.CreateUser(ctx, arg)
	if err != nil {
		if msErr, ok := err.(*mysql.MySQLError); ok {
			if msErr.Number == 1062 {
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	fetchedUser, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := createUserResponse{
		Username:          fetchedUser.Username,
		FullName:          fetchedUser.FullName,
		Email:             fetchedUser.Email,
		PasswordChangedAt: fetchedUser.PasswordChangedAt,
		CreatedAt:         fetchedUser.CreatedAt,
	}
	ctx.JSON(http.StatusOK, rsp)
}
