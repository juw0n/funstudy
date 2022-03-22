package data

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// var err error

func OpenDatabase() error {
	var err error
	db, err = sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		return err
	}
	return db.Ping()
}
func CreateTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS funstudy (
		"idNote" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"word" TEXT,
		"definition" TEXT,
		"category" TEXT
	  );`
	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Println(err.Error())
	}
	statement.Exec()
	log.Println("funstudy table created successfully")
}

func InsertNote(word string, definition string, category string) {
	insertNoteSQL := `INSERT INTO funstudy(word, definition, category) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertNoteSQL)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = statement.Exec(word, definition, category)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Inserted new note successfully")
}

func DisplayAllNotes() {
	row, err := db.Query("SELECT * FROM funstudy ORDER BY word")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() {
		var idNote int
		var word string
		var definition string
		var category string
		row.Scan(&idNote, &word, &definition, &category)
		log.Println("[", category, "] ", word, "—", definition)
	}
}
