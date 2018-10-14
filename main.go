package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"strings"
	"time"
)

func dbConn() *sql.DB {
	const (
		HOST     = "localhost"
		PORT     = 5432
		USER     = "postgres"
		PASSWORD = "1303"
		DBNAME   = "portfolio"
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		HOST, PORT, USER, PASSWORD, DBNAME)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func getNow() string {
	splitTime := strings.Split(time.Now().String(), " ")
	now := splitTime[0]
	return now
}

//func dbConn() *sql.DB {
//	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
//	if err != nil {
//		log.Fatalf("Error opening database: %q", err)
//	}
//	return db
//}

func main(){
	defer DB.Close()
	Routes()
}
