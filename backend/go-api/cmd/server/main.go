package main

import (
	"fmt"
	"log"

	"go-api/internal/repository"

	"gorm.io/driver/postgres" // 💡 追加: postgresドライバ
	"gorm.io/gorm"
)

func main() {
	// Docker Composeサービス名 (db) をホストとして使用
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		"db", // 💡 変更: ホスト名はサービス名 'db'
		"user",
		"password",
		"app_db",
		"5432",
	)

	// データベース接続の初期化
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) // 💡 変更: postgres.Open() を使用
	if err != nil {
		log.Fatalf("failed to connect database. Error: %v", err)
	}

	// ... (以降のコードは変更なし)
	db.AutoMigrate(&repository.User{})

	// ... (Handler, Service, Routerの設定)

	log.Println("Server starting on :8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("server run failed: %v", err)
	}
}
