package db

import "github.com/smartwalle/dbr"

var rPoor = RedisInit()

func RedisInit() *dbr.Pool {
	p := dbr.NewRedis("127.0.0.1:6379","",3,30,5)
	return p
}

func GetRedisSession() *dbr.Session {
	return rPoor.GetSession()
}

