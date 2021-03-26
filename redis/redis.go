package redis

import (
	"context"
	"fmt"

	"github.com/SevenTV/ServerGo/configure"
	"github.com/go-redis/redis/v8"
	"github.com/gobuffalo/packr/v2"
)

var Ctx = context.Background()

var (
	errInvalidResp = fmt.Errorf("invalid resp from redis")
)

func init() {
	options, err := redis.ParseURL(configure.Config.GetString("redis_uri"))
	if err != nil {
		panic(err)
	}

	Client = redis.NewClient(options)

	box := packr.New("lua", "./lua")

	tokenConsumerLuaScript, err := box.FindString("token-consumer.lua")
	if err != nil {
		panic(err)
	}
	v, err := Client.ScriptLoad(Ctx, tokenConsumerLuaScript).Result()
	if err != nil {
		panic(err)
	}
	tokenConsumerLuaScriptSHA1 = v

	getCacheLuaScript, err := box.FindString("get-cache.lua")
	if err != nil {
		panic(err)
	}
	v, err = Client.ScriptLoad(Ctx, getCacheLuaScript).Result()
	if err != nil {
		panic(err)
	}
	getCacheLuaScriptSHA1 = v

	setCacheLuaScript, err := box.FindString("set-cache.lua")
	if err != nil {
		panic(err)
	}
	v, err = Client.ScriptLoad(Ctx, setCacheLuaScript).Result()
	if err != nil {
		panic(err)
	}
	setCacheLuaScriptSHA1 = v

	invalidateCacheLuaScript, err := box.FindString("invalidate-cache.lua")
	if err != nil {
		panic(err)
	}
	v, err = Client.ScriptLoad(Ctx, invalidateCacheLuaScript).Result()
	if err != nil {
		panic(err)
	}
	invalidateCacheLuaScriptSHA1 = v

	invalidateCommonIndexCacheLuaScript, err := box.FindString("invalidate-common-index-cache.lua")
	if err != nil {
		panic(err)
	}
	v, err = Client.ScriptLoad(Ctx, invalidateCommonIndexCacheLuaScript).Result()
	if err != nil {
		panic(err)
	}
	invalidateCommonIndexCacheLuaScriptSHA1 = v
}

var Client *redis.Client

type Message = redis.Message

type StringCmd = redis.StringCmd

type StringStringMapCmd = redis.StringStringMapCmd

const ErrNil = redis.Nil
