package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

type RoleInfo struct {
	ServerID   int       `json:"server_id" redis:"serverid"`
	Account    string    `json:"account" redis:"-"`
	UserID     uint64    `json:"user_id" redis:"-"`
	RoleName   string    `json:"role_name" redis:"-"`
	Level      int       `json:"level" redis:"-"`
	GuildID    int       `json:"guild_id" redis:"-"`
	ClanID     int       `json:"clan_id" redis:"-"`
	GID        string    `json:"gid" redis:"-"`
	PID        string    `json:"pid" redis:"-"`
	UpdateTime time.Time `json:"update_time" redis:"-"`
	CreateTime int64     `json:"create_time" redis:"-"`
	IsEnable   int `redis:"-"`
}

func main() {
	setPasswd := redis.DialPassword("zms123456")
	c, err := redis.Dial("tcp", ":6379", setPasswd)
	if err != nil {
		// handle error
	}
	defer c.Close()

	v, err := redis.Strings(c.Do("HKEYS", "user_data_new:zms"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
}
