package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
)

type User struct {
	Id   int `json:"id"`
	Name string `json:"name"`
}

func main() {
	db, err := sql.Open("mysql", "root:password@/test_db")
	if err != nil {
		panic("fail to open")
	}

	rows, err := db.Query("SELECT * FROM test_table")
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
	}

	users := []User{}
	user := User{}
	for rows.Next() {
		users = append(users, user)
	}
	jsonBytes, err := json.Marshal(users)
	fmt.Println(string(jsonBytes))
}