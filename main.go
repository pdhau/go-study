package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

func main() {
	db, err := gorm.Open("mysql", "root:@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database") // Kiểm tra kết nối tới database
	}
	defer db.Close()
	index := 1
	for {
		time.Sleep(5 * time.Second)
		fmt.Println("hi, dummies man")
		db.Exec("INSERT INTO `users` (`id`, `name`) VALUES (?, ?)", index, "name-"+strconv.Itoa(index))
		index++
	}
}
