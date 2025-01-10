package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`

	Security struct {
		Jwt struct {
			Secret     string `mapstructure:"secret"`
			Expiration int    `mapstructure:"expiration"`
		} `mapstructure:"jwt"`
	} `mapstructure:"security"`

	Database []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
		DbName   string `mapstructure:"dbName"`
		Port     int    `mapstructure:"port"`
	} `mapstructure:"databases"`
}

func main() {

	// Configurations for Viper
	viper := viper.New()             // tạo viper instance
	viper.AddConfigPath("./config/") // nơi chứa file config
	viper.SetConfigName("local")     // tên file config
	viper.SetConfigType("yaml")      // loại file config

	err := viper.ReadInConfig()
	if err != nil {
		// panic là dừng luôn chương trình
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// read config
	serverPort := viper.GetInt("server.port")
	fmt.Println("Server port::", serverPort)

	jwtCodeSecret := viper.GetString("security.jwt.secret")
	fmt.Println("JWT secret::", jwtCodeSecret)
	fmt.Println("=====================================")

	// read config via structure
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Enable to decode configuration structure: %v", err)
	}
	fmt.Println("Server port::", config.Server.Port)
	fmt.Println("=====================================")

	for _, db := range config.Database {
		fmt.Println("User::", db.User)
		fmt.Println("Password::", db.Password)
		fmt.Println("Host::", db.Host)
		fmt.Println("DbName::", db.DbName)
		fmt.Println("Port::", db.Port)
		fmt.Println("=====================================")
	}
}
