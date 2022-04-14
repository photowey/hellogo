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
	defaultSession RedizSession
)

var (
	protocol      = "tcp"
	auth          = "AUTH"
	emptyString   = ""
	defaultString = emptyString
)

// ---------------------------------------------------------------- init

func init() {
	defaultSession = RedizSession{}
}

// ---------------------------------------------------------------- redis connection factory

// RedizConnectionFactory 简单的 {@code Redis} 连接工厂
type RedizConnectionFactory struct {
	protocol string
	address  string             // {@code Redis} 连接地址 -> 主机:端口
	password string             // {@code Redis} 密码
	options  []redis.DialOption // {@code Redis} 选项配置
}

// NewConnectionFactory 创建一个连接工厂
// 为什么要将连接工厂的创建暴露出去?
// 1.{@code RedizTemplate} 提供的命令操作还不完善, 这样给外界一个机会 -> 手动创建连接工厂, 然后开启连接
// 2.⭐⭐ 不推荐:开发者使用该方式
func NewConnectionFactory(address string, password string, options ...redis.DialOption) RedizConnectionFactory {
	return RedizConnectionFactory{
		protocol: protocol,
		address:  address,
		password: password,
		options:  options,
	}
}

// OpenConnect 连接工厂创建连接
// 为外界提供一个机会: 通过连接工厂创建 {@code Redis} 连接
func (factory RedizConnectionFactory) OpenConnect() (redis.Conn, error) {
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

// RedizSession {@code Redis} 会话
type RedizSession struct {
	Conn redis.Conn
}

// Borrow 获取 {@code Redis} 会话连接
func (rs RedizSession) Borrow() redis.Conn {
	return rs.Conn
}

// Exec 执行 {@code Redis} 命令
func (rs RedizSession) Exec(command string, args ...any) (reply any, err error) {
	defer rs.Release()
	conn := rs.Borrow()
	return conn.Do(command, args)
}

// Release 释放连接
func (rs RedizSession) Release() {
	err := rs.Conn.Close()
	if err != nil {
		return
	}
}

// ---------------------------------------------------------------- redis template

// RedizTemplate 操作 {@code Redis} 的模板
// {@code RedizTemplate} 避免采用包名作为结构体前置
type RedizTemplate struct {
	factory *RedizConnectionFactory
}

// NewRedizTemplate 创建 {@code RedizTemplate} 模板实例
func (rt RedizTemplate) NewRedizTemplate(address string, password string, options ...redis.DialOption) RedizTemplate {
	factory := NewConnectionFactory(address, password, options...)
	return RedizTemplate{
		factory: &factory,
	}
}

// OpenSession 开启 {@code Redis} 会话
func (rt RedizTemplate) OpenSession() (RedizSession, error) {
	conn, err := rt.factory.OpenConnect()
	if err != nil {
		return defaultSession, err
	}

	return RedizSession{conn}, nil
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
	var rs, err = rt.OpenSession()
	if err != nil {
		return err
	}

	if expireSeconds < 0 {
		_, err = rs.Exec(SET, key, value)
	} else {
		_, err = rs.Exec(SET, key, value, EXPIRED, strconv.FormatInt(expireSeconds, 10))
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
	var rs, err = rt.OpenSession()
	if err != nil {
		return defaultString, err
	}

	reply, err := redis.String(rs.Exec(GET, key))
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
	var rs, err = rt.OpenSession()
	if err != nil {
		return err
	}

	_, err = rs.Exec(LPUSH, key, value)
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
	var rs, err = rt.OpenSession()
	if err != nil {
		return err
	}

	_, err = rs.Exec(RPUSH, key, value)
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
	var rs, err = rt.OpenSession()
	if err != nil {
		return defaultString, err
	}

	reply, err := redis.String(rs.Exec(LPOP, key))
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
	var rs, err = rt.OpenSession()
	if err != nil {
		return defaultString, err
	}

	reply, err := redis.String(rs.Exec(RPOP, key))
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
	var rs, err = rt.OpenSession()
	if err != nil {
		return err
	}

	_, err = rs.Exec(DELETE, key)
	if err != nil {
		return err
	}

	return nil
}

// other cmd...
