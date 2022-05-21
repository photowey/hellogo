package session

import (
	"github.com/gomodule/redigo/redis"
)

var DefaultSession = RedizSession{}

// RedizSession {@code Redis} 会话
type RedizSession struct {
	conn redis.Conn
}

// OpenSession 开启会话
func OpenSession(conn redis.Conn) RedizSession {
	return RedizSession{
		conn,
	}
}

// Connection 获取 {@code Redis} 会话连接
func (session RedizSession) Connection() redis.Conn {
	return session.conn
}
