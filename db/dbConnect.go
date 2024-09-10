package db

import (
	"database/sql"
	"fmt"

	"github.com/dinosaur1258/GolangBackEnd/model" // 替換成你的模塊路徑
	_ "github.com/lib/pq"
)

func DatabaseConnection() (*sql.DB, error) {
	// 創建 YAMLConfigLoader 實例
	var loader model.ConfigLoader = &model.YAMLConfigLoader{}

	// 讀取並解析 YAML 檔案
	config, err := loader.LoadConfig("./conf/conf.yaml")
	if err != nil {
		return nil, fmt.Errorf("error loading config: %w", err)
	}
	// 打印配置確認讀取成功
	fmt.Printf("Loaded config: %+v\n", config.Database)

	// 使用解析的資料構建連線字符串
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password, config.Database.Dbname)

	// 打開資料庫連線
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	// 確保連線能夠連接到資料庫
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	fmt.Println("Database connection established successfully")
	return db, nil
}
