package main

import (
	"github.com/gomodule/redigo/redis"
	"fmt"
)

func main() {
	pass := redis.DialPassword("zms123456")
	c, err := redis.Dial("tcp", ":6379", pass)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	// reply, err := c.Do("SETNX", "1", 1)
	reply, err := redis.String(c.Do("SETNX", "2", 1))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(reply)
}

