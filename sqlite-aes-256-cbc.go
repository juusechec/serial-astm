package main

import (
	"database/sql"
	"fmt"
	_ "github.com/xeodou/go-sqlcipher"
)

func main() {
	db, err := sql.Open("sqlite3", "./encrip.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	p := "PRAGMA key = 'Ab123456';"
	_, err = db.Exec(p)
	if err != nil {
		fmt.Println(err)
	}

	rows, err := db.Query("select name from foo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}
