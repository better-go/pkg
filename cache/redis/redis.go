package redis

/*

ref:

	- redis 官方 sdk 列表: https://redis.io/clients#go
	- 讨论: https://www.reddit.com/r/golang/comments/3dspr9/which_redis_client/
	- redis 客户端lib 选择: https://zhuanlan.zhihu.com/p/47480504
	- https://segmentfault.com/a/1190000007078961
	- 分布式锁: https://github.com/bsm/redislock

go 可选包:
	- https://github.com/gomodule/redigo
		- 支持所有 redis cmd
		- 部分大厂选择, 生产验证可靠
	- https://github.com/go-redis/redis
		- 接口设计友好

*/

import (
	sdk1 "github.com/go-redis/redis/v7"
	sdk2 "github.com/gomodule/redigo/redis"
)

// redis client:
type Client struct {
	v1 *sdk1.Client // use client
	v2 *sdk2.Pool   // use pool
}

func NewClient(opt *Options) *Client {
	return &Client{
		v1: NewRedis(opt),
		v2: NewRedisFull(),
	}
}

func NewRedis(opt *Options) *sdk1.Client {
	optV1 := opt.ToOption1()

	// client with pool
	client := sdk1.NewClient(optV1)
	return client
}

func NewRedisFull() *sdk2.Pool {
	return sdk2.NewPool(nil, 2)
}

func (m *Client) V1() *sdk1.Client {
	return m.v1
}

func (m *Client) V2() *sdk2.Pool {
	return m.v2
}
