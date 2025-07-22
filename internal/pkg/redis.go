package pkg

import (
	"context"
	"fmt"
	"sync"

	"github.com/mediocregopher/radix/v4"
	"github.com/spf13/viper"
)

var (
	redisOnce sync.Once
	redisConn radix.Conn
	redisErr  error
)

// GetRedis 返回全局 Redis 连接
func GetRedis() (radix.Conn, error) {
	redisOnce.Do(func() {
		addr := viper.GetString("redis.addr")
		password := viper.GetString("redis.password")
		db := viper.GetInt("redis.db")

		redisConn, redisErr = radix.Dial(context.Background(), "tcp", addr)
		if redisErr != nil {
			return
		}
		if password != "" {
			if err := redisConn.Do(context.Background(), radix.Cmd(nil, "AUTH", password)); err != nil {
				redisConn.Close()
				redisConn = nil
				redisErr = err
				return
			}
		}
		if db > 0 {
			if err := redisConn.Do(context.Background(), radix.Cmd(nil, "SELECT", fmt.Sprintf("%d", db))); err != nil {
				redisConn.Close()
				redisConn = nil
				redisErr = err
				return
			}
		}
	})
	return redisConn, redisErr
}

// PingRedis 用于健康检查
func PingRedis() error {
	conn, err := GetRedis()
	if err != nil {
		return err
	}
	var resp string
	err = conn.Do(context.Background(), radix.Cmd(&resp, "PING"))
	if err != nil {
		return err
	}
	if resp != "PONG" {
		return fmt.Errorf("unexpected redis ping response: %s", resp)
	}
	return nil
}
