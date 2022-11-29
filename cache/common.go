package cache

import (
	"github.com/go-redis/redis"
	"gopkg.in/ini.v1"
	"strconv"
)

var (
	RedisClient *redis.Client
	RedisDb     string
	RedisAddr   string
	RedisPw     string
	RedisDbName string
)

func init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		panic(err)
	}
	LoadRedis(file)
	Redis()
}

func LoadRedis(file *ini.File) {
	RedisDb = file.Section("redis").Key("RedisDb").String()
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
}

func Redis() {
	db, _ := strconv.ParseUint(RedisDbName, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr: RedisAddr,
		DB:   int(db),
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	RedisClient = client
}
