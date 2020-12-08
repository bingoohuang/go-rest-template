package db

import (
	"log"
	"time"

	"github.com/bingoohuang/go-rest-template/internal/pkg/conf"
	"github.com/bingoohuang/go-rest-template/internal/pkg/models/tasks"
	"github.com/bingoohuang/go-rest-template/internal/pkg/models/users"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

// SetupDB opens a database and saves the reference to `Database` struct.
func SetupDB() {
	c := conf.GetConf()
	dialect, datasource := c.Db.GormOpen()
	db, err := gorm.Open(dialect, datasource)
	if err != nil {
		log.Fatalf("db err: %v", err)
	}

	// Change this to true if you want to see SQL queries
	db.LogMode(false)
	s := db.DB()
	s.SetMaxIdleConns(c.Db.MaxIdleConns)
	s.SetMaxOpenConns(c.Db.MaxOpenConns)
	s.SetConnMaxLifetime(time.Duration(c.Db.MaxLifetime) * time.Second)
	DB = db
	migration()
}

// Auto migrate project models
func migration() {
	DB.AutoMigrate(&users.User{})
	DB.AutoMigrate(&users.UserRole{})
	DB.AutoMigrate(&tasks.Task{})
}

func GetDB() *gorm.DB {
	return DB
}
