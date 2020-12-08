package conf

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var Config *Conf

type Conf struct {
	Server ServerConf
	Db     DBConf
}

type DBConf struct {
	Driver       string
	Dbname       string
	Username     string
	Password     string
	Host         string
	Port         int
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
}

func (d DBConf) GormOpen() (dialect, datasource string) {
	switch d.Driver {
	case "sqlite":
		return "sqlite3", d.Dbname
	case "postgres":
		return "postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s  sslmode=disable password=%s",
			d.Host, d.Port, d.Username, d.Dbname, d.Password)
	case "mysql":
		return "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d))/%s?charset=utf8&parseTime=True&loc=Local",
			d.Username, d.Password, d.Host, d.Port, d.Dbname)
	}
	return "", ""
}

type ServerConf struct {
	Port   int
	Secret string
	Mode   string
}

// SetupDB initialize configuration
func Setup(confPath string) {
	var conf *Conf

	viper.SetConfigFile(confPath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading conf file, %s", err)
	}

	err := viper.Unmarshal(&conf)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	Config = conf
}

// GetConf helps you to get configuration data
func GetConf() *Conf {
	return Config
}
