package initialize

import (
	"fmt"
	"src/global"
	"src/internal/po"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql() {
	// Lấy thông tin cấu hình kết nối MySQL từ biến global.Config.MySql
	m := global.Config.MySql
	// kêt nối mysql thông qua gorm
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	s := fmt.Sprintf(dsn, m.User, m.Password, m.Host, m.Port, m.DbName)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		// này là cái gì á mà nó đc set măc định là true
		SkipDefaultTransaction: false,
	})
	if err != nil {
		global.Logger.Error("InitMysql error: ", zap.Error(err))
		panic(err)
	}
	global.Mdb = db
	global.Logger.Info("InitMysql success")

	// setPool
	setPool()
	// migrateDb
	migrateDb()
}

func setPool() {
	m := global.Config.MySql
	sqlDb, err := global.Mdb.DB()

	if err != nil {
		global.Logger.Error("setPool error: ", zap.Error(err))
		panic(err)
	}
	// set config cho mysql
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
}

func migrateDb() {
	// tạo folder pool để chứa các file migration
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)
	if err != nil {
		global.Logger.Error("migrateDb error: ", zap.Error(err))
		panic(err)
	}
}
