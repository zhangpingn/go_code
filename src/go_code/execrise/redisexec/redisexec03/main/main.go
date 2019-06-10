package main

import (
	"go_code/execrise/redisexec/redisexec03/utils"
	_ "fmt"
)

func main() {
	conn := utils.Pool.Get()
	defer conn.Close()
	conn.Do("set", "k1", "v1")
	conn.Do("flushdb")
}