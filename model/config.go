package model

// dbConfig 定義了數據庫配置的結構
type dbConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

// Config 定義了配置的結構
type Config struct {
	Database dbConfig `yaml:"database"`
}
