package redislock

import (
	"time"

	"github.com/fwhezfwhez/errorx"
	"github.com/garyburd/redigo/redis"
)

type RedisLock struct {
	key           string
	maxLockSecond int           // 锁定状态标记的最大时间
	maxRetryTimes int           // 最大阻塞重试次数
	retryInterval time.Duration // 重试的间隔
}

func NewRedisRWMutex(key string, maxLockSecond int, maxRetryTimes int, retryInterval time.Duration) *RedisLock {
	return &RedisLock{
		key:           key,
		maxLockSecond: maxLockSecond,
		maxRetryTimes: maxRetryTimes,
		retryInterval: retryInterval,
	}
}

func (lock *RedisLock) RLock(conn redis.Conn) (int, error) {
	maxRetry := lock.maxRetryTimes

	rs, e := lock.rlock(&maxRetry, conn)

	if e != nil {
		return 0, errorx.Wrap(e)
	}

	return rs, nil
}

func (lock *RedisLock) rlock(retryTimes *int, conn redis.Conn) (int, error) {
	/*
		--- 读锁定时,锁状态置为1,不做任何阻塞
		--- 写锁定时, 锁状态置为2, 阻塞其他读写
		--- 无锁时,锁状态为0,或者不存在该 key

		--- 返回2表示可执行
		--- 返回3表示需要阻塞
	*/
	var script = `
local stat = redis.call('GET', KEYS[1]);
if not stat then
    redis.call('SETEX', KEYS[1], ARGV[1], 1)
    return 2;
end
if tonumber(stat) == 0 then
    redis.call('SETEX', KEYS[1], ARGV[1], 1)
    return 2;
end
if tonumber(stat) == 2 then
    return 3;
end
if tonumber(stat) == 1 then
    return 2;
end
return 4;
`
	vint, e := redis.Int(conn.Do("eval", script, 1, lock.key, lock.maxLockSecond))
	if e != nil {
		return 0, errorx.Wrap(e)
	}

	// 可执行
	if vint == 2 {
		return 0, nil
	}

	if vint == 3 {

		*retryTimes--
		if *retryTimes == 0 {
			return 1, nil
		}

		time.Sleep(lock.retryInterval)
		return lock.rlock(retryTimes, conn)
	}

	return 0, errorx.NewFromStringf("unexpected lock stat return %d", vint)
}

func (lock *RedisLock) RUnLock(conn redis.Conn) {
	_, _ = conn.Do("del", lock.key)
}

func (lock RedisLock) Lock(conn redis.Conn) (int, error) {
	var max = lock.maxRetryTimes

	rs, e := lock.lock(&max, conn)
	if e != nil {
		return 0, errorx.Wrap(e)
	}

	return rs, nil
}

func (lock *RedisLock) lock(retryTimes *int, conn redis.Conn) (int, error) {
	/*
	   --- 读锁定时,锁状态置为1,不做任何阻塞
	   --- 写锁定时,锁状态置为2,阻塞其他读写
	   --- 无锁时,锁状态为0,或者不存在该 key

	   --- 返回 3表示需要阻塞
	   --- 返回 2表示可执行
	*/
	var script = `
local stat = redis.call('GET', KEYS[1]);
if not stat then
    redis.call('SETEX', KEYS[1], ARGV[1], 2)
    return 2;
end
if math.abs(tonumber(stat)) < 0.1 then
    redis.call('SETEX', KEYS[1], ARGV[1], 2)
    return 2;
end
if math.abs(tonumber(stat) - 2) < 0.1 then
    return 3;
end
if math.abs(tonumber(stat) - 1) < 0.1 then
    return 3;
end
return 4;
`
	vint, e := redis.Int(conn.Do("eval", script, 1, lock.key, lock.maxLockSecond))
	if e != nil {
		return 0, errorx.Wrap(e)
	}

	// 可执行
	if vint == 2 {
		return 0, nil
	}

	if vint == 3 {

		*retryTimes--
		if *retryTimes == 0 {
			return 1, nil
		}
		time.Sleep(lock.retryInterval)

		return lock.lock(retryTimes, conn)
	}

	return 0, errorx.NewFromStringf("unexpected lock stat return %d", vint)
}

func (lock *RedisLock) UnLock(conn redis.Conn) {
	_, _ = conn.Do("del", lock.key)
}
