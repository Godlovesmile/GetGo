package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// sql.Open()中的数据库连接串格式为："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	db, err := sql.Open("mysql", "root:12345678@tcp(47.104.156.147:3306)/ylzdata?charset=utf8")
	checkErr(err)

	// 插入数据
	// insertData(db)

	// 更新数据
	// updateData(db)

	// 查询数据
	// searchData(db)

	// 删除数据
	deleteData(db)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func insertData(db *sql.DB) {
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("ylz", "FE", "2020-9-17")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
}

func updateData(db *sql.DB) {
	stmt, err := db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err := stmt.Exec("xiaohua", 1)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}

func searchData(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string

		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}
}

func deleteData(db *sql.DB) {
	stmt, err := db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err := stmt.Exec(1)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()
}
