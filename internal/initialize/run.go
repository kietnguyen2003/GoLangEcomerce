package initialize

import (
	"fmt"
	"src/global"
)

func Run() {
	// phải đúng thứ tự
	LoadConfig()
	m := global.Config.MySql
	l := global.Config.Logger
	fmt.Print("Loading configuration successfully\n")
	fmt.Printf("MySql port:: %d\n", m.Port)
	fmt.Printf("Logger file:: %s\n", l.File)
	Initlogger()
	global.Logger.Info("Logger init success")
	InitMysql()
	// InitRedis()
	r := InitRouter()
	r.Run(":8002")
}
