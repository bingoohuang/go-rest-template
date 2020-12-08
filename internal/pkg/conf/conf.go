package conf

import (
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
	Port         string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
}

func (d DBConf) GormOpen() (dialect, datasource string) {
	switch d.Driver {
	case "sqlite":
		return "sqlite3", "./" + d.Dbname + ".db"
	case "postgres":
		return "postgres", "host=" + d.Host + " port=" + d.Port + " user=" + d.Username + " dbname=" + d.Dbname + "  sslmode=disable password=" + d.Password
	case "mysql":
		return "mysql", d.Username + ":" + d.Password + "@tcp(" + d.Host + ":" + d.Port + ")/" + d.Dbname + "?charset=utf8&parseTime=True&loc=Local"
	}
	return "", ""
}

type ServerConf struct {
	Port   string
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
