/**
 *
 * 标题：
 * 描述：
 * 作者：黄好杨
 * 创建时间：2020/1/12 9:27 下午
 **/
package main

import (
	"GoDemo/configs"
	"GoDemo/model"
	"GoDemo/pkg/util"
	"GoDemo/routers"
	"GoDemo/src/dao"
	"GoDemo/src/redis"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func main() {
	//链接数据库
	dao.Open()
	defer dao.Close()
	//redis
	redis.Client()

	routersInit := routers.InitRouter(util.AppNameApi)
	readTimeout := configs.ReadTimeout
	writeTimeout := configs.WriteTimeout
	endPoint := fmt.Sprintf(":%d", configs.ApiHttpPort)

	server := &http.Server{
		Addr:         endPoint,
		Handler:      routersInit,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}
	fmt.Printf("[info] start http server listening %s\n", endPoint)

	server.ListenAndServe()

}

/*获取砸蛋配置
getBoxConfig?boxId=1
*/
func getBoxConfig(c *gin.Context) {

	boxId, err := strconv.Atoi(c.Query("boxId"))
	fmt.Printf("接收到的参数是=%d", boxId)
	if err == nil {
		var requestRedisKey string
		if boxId == configs.BOX_ID_GOLD {
			requestRedisKey = redis.Key.BoxGold.GetKey()
		} else if boxId == configs.BOX_ID_DIAMONDS {
			requestRedisKey = redis.Key.BoxDiamonds.GetKey()
		} else {
			var resp Resp
			resp.Code = "400"
			resp.Msg = "找不到 box id"
			c.JSON(http.StatusOK, resp)
			return
		}
		redisBoxconfig := model.Boxconfig{}
		if err := redis.Client().GetObject(requestRedisKey, &redisBoxconfig); err != nil {
			boxConfig, er := dao.QueryBoxConfig(boxId)
			if er != nil {
				var resp Resp
				resp.Code = "400"
				resp.Msg = er.Error()
				c.JSON(http.StatusOK, resp)
			} else if boxConfig == nil {
				var resp Resp
				resp.Code = "200"
				resp.Msg = "没有数据"
				c.JSON(http.StatusOK, resp)
			} else {
				var successRes RespSuccess
				successRes.Code = "200"
				successRes.Msg = "获取成功"
				successRes.Data = boxConfig
				c.JSON(http.StatusOK, successRes)
				if err := redis.Client().SetObject(requestRedisKey, &boxConfig); err != nil {
					fmt.Printf("添加到redis错误=%s", err.Error())
				}
			}
		} else {
			var successRes RespSuccess
			successRes.Code = "200"
			successRes.Msg = "redis获取成功"
			successRes.Data = &redisBoxconfig
			c.JSON(http.StatusOK, successRes)
		}
	} else {
		var resp Resp
		resp.Code = "400"
		resp.Msg = "box id 错误"
		c.JSON(http.StatusOK, resp)
	}
}

//更新砸蛋配置
func updateBoxConfig(c *gin.Context) {
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

/**
http.HandleFunc("/demo", testGet)
http.ListenAndServe("0.0.0.0:8083", nil)
*/

//接收x-www-form-urlencoded类型的post请求或者普通get请求
func testGet(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	var result Resp
	test := request.Form["test"][0]
	i, err := strconv.Atoi(test)
	if err == nil {
		fmt.Printf("接收到的参数是=%s", test)
		switch i {
		case 1:
			var us model.User
			us.Avatar = "头像"
			us.Gender = 1
			us.Age = 18
			us.UserName = "小小"
			if err := dao.InsterUser(us); err != nil {
				result.Code = "400"
				result.Msg = "入库失败"
			} else {
				result.Code = "200"
				result.Msg = "入库成功"
			}
		case 2:
			if err := dao.QueryUser(1); err != nil {
				result.Code = "400"
				result.Msg = "查询失败"
			} else {
				result.Code = "200"
				result.Msg = "查询成功"
			}
		case 3:
			if err := dao.UpdateUser(1, "xiaoxiao"); err != nil {
				result.Code = "400"
				result.Msg = "更新失败"
			} else {
				result.Code = "200"
				result.Msg = "更新成功"
			}
		case 4:
			if err := dao.DeleteUser(2); err != nil {
				result.Code = "400"
				result.Msg = "删除失败"
			} else {
				result.Code = "200"
				result.Msg = "删除成功"
			}
		default:
			result.Code = "200"
			result.Msg = "没有进行任何的操作"
		}
	}

	if err := json.NewEncoder(writer).Encode(result); err != nil {
		log.Fatal(err)
	}
}
