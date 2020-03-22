/**
 *
 * 标题：
 * 描述：
 * 作者：黄好杨
 * 创建时间：2020/3/22 8:55 下午
 **/
package e

var MsgFlags = map[int]string{
	SUCCESS:        "成功",
	ERROR:          "失败",
	INVALID_PARAMS: "提交内容格式或者长度错误",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
