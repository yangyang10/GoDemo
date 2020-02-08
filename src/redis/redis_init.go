/**
 *
 * 标题：
 * 描述：redis 链接，操作
 * 作者：黄好杨
 * 创建时间：2020/1/17 4:00 下午
 **/
package redis

import (
	"GoDemo/configs"
	"encoding/json"
	"fmt"
	"github.com/go-redis/cache"
	"github.com/go-redis/redis"
	"time"
)

type Redis struct {
	client *redis.Client
	cache  *cache.Codec
}

var client *Redis

func Client() *Redis {
	if client == nil {
		client = New()
	}

	return client
}

func New() *Redis {
	r := &Redis{
		client: redis.NewClient(&redis.Options{
			Addr:     configs.REDIS_HOST + ":" + configs.REDIS_PORT,
			Password: configs.REDIS_PASS,
			DB:       0,
		}),
	}

	r.cache = &cache.Codec{
		Redis: r.client,
		Marshal: func(i interface{}) (bytes []byte, e error) {
			return json.Marshal(i)
		},
		Unmarshal: func(bytes []byte, i interface{}) error {
			return json.Unmarshal(bytes, i)
		},
	}

	var _, err = r.client.Ping().Result()
	if err != nil {
		fmt.Println("redis connect fail")
	} else {
		fmt.Println("[info] redis starting")
	}

	return r
}

func (r *Redis) RedisClient() *redis.Client {
	return r.client
}

//移除redis key
func (r *Redis) Del(keys ...string) (int64, error) {
	return r.client.Del(keys...).Result()
}

func (r *Redis) Get(key string) (string, error) {
	return r.client.Get(key).Result()
}

//获取hash key 值
func (r *Redis) HGetString(key, field string) (string, error) {
	return r.client.HGet(key, field).Result()
}

//对象json序列化后存储到redis
func (r *Redis) SetObject(key string, obj interface{}, exp ...time.Duration) error {
	expiration := time.Hour * 24 * 366 * 100 //默认100年过期
	if len(exp) > 0 {
		expiration = exp[0]
	}

	return r.cache.Set(&cache.Item{
		Key:        key,
		Object:     obj,
		Expiration: expiration,
	})
}

func (r *Redis) GetObject(key string, obj interface{}) error {
	return r.cache.Get(key, obj)
}
