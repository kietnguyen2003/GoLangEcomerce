package initialize

import (
	"fmt"
	"src/global"
)

func Run() {
	// phải đúng thứ tự
	LoadConfig()
	m := global.Config.MySql
	fmt.Print("Loading configuration successfully\n")
	fmt.Printf("MySql port:: %d\n", m.Port)
	Initlogger()
	InitMysql()
	InitRedis()
	r := InitRouter()
	r.Run(":8002")
}
