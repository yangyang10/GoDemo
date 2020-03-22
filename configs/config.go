package configs

import (
	"os"
	"strconv"
	"time"
)

var (

	//db
	MYSQL_USER     = "timel"
	MYSQL_PASSWORD = "123456"
	MYSQL_HOST     = "mysql:3306"
	MYSQL_DB_NAME  = "godb"

	//redis
	REDIS_HOST = "redis"
	REDIS_PORT = "6379"
	REDIS_PASS = "123456"

	BOX_ID_GOLD     = 1 //黄金蛋
	BOX_ID_DIAMONDS = 2 //钻石蛋

	//server
	RunMode       = getEnv("RunMode", "debug")
	WsHttpPort    = getenvInt("Http_Port", "8081")
	AdminHttpPort = getenvInt("Http_Port", "8082")
	ApiHttpPort   = getenvInt("Http_Port", "8083")
	ReadTimeout   = getenvTime("ReadTimeout", "60")
	WriteTimeout  = getenvTime("WriteTimeout", "60")
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getenvInt(key, fallback string) int {
	val := getEnv(key, fallback)
	i, _ := strconv.Atoi(val)
	return i
}

func getenvTime(key, fallback string) time.Duration {
	val := getEnv(key, fallback)
	i, _ := strconv.Atoi(val)
	return time.Duration(i) * time.Second
}
