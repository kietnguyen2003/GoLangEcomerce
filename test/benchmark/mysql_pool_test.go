package benchmark

import (
	"fmt"
	"log"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id   uint
	Name string
}

func insertRecord(b *testing.B, db *gorm.DB) {
	user := User{Name: "test"}
	if err := db.Create(&user).Error; err != nil {
		b.Fatal(err)
	}
}

// go test -bench=. -benchmem

// func BenchmarkMaxOpenConss1(b *testing.B) {
// 	// kêt nối mysql thông qua gorm
// 	dsn := "root:123456@tcp(127.0.0.1:3306)/shopdevgo?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
// 		// này là cái gì á mà nó đc set măc định là true
// 		// SkipDefaultTransaction: false,
// 	})

// 	if err != nil {
// 		log.Fatalf("Failed to connect to mysql: %v", err)
// 	}

// 	// Drop bảng user nếu tồn tại
// 	if db.Migrator().HasTable(&User{}) {
// 		if err := db.Migrator().DropTable(&User{}); err != nil {
// 			fmt.Print(err)
// 		}
// 	}

// 	// Tạo bảng user
// 	db.AutoMigrate(&User{})
// 	sql, err := db.DB()
// 	if err != nil {
// 		log.Fatalf("Failed to get sql.DB: %v", err)
// 	}
// 	// set config cho mysql
// 	sql.SetMaxOpenConns(1)
// 	defer sql.Close()

// 	//
// 	b.RunParallel(func(pb *testing.PB) {
// 		for pb.Next() {
// 			insertRecord(b, db)
// 		}
// 	})
// }

func BenchmarkMaxOpenConss10(b *testing.B) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/shopdevgo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// này là cái gì á mà nó đc set măc định là true
		// SkipDefaultTransaction: false,
	})

	if err != nil {
		log.Fatalf("Failed to connect to mysql: %v", err)
	}

	// Drop bảng user nếu tồn tại
	if db.Migrator().HasTable(&User{}) {
		if err := db.Migrator().DropTable(&User{}); err != nil {
			fmt.Print(err)
		}
	}

	// Tạo bảng user
	db.AutoMigrate(&User{})
	sql, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get sql.DB: %v", err)
	}
	// set config cho mysql
	sql.SetMaxOpenConns(10)
	defer sql.Close()

	//
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecord(b, db)
		}
	})
}
