package initialize

import (
	"src/global"
	"src/package/logger"
)

func Initlogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
