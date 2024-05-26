if redis.call('get',KEYS[1]) == ARGV[1] then
    -- 说明确实是你的锁，可以释放
    return redis.call('del',KEYS[1]) -- 那就释放这个锁 没问题
else
    -- 说明不是你的锁，不能乱释放
    return 0
end