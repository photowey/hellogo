package redis

import (
	"errors"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

// ---------------------------------------------------------------- const

const (
	SET     string = "SET"
	GET            = "GET"
	LPUSH          = "LPUSH"
	RPUSH          = "RPUSH"
	LPOP           = "LPOP"
	RPOP           = "RPOP"
	DELETE         = "DEL"
	EXPIRED        = "EX"
)

// ---------------------------------------------------------------- var

var (
	defaultSession redizSession
)

var (
	protocol      = "tcp"
	auth          = "AUTH"
	emptyString   = ""
	defaultString = emptyString
)

// ---------------------------------------------------------------- init

func init() {
	defaultSession = redizSession{}
}

// ---------------------------------------------------------------- redis connection factory

// redizConnectionFactory 简单的 {@code Redis} 连接工厂
type redizConnectionFactory struct {
	protocol string
	address  string             // {@code Redis} 连接地址 -> 主机:端口
	password string             // {@code Redis} 密码
	options  []redis.DialOption // {@code Redis} 选项配置
}

// openConnect 连接工厂创建连接
func (factory redizConnectionFactory) openConnect() (redis.Conn, error) {
	conn, err := redis.Dial(factory.protocol, factory.address, factory.options...)
	if err != nil {
		return nil, err
	}
	if factory.password != "" {
		_, err := conn.Do(auth, factory.password)
		if err != nil {
			return nil, err
		}
	}

	return conn, nil
}

// ---------------------------------------------------------------- redis session

// redizSession {@code Redis} 会话
type redizSession struct {
	conn redis.Conn
}

// Borrow 获取 {@code Redis} 会话连接
func (rs redizSession) borrow() redis.Conn {
	return rs.conn
}

// exec 执行 {@code Redis} 命令
func (rs redizSession) exec(command string, args ...any) (reply any, err error) {

	conn := rs.borrow()
	return conn.Do(command, args)
}

// Release 释放连接
func (rs redizSession) release() {
	err := rs.conn.Close()
	if err != nil {
		return
	}
}

// ---------------------------------------------------------------- redis template

// RedizTemplate 操作 {@code Redis} 的模板
// {@code RedizTemplate} 避免采用包名作为结构体前置
type RedizTemplate struct {
	factory *redizConnectionFactory
}

// NewRedizTemplate 创建 {@code RedizTemplate} 模板实例
func (rt RedizTemplate) NewRedizTemplate(address string, password string, options ...redis.DialOption) RedizTemplate {
	return RedizTemplate{
		factory: &redizConnectionFactory{
			protocol: protocol,
			address:  address,
			password: password,
			options:  options,
		},
	}
}

// openSession 开启 {@code Redis} 会话
func (rt RedizTemplate) openSession() (redizSession, error) {
	conn, err := rt.factory.openConnect()
	if err != nil {
		return defaultSession, err
	}

	return redizSession{conn}, nil
}

// Set 设置值
func (rt RedizTemplate) Set(key string, value string) error {
	return rt.Setex(key, value, -1)
}

// Setex 设置值,并设置一个过期时间
func (rt RedizTemplate) Setex(key string, value string, expireSeconds int64) error {
	if rt.factory == nil {
		return errors.New("redis ConnectionFactory == nil")
	}
	var rs, err = rt.openSession()
	if err != nil {
		return err
	}

	if expireSeconds < 0 {
		_, err = rs.exec(SET, key, value)
	} else {
		_, err = rs.exec(SET, key, value, EXPIRED, strconv.FormatInt(expireSeconds, 10))
	}

	if err != nil {
		return err
	}

	return nil
}

// Get the get action
func (rt RedizTemplate) Get(key string) (string, error) {
	if rt.factory == nil {
		return defaultString, errors.New("redis ConnectionFactory == nil")
	}
	var rs, err = rt.openSession()
	if err != nil {
		return defaultString, err
	}

	reply, err := redis.String(rs.exec(GET, key))
	if err != nil {
		return defaultString, err
	} else {
		return reply, nil
	}
}

// LPush left push
func (rt RedizTemplate) LPush(key string, value string) error {
	if rt.factory == nil {
		return errors.New("redis ConnectionFactory == nil")
	}
	var rs, err = rt.openSession()
	if err != nil {
		return err
	}

	_, err = rs.exec(LPUSH, key, value)
	if err != nil {
		return err
	}

	return nil
}

// RPush right push
func (rt RedizTemplate) RPush(key string, value string) error {
	if rt.factory == nil {
		return errors.New("redis ConnectionFactory == nil")
	}
	var rs, err = rt.openSession()
	if err != nil {
		return err
	}

	_, err = rs.exec(RPUSH, key, value)
	if err != nil {
		return err
	}

	return nil
}

// LPop left pop
func (rt RedizTemplate) LPop(key string) (string, error) {
	if rt.factory == nil {
		return defaultString, errors.New("redis ConnectionFactory == nil")
	}
	var rs, err = rt.openSession()
	if err != nil {
		return defaultString, err
	}

	reply, err := redis.String(rs.exec(LPOP, key))
	if err != nil {
		return defaultString, err
	} else {
		return reply, nil
	}
}

// RPop right pop
func (rt RedizTemplate) RPop(key string) (string, error) {
	if rt.factory == nil {
		return defaultString, errors.New("redis ConnectionFactory == nil")
	}
	var rs, err = rt.openSession()
	if err != nil {
		return defaultString, err
	}

	reply, err := redis.String(rs.exec(RPOP, key))
	if err != nil {
		return defaultString, err
	} else {
		return reply, nil
	}
}

// Delete the delete action
func (rt RedizTemplate) Delete(key string) error {
	if rt.factory == nil {
		return errors.New("redis ConnectionFactory == nil")
	}
	var rs, err = rt.openSession()
	if err != nil {
		return err
	}

	_, err = rs.exec(DELETE, key)
	if err != nil {
		return err
	}

	return nil
}

// other cmd...
