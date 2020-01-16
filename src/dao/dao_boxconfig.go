/**
 *
 * 标题：
 * 描述：砸蛋配置表操作
 * 作者：黄好杨
 * 创建时间：2020/1/15 2:41 下午
 **/
package dao

import (
	"GoDemo/model"
	"fmt"
)

// 砸蛋配置入库
func InsterBoxConfig(bean *model.Boxconfig) error {
	var sql = "insert into godb_box_config " +
		"(box_id,status,key_price,open_time,end_time,image_url) values(?,?,?,?,?,?)"
	if err := DB.Exec(sql,bean.BoxId,bean.Status,bean.KeyPrice,bean.OpenTime,bean.EndTime,bean.ImageUrl).Error; err != nil {
		fmt.Printf("入库失败 error=%s",err)
		return err
	}
	return nil
}

// 更新砸蛋配置信息
func UpdateBoxConfig(bean *model.Boxconfig) error  {
	result, err := QueryBoxConfig(bean.BoxId)
	if result == nil{
		err = InsterBoxConfig(bean)
	}else{
		var sql = "update godb_box_config set status=?,key_price=?,open_time=?,end_time=?,image_url=? where box_id=?"
		err = DB.Exec(sql,bean.Status,bean.KeyPrice,bean.OpenTime,bean.EndTime,bean.ImageUrl,bean.BoxId).Error
	}
	return err
}

//砸蛋配置查询
func QueryBoxConfig(boxId int) (*model.Boxconfig,error){
	var sql = "select box_id,status,key_price,open_time,end_time,image_url from godb_box_config where box_id =?"
	rows,err := DB.Raw(sql,boxId).Rows()
	defer rows.Close()

	for rows.Next() {
		var bean model.Boxconfig
		DB.ScanRows(rows,&bean)
		return &bean,nil
	}
	return nil,err
}