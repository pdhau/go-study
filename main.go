package main

import (
	"fmt"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	db, err := gorm.Open("mysql", "root:@(localhost)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	i := 1
	for i <= 20 {
		go saveMulti(i, db)
		i++
	}
	for {
	}
}

func saveMulti(index int, db *gorm.DB) {
	db, err := gorm.Open("mysql", "root:@(localhost)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	fmt.Println(fmt.Sprintf("%d is running", index))

	for i := 0; i < 50; i++ {
		fmt.Println("Start ", i)
		tx := db.Begin()
		for index := 0; index < 30; index++ {
			err := save(tx, index)

			if err != nil {
				fmt.Println("err ==============", err)
				tx.Rollback()
			}
		}
		tx.Commit()
		fmt.Println("Commit ", i)
		time.Sleep(1 * time.Second)
	}
}

func save(db *gorm.DB, index int) error {
	sql := "INSERT INTO todo SET time=?, name=?, description=? ON DUPLICATE KEY UPDATE description = ?"

	return db.Exec(sql, "time "+strconv.Itoa(index), "dummy 1", "description", time.Now().UnixNano()).Error
}
