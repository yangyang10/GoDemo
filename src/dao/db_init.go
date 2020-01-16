/**
 *
 * 标题：
 * 描述：gorm文档
	<p>
		https://gorm.io/docs/connecting_to_the_database.html
	</p>
 * 作者：黄好杨
 * 创建时间：2020/1/12 5:57 下午
 **/
package dao

import (
	"GoDemo/configs"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"runtime/debug"
)

var DB *gorm.DB

//开启连接数据库
func Open() {

	d, err := gorm.Open("mysql", fmt.Sprintf(
		"%s:%s@(%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		configs.MYSQL_USER,
		configs.MYSQL_PASSWORD,
		configs.MYSQL_HOST,
		configs.MYSQL_DB_NAME))

	if err != nil{
		log.Fatal(err)
		panic("链接数据库失败，请检查配置文件")
	}

	fmt.Println("[info] mysql starting")
	DB = d

}

//关闭数据库
func Close()  {
	if DB != nil{
		DB.Close()
	}
}

//开始数据库事务，错误自动回滚
func BeginTransaction(cb func(*gorm.DB) error) error {

	transaction := DB.Begin()
	if err := transaction.Error; err != nil{
		return err
	}

	defer func() {
		if err := recover();err != nil{
			log.Printf("事务错误=%v",err)
			debug.PrintStack()
			if err := transaction.Rollback().Error; err != nil{
				log.Printf("事务回滚错误=%v",err)
			}
		}
	}()

	if err := cb(transaction); err != nil{
		transaction.Rollback()
		return err
	}else if err := transaction.Commit().Error;err != nil{
		return err
	}
	return nil
}
