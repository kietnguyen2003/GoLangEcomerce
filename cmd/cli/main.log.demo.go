package main

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// // 1. Example
	// logger1 := zap.NewExample()
	// logger1.Info("Example logger1")

	// // 2. Development
	// logger2, _ := zap.NewDevelopment()
	// logger2.Info("Development logger2")

	// // 3. Production
	// logger3, _ := zap.NewProduction()
	// logger3.Info("Production logger3")

	// // 4. Custom log
	// core := zapcore.NewCore(getEncoderLog(), zapcore.AddSync(os.Stdout), zapcore.InfoLevel)
	// logger4 := zap.New(core)
	// logger4.Info("Custom logger4")

	// 5. Ghi log with WriteSyncer
	encoder := getEncoderLog()
	sync := getWriteSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger5 := zap.New(core)

	logger5.Info("Info logger1", zap.Int("line", 1))
	logger5.Error("Error logger2", zap.Int("line", 2))

	// Kết quả
	// {"level":"info","msg":"Example logger1"}
	// 2025-01-10T16:48:28.350+0700    INFO    cli/main.log.demo.go:14 Development logger2
	// {"level":"info","ts":1736502508.3511257,"caller":"cli/main.log.demo.go:18","msg":"Production logger3"}
}

// custom log
func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionConfig().EncoderConfig

	// 1736502508.3511257 -> 2025-01-10T16:48:28.350+0700
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// ts -> Time
	encodeConfig.TimeKey = "Time"

	// level -> Level
	encodeConfig.LevelKey = "Level"

	// thêm caller vào log
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encodeConfig)
}

func getWriteSync() zapcore.WriteSyncer {
	// Kiểm tra và tạo thư mục log nếu chưa tồn tại
	logDir := "./cmd/cli/log/"
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		panic(fmt.Sprintf("Failed to create log directory: %v", err))
	}

	// Mở file log
	file, err := os.OpenFile(logDir+"log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("Failed to open log file: %v", err))
	}

	// Kết hợp ghi log vào file và console
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
