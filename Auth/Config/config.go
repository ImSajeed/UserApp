package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var DB struct {
	Username string
	Password string
	Endpoint string
	DBName   string
}

func Load() {
	fmt.Println("..Loading Config Data..")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()

	if err != nil {
		//panic(fmt.Errorf("Fatal error config file: %s", err.Error()))
		fmt.Println(err.Error())
		panic("Failed loading config file")
	}

	populateConfigurableVariable()
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		populateConfigurableVariable()
	})

	fmt.Println(" Config loading Done ")

}

func populateConfigurableVariable() {
	DB.Username = viper.GetString("DB.username")
	DB.Password = viper.GetString("DB.password")
	DB.Endpoint = viper.GetString("DB.endpoint")
	DB.DBName = viper.GetString("DB.dbname")

}
