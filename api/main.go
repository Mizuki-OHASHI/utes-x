package main

import (
	"database/sql"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// DB_HOST=localhost
// DB_PORT=3307
// DB_USER=utes_x_user
// DB_PASSWORD=utes_x_password
// DB_NAME=utes_x_database
// API_PORT=8888
type EnvironmentalVariables struct {
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	API_PORT    string
}

var envs EnvironmentalVariables = EnvironmentalVariables{}

// 環境変数を読み込む
func load_envs() {
	godotenv.Load(".env")
	envs.DB_HOST = os.Getenv("DB_HOST")
	envs.DB_PORT = os.Getenv("DB_PORT")
	envs.DB_USER = os.Getenv("DB_USER")
	envs.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	envs.DB_NAME = os.Getenv("DB_NAME")
	envs.API_PORT = os.Getenv("API_PORT")
	if envs.DB_HOST == "" || envs.DB_PORT == "" || envs.DB_USER == "" || envs.DB_PASSWORD == "" || envs.DB_NAME == "" || envs.API_PORT == "" {
		panic("Environment variables not set correctly. Please check your .env file.")
	}
}

// データベースに接続する
func connectDB() *sql.DB {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic("Failed to load Asia/Tokyo timezone: " + err.Error())
	}
	cfg := mysql.Config{
		User:                 envs.DB_USER,
		Passwd:               envs.DB_PASSWORD,
		Addr:                 envs.DB_HOST + ":" + envs.DB_PORT,
		DBName:               envs.DB_NAME,
		Net:                  "tcp",
		ParseTime:            true,
		Collation:            "utf8mb4_unicode_ci",
		Loc:                  jst,
		AllowNativePasswords: true,
	}
	dsn := cfg.FormatDSN()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	if err := db.Ping(); err != nil {
		panic("Failed to ping the database: " + err.Error())
	}
	return db
}

func registerHandlers(engine *gin.Engine, db *sql.DB) {
	// TODO
}

func init() {
	load_envs()
}

func main() {
	db := connectDB()
	defer db.Close()

	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool { return true },
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Length",
			"Content-Type",
			"Authorization",
		},
	}))

	registerHandlers(engine, db)
	engine.Run(":" + envs.API_PORT)
}
