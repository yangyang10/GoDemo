/**
 *
 * 标题：
 * 描述：
 * 作者：黄好杨
 * 创建时间：2020/3/22 8:09 下午
 **/
package util

type AppName string

const (
	AppNameApi  = AppName("api")
	AppNamAdmin = AppName("admin")
)

type Role int

const (
	RoleAdmin   = Role(0)   //超级管理员
	RoleBackend = Role(1)   //后端
	RoleUser    = Role(100) //普通用户
)
