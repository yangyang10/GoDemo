/**
 *
 * 标题：
 * 描述：
 * 作者：黄好杨
 * 创建时间：2020/3/22 8:07 下午
 **/
package routers

import (
	"GoDemo/pkg/util"
	"GoDemo/routers/api"
	"github.com/gin-gonic/gin"
)

func InitRouter(appName util.AppName) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//

	if appName == util.AppNameApi {
		//客户端
		r.GET("/getBoxConfig", api.GetBoxConfig)
		r.POST("/updateBoxConfig", api.UpdateBoxConfig)
	}

	return r
}
