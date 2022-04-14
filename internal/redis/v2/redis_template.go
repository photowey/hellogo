package redis

import (
	"errors"
	"strconv"

	rediz "github.com/gomodule/redigo/redis"
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
	DefaultSession RedizSession // 默认的: {@code redizSession} 空实例
)

var (
	protocol      = "tcp"  // {@code Redis} 协议
	auth          = "AUTH" // {@code Redis} 认证
	emptyString   = ""     // 空字符串
	defaultString = emptyString
)

// ---------------------------------------------------------------- init

func init() {
	DefaultSession = RedizSession{}
}

// ---------------------------------------------------------------- rediz connection-factory interface

// ConnectionFactory {@code Redis} 连接工厂抽象接口
type ConnectionFactory interface {
	OpenConnect() (rediz.Conn, error) // 开启连接
}

// ---------------------------------------------------------------- rediz session interface

// Session {@code Redis} 连接会话抽象接口
type Session interface {
	Borrow() rediz.Conn                                      // 获取连接
	Exec(command string, args ...any) (reply any, err error) // 执行命令
	Release()                                                // 释放连接
}

// ---------------------------------------------------------------- rediz session interface

// Template {@code Redis} 操作模板抽象接口
type Template interface {
	OpenSession() (RedizSession, error)                        //开启会话
	Set(key string, value string) error                        // 设置值
	Setex(key string, value string, expireSeconds int64) error // 设置过期值
	Get(key string) (string, error)                            // 获取值
	LPush(key string, value string) error                      // 左 Push
	RPush(key string, value string) error                      // 右 Push
	LPop(key string) (string, error)                           // 左 弹出
	RPop(key string) (string, error)                           // 右弹出
	Delete(key string) error                                   // 删除
}

// RedizConnectionFactory 简单的 {@code Redis} 连接工厂
type RedizConnectionFactory struct {
	protocol string
	address  string             // {@code Redis} 连接地址 -> 主机:端口
	password string             // {@code Redis} 密码
	options  []rediz.DialOption // {@code Redis} 选项配置
}

// NewConnectionFactory 创建一个连接工厂
// 为什么要将连接工厂的创建暴露出去?
// 1.{@code RedizTemplate} 提供的命令操作还不完善, 这样给外界一个机会 -> 手动创建连接工厂, 然后开启连接
// 2.⭐⭐ 不推荐:开发者使用该方式
func NewConnectionFactory(address string, password string, options ...rediz.DialOption) RedizConnectionFactory {
	return RedizConnectionFactory{
		protocol: protocol,
		address:  address,
		password: password,
		options:  options,
	}
}

// OpenConnect 连接工厂创建连接
// 为外界提供一个机会: 通过连接工厂创建 {@code Redis} 连接
func (factory RedizConnectionFactory) OpenConnect() (rediz.Conn, error) {
	conn, err := rediz.Dial(factory.protocol, factory.address, factory.options...)
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

// ---------------------------------------------------------------- rediz session

// RedizSession {@code Redis} 会话
type RedizSession struct {
	Conn rediz.Conn
}

// Borrow 获取 {@code Redis} 会话连接
func (rs RedizSession) Borrow() rediz.Conn {
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

// ---------------------------------------------------------------- rediz template

// RedizTemplate 操作 {@code Redis} 的模板
// {@code RedizTemplate} 避免采用包名作为结构体前置
type RedizTemplate struct {
	factory *RedizConnectionFactory
}

// NewRedizTemplate 创建 {@code RedizTemplate} 模板实例
func (rt RedizTemplate) NewRedizTemplate(address string, password string, options ...rediz.DialOption) RedizTemplate {
	factory := NewConnectionFactory(address, password, options...)
	return RedizTemplate{
		factory: &factory,
	}
}

// OpenSession 开启 {@code Redis} 会话
func (rt RedizTemplate) OpenSession() (RedizSession, error) {
	conn, err := rt.factory.OpenConnect()
	if err != nil {
		return DefaultSession, err
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
		return errors.New("rediz ConnectionFactory == nil")
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
		return defaultString, errors.New("rediz ConnectionFactory == nil")
	}
	var rs, err = rt.OpenSession()
	if err != nil {
		return defaultString, err
	}

	reply, err := rediz.String(rs.Exec(GET, key))
	if err != nil {
		return defaultString, err
	} else {
		return reply, nil
	}
}

// LPush left push
func (rt RedizTemplate) LPush(key string, value string) error {
	if rt.factory == nil {
		return errors.New("rediz ConnectionFactory == nil")
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
		return errors.New("rediz ConnectionFactory == nil")
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
		return defaultString, errors.New("rediz ConnectionFactory == nil")
	}
	var rs, err = rt.OpenSession()
	if err != nil {
		return defaultString, err
	}

	reply, err := rediz.String(rs.Exec(LPOP, key))
	if err != nil {
		return defaultString, err
	} else {
		return reply, nil
	}
}

// RPop right pop
func (rt RedizTemplate) RPop(key string) (string, error) {
	if rt.factory == nil {
		return defaultString, errors.New("rediz ConnectionFactory == nil")
	}
	var rs, err = rt.OpenSession()
	if err != nil {
		return defaultString, err
	}

	reply, err := rediz.String(rs.Exec(RPOP, key))
	if err != nil {
		return defaultString, err
	} else {
		return reply, nil
	}
}

// Delete the delete action
func (rt RedizTemplate) Delete(key string) error {
	if rt.factory == nil {
		return errors.New("rediz ConnectionFactory == nil")
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
