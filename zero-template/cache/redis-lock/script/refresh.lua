if redis.call('get',KEYS[1]) == ARGV[1] then
    -- 说明确实是你的锁，可以续期
    return redis.call('expire',KEYS[1],ARGV[2])
else
    -- 说明不是你的锁，不能续期
    return 0
end