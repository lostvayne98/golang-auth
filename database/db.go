package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Main() {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("Подключение к базе данных")
}
func InsertUser(name, phone, password string) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("Подключение к базе данных")

	insert, err := db.Query(fmt.Sprintf("INSERT INTO `users` (`name`,`phone`,`password`,`role`)VALUES ('%s','%s','%s','4')", name, phone, password))
	if err != nil {
		fmt.Println(err)
	}
	defer insert.Close()
}
func GetUser(phone, password string) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/golang")
	if err != nil {
		panic("Не могу подключиться")
	}

	fmt.Println("Проверяю пользователя")

	get, err := db.Query(fmt.Sprintf("SELECT * FROM `users` WHERE `phone`='%s' and `password` = '%s' ", phone, password))
	if err != nil {
		fmt.Println("err")
	}
	defer get.Close()

}
