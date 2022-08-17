package redislock

const (
	lockLua = `local stat = redis.call('GET', KEYS[1]);

-- 无锁时,返回可执行,并标记为写锁中
if not stat then
    redis.call('SETEX', KEYS[1], ARGV[1], 2)
    return 2;
end

-- 无锁,返回可执行,标记为写锁中
if math.abs(tonumber(stat)) < 0.1 then
    redis.call('SETEX', KEYS[1], ARGV[1], 2)
    return 2;
end

-- 写锁定时,返回阻塞
if math.abs(tonumber(stat) - 2) < 0.1 then
    return 3;
end

-- 读锁定时,返回阻塞
if math.abs(tonumber(stat) - 1) < 0.1 then
    return 3;
end

-- 预期之外的结果
return 4;`
	readLockLua = `local stat = redis.call('GET', KEYS[1]);

-- 不存在,无锁时,返回可执行,并标记为读锁中
if not stat then
    redis.call('SETEX', KEYS[1], ARGV[1], 1)
    return 2;
end

-- 存在,但是出于无锁状态,返回可执行,标记为读锁中
if tonumber(stat) == 0 then
    redis.call('SETEX', KEYS[1], ARGV[1], 1)
    return 2;
end

-- 写锁定时,返回阻塞
if tonumber(stat) == 2 then
    return 3;
end

-- 读锁定时,返回放行
if tonumber(stat) == 1 then
    return 2;
end

-- 预期之外的结果
return 4;`
)
