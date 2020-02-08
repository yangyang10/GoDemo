/**
 *
 * 标题：
 * 描述：
 * 作者：黄好杨
 * 创建时间：2020/1/17 7:04 下午
 **/
package redis

import (
	"fmt"
	"github.com/fatih/camelcase"
	"reflect"
	"strings"
)

const (
	RedisPrefix = "timel:"
)

type KeyValue struct {
	value string
}

func NewKey(value string) KeyValue {
	return KeyValue{value: value}
}

func NewKeyStore() RedisKey {
	redisKey := RedisKey{}
	t := reflect.TypeOf(redisKey)
	v := reflect.ValueOf(&redisKey).Elem()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		value := NewKey(strings.ToLower(strings.Join(camelcase.Split(f.Name), ":")))
		v.FieldByName(f.Name).Set(reflect.ValueOf(value))
	}
	return redisKey
}

func (k *KeyValue) GetKey(sufix ...interface{}) string {
	key := RedisPrefix + k.value
	for _, suf := range sufix {
		key = fmt.Sprintf("%v:%v", key, suf)
	}

	return key
}

var Key = NewKeyStore()
