package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"ronorm/ronorm"
)

func main() {
	//sqlite3Demo()
	//sessionDemo()
	transactionDemo()
}

func sqlite3Demo() {
	db, _ := sql.Open("sqlite3", "ron.db")
	defer func() { _ = db.Close() }()
	_, _ = db.Exec("DROP TABLE IF EXISTS User;")
	_, _ = db.Exec("CREATE TABLE User(name text);")
	result, err := db.Exec("INSERT INTO User(`name`) values (?), (?)", "Tom", "Sam")
	if err == nil {
		affected, _ := result.RowsAffected()
		log.Println(affected)
	}
	row := db.QueryRow("SELECT name FROM User LIMIT 1")
	var name string
	if err := row.Scan(&name); err == nil {
		log.Println(name)
	}
}

func sessionDemo() {
	engine, _ := ronorm.NewEngine("sqlite3", "ron.db")
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(name text);").Exec()
	_, _ = s.Raw("CREATE TABLE User(name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`name`) values (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)
}

func transactionDemo() {
	db, _ := sql.Open("sqlite3", "ron.db")
	defer func() { _ = db.Close() }()
	_, _ = db.Exec("CREATE TABLE IF NOT EXISTS User(`name` text);")

	tx, _ := db.Begin()
	_, err1 := tx.Exec("INSERT INTO User(`name`) VALUES (?)", "Tom")
	_, err2 := tx.Exec("INSERT INTO User(`name`) VALUES (?)", "Jack")
	if err1 != nil || err2 != nil {
		_ = tx.Rollback()
		log.Println("Rollback", err1, err2)
	} else {
		_ = tx.Commit()
		log.Println("Commit")
	}
}
