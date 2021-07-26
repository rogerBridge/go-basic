package main

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

// 从mysql数据库中拿到所需要的数据
func getEcCommand(vcuNo string, bgnTime string, endTime string) []string {
	// sql.DB 设计之初是为了保持长时间的连接的
	db, err := sql.Open("mysql", "global:senthink@tcp(10.16.10.246:3306)/columbia_global")
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Ping()
	if err != nil {
		log.Println("mysql连接失败~", err)
	}
	defer db.Close()
	// 定义需要查找的变量名
	var cmdDetail string
	output := make([]string, 0, 1000)
	rows, err := db.Query("select cmd_detail from bump_dvc_dev_log where cmd_value = ? and vcuno = ? and cmd_time>? and cmd_time<?", "0x0A", vcuNo, bgnTime, endTime)
	if err != nil {
		log.Fatalln("query mysql数据库出现错误!", err)
	}
	defer rows.Close()
	// 遍历rows里面符合vcuno要求的row
	for rows.Next() {
		err := rows.Scan(&cmdDetail)
		if err != nil {
			log.Fatalln(err)
		}
		output = append(output, cmdDetail)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return output
}

func main() {
	t0 := time.Now()
	fmt.Println(getEcCommand("V1AB0100A19I240001", "2019-10-08 16:22:00", "2019-10-11 16:47:00"))
	fmt.Println("获取此次数据共用时:", time.Since(t0))
}
