package factory

import (
	"github.com/gomodule/redigo/redis"
	"github.com/hellogo/internal/redis/session"
	"github.com/hellogo/pkg/logger"
)

var (
	Protocol = "tcp"
	Auth     = "AUTH"
)

// RedisConnectionFactory 简单的 {@code Redis} 连接工厂
type RedisConnectionFactory struct {
	Url      string
	Password string
	Database int
}

// New 创建一个简单的 {@code Redis} 连接工厂
func New(url string, password string, database int) RedisConnectionFactory {
	return RedisConnectionFactory{
		Url:      url,
		Password: password,
		Database: database,
	}
}

// OpenSession 开启连接
func (factory RedisConnectionFactory) OpenSession() (session.RedizSession, error) {
	conn, err := redis.Dial(Protocol, factory.Url, redis.DialDatabase(factory.Database))
	if err != nil {
		logger.Info("connect to redis error", err)
		return session.DefaultSession, err
	}
	if factory.Password != "" {
		_, err := conn.Do(Auth, factory.Password)
		if err != nil {
			return session.DefaultSession, err
		}
	}

	return session.OpenSession(conn), nil
}

// Close 关闭连接
func (factory RedisConnectionFactory) Close(rs session.RedizSession) {
	if rs.Connection() != nil {
		err := rs.Connection().Close()
		if err != nil {
			logger.Info("close to redis connection failed", err)
			return
		}
	}
}
