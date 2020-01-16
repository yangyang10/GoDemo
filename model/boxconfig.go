/**
 *
 * 标题：
 * 描述：砸蛋配置
 * 作者：黄好杨
 * 创建时间：2020/1/15 2:59 下午
 **/
package model

type Boxconfig struct {
		BoxId int `json:"boxId" binding:"required"`
		Status int	`json:"status" binding:"required"`
		KeyPrice float32 `json:"keyPrice" binding:"required"`
		OpenTime string `json:"openTime" binding:"required"`
		EndTime string  `json:"endTime" binding:"required"`
		ImageUrl string `json:"imageUrl" binding:"required"`
}
