package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, error := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test_db?charset=utf8")
	if error != nil {
		fmt.Println(error.Error())
	}
	defer db.Close()
	result, error := db.Exec("create table user_info(id int(4) not null primary key auto_increment,name char(20) not null)")
	if error != nil {
		fmt.Println(error.Error())
	}
	fmt.Println(result)

}
