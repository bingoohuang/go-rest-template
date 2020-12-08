package test

import (
	"fmt"
	"testing"

	"github.com/bingoohuang/go-rest-template/internal/pkg/conf"
	"github.com/bingoohuang/go-rest-template/internal/pkg/db"
	models "github.com/bingoohuang/go-rest-template/internal/pkg/models/users"
	"github.com/bingoohuang/go-rest-template/internal/pkg/persist"
)

var userTest models.User

func Setup() {
	conf.Setup("config.yml")
	db.SetupDB()
	db.GetDB().Exec("DELETE FROM users")
}

func TestAddUser(t *testing.T) {
	Setup()
	user := models.User{
		Firstname: "Antonio",
		Lastname:  "Paya",
		Username:  "antonio",
		Hash:      "hash",
		Role:      models.UserRole{RoleName: "user"},
	}
	s := persist.GetUserRepo()
	if err := s.Add(&user); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	userTest = user
}

func TestGetAllUsers(t *testing.T) {
	s := persist.GetUserRepo()
	if _, err := s.All(); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestGetUserById(t *testing.T) {
	db.SetupDB()
	s := persist.GetUserRepo()
	if _, err := s.Get(fmt.Sprint(userTest.ID)); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
