package persist

import (
	"strconv"

	"github.com/bingoohuang/go-rest-template/internal/pkg/db"
	models "github.com/bingoohuang/go-rest-template/internal/pkg/models/users"
)

type UserRepo struct{}

var userRepo *UserRepo

func GetUserRepo() *UserRepo {
	if userRepo == nil {
		userRepo = &UserRepo{}
	}
	return userRepo
}

func (r *UserRepo) Get(id string) (*models.User, error) {
	var user models.User
	where := models.User{}
	where.ID, _ = strconv.ParseUint(id, 10, 64)
	_, err := First(&where, &user, []string{"Role"})
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *UserRepo) GetByUsername(username string) (*models.User, error) {
	var user models.User
	where := models.User{Username: username}
	_, err := First(&where, &user, []string{"Role"})
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *UserRepo) All() (*[]models.User, error) {
	var users []models.User
	err := Find(&models.User{}, &users, []string{"Role"}, "id asc")
	return &users, err
}

func (r *UserRepo) Query(q *models.User) (*[]models.User, error) {
	var users []models.User
	err := Find(&q, &users, []string{"Role"}, "id asc")
	return &users, err
}

func (r *UserRepo) Add(user *models.User) error {
	if err := Create(&user); err != nil {
		return err
	}

	return Save(&user)
}

func (r *UserRepo) Update(user *models.User) error {
	var userRole models.UserRole
	if _, err := First(models.UserRole{UserID: user.ID}, &userRole, []string{}); err != nil {
		return err
	}
	userRole.RoleName = user.Role.RoleName
	if err := Save(&userRole); err != nil {
		return err
	}

	if err := db.GetDB().Omit("Role").Save(&user).Error; err != nil {
		return err
	}

	user.Role = userRole
	return nil
}

func (r *UserRepo) Delete(user *models.User) error {
	d := db.GetDB().Unscoped()
	if err := d.Delete(models.UserRole{UserID: user.ID}).Error; err != nil {
		return err
	}

	return d.Delete(&user).Error
}
