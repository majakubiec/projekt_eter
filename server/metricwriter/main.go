package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	var conninfo string = "postgres://postgres@postgres:5432/postgres?sslmode=disable"

	db, err := sql.Open("postgres", conninfo)
	if err != nil {
		panic(err)
	}

	for {
		value := (rand.Int() % 20) + 100
		// insert random data to db
		r, err := db.Exec("INSERT INTO datas (time, measure) VALUES (now(), $1);", value)
		if err != nil {
			panic(err)
		}
		fmt.Println(r)
		time.Sleep(time.Second)
	}

}
