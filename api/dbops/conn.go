package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //注意：程序在操作数据库的时候只需要用到database/sql，而不需要直接使用数据库驱动，所以程序在导入数据库驱动的时候将这个包的名字设置成下划线。
)

var (
	dbConn *sql.DB
	err error

)

func init() {
	//数据源语法："用户名:密码@[连接方式](主机名:端口号)/数据库名"
	//注意：open()在执行时不会真正的与数据库进行连接，只是设置连接数据库需要的参数
	// 【ZKK】注意这里的=，不是:=（表示声明并赋值）
	dbConn, err = sql.Open("mysql", "root:root_root@tcp(localhost:3306)/video?charset=utf8")
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}

}
