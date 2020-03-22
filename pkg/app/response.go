/**
 *
 * 标题：
 * 描述：
 * 作者：黄好杨
 * 创建时间：2020/3/22 8:34 下午
 **/
package app

import (
	"GoDemo/pkg/e"
	"github.com/gin-gonic/gin"
	"time"
)

type Gin struct {
	C *gin.Context
}

type Result struct {
	Code       int         `json:"code"`
	Msg        string      `json:"msg"`
	AppVersion string      `json:"app_version"`
	Data       interface{} `json:"data"`
}

func Response(c *gin.Context, httpCode, errCode int, data interface{}) {
	if data == nil {
		data = ""
	}

	result := Result{
		Code: httpCode,
		Msg:  e.GetMsg(errCode),
		//AppVersion: "",
		Data: data,
	}

	//TODO 数据加密

	c.JSONP(httpCode, result)

	//记录运行时间

	go func() {
		if _, ok := c.Keys["enter"]; !ok {
			return
		}
		enter := c.Keys["enter"].(time.Time)
		t := time.Now().Sub(enter).Seconds()
		if t > 1.5 {
			//TODO 接口请求超过1。5秒进行日志记录
		}
	}()

	return
}
