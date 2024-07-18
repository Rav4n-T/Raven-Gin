package utils

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"

	g "Raven-Admin/global"
)

type Interface interface {
	Get() bool
	Block(seconds int64) bool
	Release() bool
	ForceRelease()
}

type lock struct {
	context context.Context
	name    string
	owner   string
	seconds int64
}

const releaseLockLuaScript = `
if redis.call("get", KEYS[1]) == ARGV[1] then
	return redis.call("del", KEYS[1])
else
	return 0
end
`

func Lock(name string, seconds int64) Interface {
	return &lock{
		context.Background(),
		name,
		RandString(16),
		seconds,
	}
}

func (lock *lock) Get() bool {
	return g.Redis.SetNX(lock.context, lock.name, lock.owner, time.Duration(lock.seconds)*time.Second).Val()
}

func (lock *lock) Block(seconds int64) bool {
	starting := time.Now().Unix()
	for {
		if !lock.Get() {
			time.Sleep(time.Duration(1) * time.Second)
			if time.Now().Unix()-seconds >= starting {
				return false
			}
		} else {
			return true
		}
	}
}

func (lock *lock) Release() bool {
	luaScript := redis.NewScript(releaseLockLuaScript)
	result := luaScript.Run(lock.context, g.Redis, []string{lock.name}, lock.owner).Val().(int64)
	return result != 0
}

func (lock *lock) ForceRelease() {
	g.Redis.Del(lock.context, lock.name).Val()
}
