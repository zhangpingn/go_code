package utils

import (
	"github.com/garyburd/redigo/redis"
)

var (
	Pool *redis.Pool
)
func init(){
	Pool = &redis.Pool{
		MaxActive : 8,
		Dial : func()(conn redis.Conn, err error){
			conn, err = redis.Dial("tcp", "localhost:6379")
			return 
		},
	}
}