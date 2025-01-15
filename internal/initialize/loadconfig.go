// file này để đọc các config trong file Viper (các file yaml đã làm)
package initialize

import (
	"fmt"
	"src/global"

	"github.com/spf13/viper"
)

func LoadConfig() {
	// Configurations for Viper
	viper := viper.New()             // tạo viper instance
	viper.AddConfigPath("./config/") // nơi chứa file config
	viper.SetConfigName("local")     // tên file config
	viper.SetConfigType("yaml")      // loại file config

	// Đọc file cấu hình
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// read config via structure
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Enable to decode configuration structure: %v", err)
	}
}
