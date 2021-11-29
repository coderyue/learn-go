package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3308)/test")
	if err != nil {
		fmt.Println("mysql connect fail");
		return
	}
	defer db.Close()

	// 插入一条数据
	//sql := "insert into user (id, user, name) values('testid', 'testuser', 'ruby')"
	//exec, _:= db.Exec(sql)
	//rowsAffected, _ := exec.RowsAffected()
	//fmt.Println(">>>>>>插入条数：", rowsAffected)

	// 预处理
	//user := [2][3]string{{"555", "tom", "tomName"}, {"666", "berry", "berryName"}}
	//preStmt, _ := db.Prepare("insert into user values (?, ?, ?)")   // 获取预处理对象
	//
	//for _, s := range user{
	//	preStmt.Exec(s[0], s[1], s[2]) 			// 调用预处理语句
	//}

	// 查询一条数据
	//var id, user, name string
	//sql := "select id, user, name from user"
	//row := db.QueryRow(sql)
	//row.Scan(&id, &user, &name)
	//fmt.Println(id, user, name)

	// 查询多条数据
	sql := "select id, user, name from user"
	rows, _ := db.Query(sql)
	var id, user, name string

	for rows.Next() {		// 循环拿到所有数据
		rows.Scan(&id, &user, &name)
		fmt.Println(id, user, name)
	}

}
