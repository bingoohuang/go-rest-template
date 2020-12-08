package controllers

import (
	"github.com/bingoohuang/go-rest-template/internal/pkg/ginx"
	"github.com/bingoohuang/go-rest-template/internal/pkg/persist"
	"github.com/bingoohuang/go-rest-template/pkg/crypto"
	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context, bindJSON interface{}) ginx.Render {
	loginInput := bindJSON.(LoginInput)
	s := persist.GetUserRepo()
	user, err := s.GetByUsername(loginInput.Username)
	if err != nil {
		return ginx.New404Error("user not found", err)
	}

	if !crypto.ComparePasswords(user.Hash, []byte(loginInput.Password)) {
		return ginx.New403Error("user and password not match", nil)
	}

	token, _ := crypto.CreateToken(user.Username)
	return ginx.JSON(token)
}
