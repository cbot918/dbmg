package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var lg = fmt.Println

type Config struct {
	DB_Type string
	DB_URL  string
}

func NewConfig() *Config {
	return &Config{
		DB_Type: os.Getenv("DB_TYPE"),
		DB_URL:  os.Getenv("DB_URL"),
	}
}

func main() {

	// 處理讀入的第一個參數
	read := os.Args
	if len(read) != 2 {
		fmt.Println("請指定一個檔案")
		return
	}
	file := read[1]

	// 讀入sql內容
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

	// 讀入連線db的設定檔資訊
	godotenv.Load()
	cfg := NewConfig()

	// new一個dby instance
	db, err := NewDBy(cfg.DB_Type, cfg.DB_URL)
	if err != nil {
		log.Fatal(err)
	}

	// 執行sql
	result, err := db.DB.Exec(string(data))
	if err != nil {
		log.Fatal(err)
		fmt.Printf("%+v", err)
	}
	fmt.Println(result)
}
