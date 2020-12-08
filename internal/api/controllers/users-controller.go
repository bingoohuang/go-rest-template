package controllers

import (
	"net/http"

	"github.com/bingoohuang/go-rest-template/internal/pkg/ginx"
	models "github.com/bingoohuang/go-rest-template/internal/pkg/models/users"
	"github.com/bingoohuang/go-rest-template/internal/pkg/persist"
	"github.com/bingoohuang/go-rest-template/pkg/crypto"
	"github.com/gin-gonic/gin"
)

type UserInput struct {
	Username  string `json:"username" binding:"required"`
	Lastname  string `json:"lastname"`
	Firstname string `json:"firstname"`
	Password  string `json:"password" binding:"required"`
	Role      string `json:"role"`
}

// GetUserById godoc
// @Summary Retrieves user based on given ID
// @Description get User by ID
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} users.User
// @Router /api/users/{id} [get]
// @Security Authorization Token
func GetUserById(c *gin.Context, id string) ginx.Render {
	user, err := persist.GetUserRepo().Get(id)
	if err != nil {
		return ginx.New404Error("user not found", err)
	}

	return ginx.JSON(user)
}

// GetUsers godoc
// @Summary Retrieves users based on query
// @Description Get Users
// @Produce json
// @Param username query string false "Username"
// @Param firstname query string false "Firstname"
// @Param lastname query string false "Lastname"
// @Success 200 {array} []users.User
// @Router /api/users [get]
// @Security Authorization Token
func GetUsers(c *gin.Context, bind interface{}) ginx.Render {
	q := bind.(models.User)
	users, err := persist.GetUserRepo().Query(&q)
	if err != nil {
		return ginx.New404Error("users not found", err)
	}

	return ginx.JSON(users)
}

func CreateUser(c *gin.Context, bindJSON interface{}) ginx.Render {
	userInput := bindJSON.(UserInput)
	user := models.User{
		Username:  userInput.Username,
		Firstname: userInput.Firstname,
		Lastname:  userInput.Lastname,
		Hash:      crypto.HashAndSalt([]byte(userInput.Password)),
		Role:      models.UserRole{RoleName: userInput.Role},
	}
	if err := persist.GetUserRepo().Add(&user); err != nil {
		return ginx.New400Error("", err)
	}

	return ginx.StatusJSON(http.StatusCreated, user)
}

func UpdateUser(c *gin.Context, id string, bindJSON interface{}) ginx.Render {
	s := persist.GetUserRepo()
	user, err := s.Get(id)
	if err != nil {
		return ginx.New404Error("user not found", err)
	}

	userInput := bindJSON.(UserInput)
	user.Username = userInput.Username
	user.Lastname = userInput.Lastname
	user.Firstname = userInput.Firstname
	user.Hash = crypto.HashAndSalt([]byte(userInput.Password))
	user.Role = models.UserRole{RoleName: userInput.Role}
	if err := s.Update(user); err != nil {
		return ginx.New404Error("", err)
	}

	return ginx.JSON(user)
}

func DeleteUser(c *gin.Context, id string) ginx.Render {
	s := persist.GetUserRepo()
	user, err := s.Get(id)
	if err != nil {
		return ginx.New404Error("user not found", err)
	}

	if err := s.Delete(user); err != nil {
		return ginx.New404Error("", err)
	}

	return ginx.StatusJSON(http.StatusNoContent, "")
}
