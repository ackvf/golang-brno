package db

import (
	"database/sql"
	_ "github.com/lib/pq" // _ means we don't use this package directly, but we need to import it for other packages
	// we need to install this package by `go get github.com/lib/pq`
)

// github dstankovic, https://gist.github.com/dstankovic / connection_string

var db *sql.DB

func getDB() *sql.DB {
	if db == nil {
		var err error
		db, err := sql.Open("postgres" /*dataSourceName*/, "host=ec2-23-23-211-21.compute-1.amazonaws.com database=ddrkheg4gaufom user=yxdgjzsqajkoia password=pJoowuHjfQ_02pzTtbBieQq801")
		if err != nil {
			panic(err)
		}
	}
	return db
}
