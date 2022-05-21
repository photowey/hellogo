package template

import (
	"errors"
	"strconv"

	"github.com/gomodule/redigo/redis"

	"github.com/hellogo/internal/redis/factory"
	"github.com/hellogo/pkg/logger"
)

const (
	LPUSH   string = "LPUSH"
	RPUSH          = "RPUSH"
	LPOP           = "LPOP"
	RPOP           = "RPOP"
	SET            = "SET"
	GET            = "GET"
	DELETE         = "DEL"
	EXPIRED        = "EX"
)

// RedizTemplate 操作 {@code Redis} 的模板
// {@code RedizTemplate} 避免采用包名作为结构体前置
type RedizTemplate struct {
	ConnectionFactory *factory.RedisConnectionFactory
}

// New 创建 {@code RedizTemplate} 模板实例
func (rt RedizTemplate) New(url string, password string, database int) RedizTemplate {
	connectionFactory := factory.New(url, password, database)
	template := RedizTemplate{
		ConnectionFactory: &connectionFactory,
	}

	return template
}

// LPush left push
func (rt RedizTemplate) LPush(key string, value string) error {
	if rt.ConnectionFactory == nil {
		return errors.New("redis ConnectionFactory == nil")
	}
	session, err := rt.ConnectionFactory.OpenSession()
	if err != nil {
		logger.Info("redis set failed:", err)
		return err
	}
	defer rt.ConnectionFactory.Close(session)
	_, err = session.Connection().Do(LPUSH, key, value)
	if err != nil {
		logger.Info("redis set failed:", err)
		return err
	}

	return nil
}

// RPush right push
func (rt RedizTemplate) RPush(key string, value string) error {
	if rt.ConnectionFactory == nil {
		return errors.New("redis ConnectionFactory == nil")
	}
	session, err := rt.ConnectionFactory.OpenSession()
	if err != nil {
		logger.Info("redis set failed:", err)
		return err
	}
	defer rt.ConnectionFactory.Close(session)
	_, err = session.Connection().Do(RPUSH, key, value)
	if err != nil {
		logger.Info("redis set failed:", err)
		return err
	}

	return nil
}

// LPop left pop
func (rt RedizTemplate) LPop(key string) (string, error) {
	if rt.ConnectionFactory == nil {
		return "", errors.New("redis ConnectionFactory == nil")
	}
	session, err := rt.ConnectionFactory.OpenSession()
	defer rt.ConnectionFactory.Close(session)
	if err != nil {
		return "", err
	}
	reply, err := redis.String(session.Connection().Do(LPOP, key))
	if err != nil {
		return "", err
	} else {
		return reply, nil
	}
}

// RPop right pop
func (rt RedizTemplate) RPop(key string) (string, error) {
	if rt.ConnectionFactory == nil {
		return "", errors.New("redis ConnectionFactory == nil")
	}
	session, err := rt.ConnectionFactory.OpenSession()
	defer rt.ConnectionFactory.Close(session)
	if err != nil {
		return "", err
	}
	reply, err := redis.String(session.Connection().Do(RPOP, key))
	if err != nil {
		return "", err
	} else {
		return reply, nil
	}
}

// Set the set action
func (rt RedizTemplate) Set(key string, value string, expireSeconds int64) error {
	if rt.ConnectionFactory == nil {
		return errors.New("redis ConnectionFactory == nil")
	}
	session, err := rt.ConnectionFactory.OpenSession()
	if err != nil {
		logger.Info("redis set failed:", err)
		return err
	}
	defer rt.ConnectionFactory.Close(session)
	_, err = session.Connection().Do(SET, key, value, EXPIRED, strconv.FormatInt(expireSeconds, 10))
	if err != nil {
		logger.Info("redis set failed:", err)
		return err
	}

	return nil
}

// Get the get action
func (rt RedizTemplate) Get(key string) (string, error) {
	if rt.ConnectionFactory == nil {
		return "", errors.New("redis ConnectionFactory == nil")
	}
	session, err := rt.ConnectionFactory.OpenSession()
	defer rt.ConnectionFactory.Close(session)
	if err != nil {
		return "", err
	}
	reply, err := redis.String(session.Connection().Do(GET, key))
	if err != nil {
		return "", err
	} else {
		return reply, nil
	}
}

// Delete the delete action
func (rt RedizTemplate) Delete(key string) error {
	if rt.ConnectionFactory == nil {
		return errors.New("redis ConnectionFactory == nil")
	}
	session, err := rt.ConnectionFactory.OpenSession()
	if err != nil {
		logger.Info("redis set failed:", err)
		return err
	}
	defer rt.ConnectionFactory.Close(session)
	_, err = session.Connection().Do(DELETE, key)
	if err != nil {
		logger.Info("redis set failed:", err)
		return err
	}

	return nil
}

// other cmd...
