package config

import (
	"BaseWebProject/cache"
	"BaseWebProject/env"
	"BaseWebProject/logger"
	"BaseWebProject/model"
	"gopkg.in/ini.v1"
	"strings"
)

func Init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		logger.Logger.Info("配置文件加载错误，请检查文件路径")
	}
	loadSever(file)
	loadRedisConfig(file)
	loadMysqlConfig(file)
	cache.InitRedis()
	path := strings.Join([]string{env.DbUser, ":", env.DbPassword, "@tcp(", env.DbHost, ":", env.DbPort, ")/", env.DbName, "?charset=utf8&parseTime=true"}, "")
	model.InitMysql(path)
}

func loadSever(file *ini.File) {
	env.HttpPort = file.Section("service").Key("HttpPort").String()
}

func loadRedisConfig(file *ini.File) {
	env.RedisAddr = file.Section("redis").Key("RedisAddr").String()
	env.RedisPw = file.Section("redis").Key("RedisPw").String()
	env.RedisDbName = file.Section("redis").Key("RedisDbName").String()
}

func loadMysqlConfig(file *ini.File) {
	env.DbHost = file.Section("mysql").Key("DbHost").String()
	env.DbPort = file.Section("mysql").Key("DbPort").String()
	env.DbUser = file.Section("mysql").Key("DbUser").String()
	env.DbPassword = file.Section("mysql").Key("DbPassword").String()
	env.DbName = file.Section("mysql").Key("DbName").String()
}
