/**
 *
 * 标题：
 * 描述：
 * 作者：黄好杨
 * 创建时间：2020/1/17 7:04 下午
 **/
package redis

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

}
