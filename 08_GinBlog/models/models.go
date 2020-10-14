package models

import (
	"blog/pkg/setting"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	user := sec.Key("USER").String()
	password := sec.Key("PASSWORD").String()
	host := sec.Key("HOST").String()
	dbName := sec.Key("NAME").String()
	// tablePrefix := sec.Key("TABLE_PREFIX").String()

	/*
		主要结论:
		在已有字符串数组的场合，使用 strings.Join() 能有比较好的性能
		在一些性能要求较高的场合，尽量使用 buffer.WriteString() 以获得更好的性能
		较少字符串连接的场景下性能最好，而且代码更简短清晰，可读性更好
		如果需要拼接的不仅仅是字符串，还有数字之类的其他需求的话，可以考虑 fmt.Sprintf

	*/
	// dsn example
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	// 1. dsn string handler by fmt.Sprintf(); 内部的逻辑比较复杂，有很多额外的判断，还用到了 interface，所以性能也不是很好
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(2, "mysql canot connect: %v", err)
	}

	// gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	// 	return tablePrefix + defaultTableName
	// }
	fmt.Println(db)
}

// CloseDB info
func CloseDB() {
	// defer db.Close()
}
