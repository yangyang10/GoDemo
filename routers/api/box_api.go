/**
 *
 * 标题：
 * 描述：
 * 作者：黄好杨
 * 创建时间：2020/3/22 8:18 下午
 **/
package api

import (
	"GoDemo/configs"
	"GoDemo/model"
	"GoDemo/pkg/app"
	"GoDemo/pkg/e"
	"GoDemo/src/dao"
	"GoDemo/src/redis"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*获取砸蛋配置
getBoxConfig?boxId=1
*/
func GetBoxConfig(c *gin.Context) {
	type getBoxConfig struct {
		BoxId int `from:"boxId" binding:"requires"`
	}
	param := getBoxConfig{}
	if err := c.ShouldBind(&param); err != nil {
		app.Response(c, http.StatusBadRequest, e.INVALID_PARAMS, map[string]interface{}{})
		return
	}
	fmt.Printf("接收到的参数是=%d", param.BoxId)
	var requestRedisKey string
	if param.BoxId == configs.BOX_ID_GOLD {
		requestRedisKey = redis.Key.BoxGold.GetKey()
	} else if param.BoxId == configs.BOX_ID_DIAMONDS {
		requestRedisKey = redis.Key.BoxDiamonds.GetKey()
	} else {
		app.Response(c, http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	redisBoxconfig := model.Boxconfig{}
	if err := redis.Client().GetObject(requestRedisKey, &redisBoxconfig); err != nil {
		boxConfig, er := dao.QueryBoxConfig(param.BoxId)
		if er != nil {
			app.Response(c, http.StatusBadRequest, e.INVALID_PARAMS, nil)
			return
		}
		app.Response(c, http.StatusOK, e.SUCCESS, &boxConfig)
		return
	}
	app.Response(c, http.StatusOK, e.SUCCESS, nil)

}

//更新砸蛋配置
func UpdateBoxConfig(c *gin.Context) {
	var boxBean model.Boxconfig
	var resp Resp

	if c.ShouldBindJSON(&boxBean) == nil {
		fmt.Printf("接收到的boxId=%d", boxBean.BoxId)
		if boxBean.BoxId == configs.BOX_ID_GOLD {
			if _, err := redis.Client().Del(redis.Key.BoxGold.GetKey()); err != nil {
				fmt.Printf("移除黄金蛋 redis 失败")
			} else {
				fmt.Printf("移除黄金蛋 redis 成功")
			}

		} else if boxBean.BoxId == configs.BOX_ID_DIAMONDS {
			if _, err := redis.Client().Del(redis.Key.BoxDiamonds.GetKey()); err != nil {
				fmt.Printf("移除钻石蛋 redis 失败")
			} else {
				fmt.Printf("移除钻石蛋 redis 成功")
			}

		}
		if err := dao.UpdateBoxConfig(&boxBean); err == nil {
			resp.Code = "200"
			resp.Msg = "更新成功过"
			c.JSON(http.StatusOK, resp)
			return
		} else {
			resp.Code = "400"
			resp.Msg = err.Error()
			c.JSON(http.StatusOK, resp)
			return
		}
	}
	resp.Code = "400"
	resp.Msg = "参数错误"
	c.JSON(http.StatusOK, resp)
}

type Resp struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

type RespSuccess struct {
	Resp
	Data *model.Boxconfig `json:"data"`
}
