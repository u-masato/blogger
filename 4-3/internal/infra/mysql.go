package infra

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	// dsn := os.Getenv("MYSQL_DSN")
	dsn := fmt.Sprintf("bloger:bloger@tcp(127.0.0.1:%d)/bloger?parseTime=true", 33306)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}

	// 接続テスト
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping MySQL: %v", err)
	}

	log.Println("Connected to MySQL successfully")
	return db
}
