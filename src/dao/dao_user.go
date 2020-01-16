/**
 *
 * 标题：
 * 描述：用户表数据库dao操作
 * 作者：黄好杨
 * 创建时间：2020/1/12 6:28 下午
 **/
package dao

import (
	"GoDemo/model"
	"fmt"
)

//插入数据库,model映射
func InsterUserByInterface(user *model.User) error {
	if err := DB.Create(user).Error; err != nil{
		return err
	}
	return nil
}

//sql插入数据库
func InsterUser(user model.User) error {
	if err := DB.Exec("insert into user (user_name,avatar,gender,age) values(?,?,?,?)",
		user.UserName,user.Avatar,user.Gender,user.Age).Error; err != nil{
			fmt.Printf("入库失败 error=%s",err)
			return err
	}
	return nil
}

//用户名查询
func QueryUser(queryId int64)error  {
	if err := DB.Exec("select user_id,user_name,avatar,gender,age from user where user_id=? ",
		queryId).Error; err != nil{
		fmt.Printf("查询语句报错 error=%s",err)
		return err
	}
	return nil
}

func DeleteUser(userId int64) error {
	if err := DB.Exec("delete from user where user_id=?",
		userId).Error; err != nil{
		fmt.Printf("删除语句报错 error=%s",err)
		return err
	}
	return nil
}

func UpdateUser(userId int64,userName string) error {
	if err := DB.Exec("update  user set user_name=? where user_id=?",
		userName,userId).Error; err != nil{
		fmt.Printf("更新用户语句报错 error=%s",err)
		return err
	}
	return nil
}
