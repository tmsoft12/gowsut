package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // SQLite sürüjisi import
)

// DB gurluşy, maglumatlar binýady bilen baglanyşygy görkezýär
var DB *sql.DB

// InitDatabase funksiýasy maglumatlar binýadyny açýar we tablisany döredýär
func InitDatabase() error {
	var err error
	// Maglumatlar binýadyny açmak
	DB, err = sql.Open("sqlite3", "./test.db")
	if err != nil {
		return err
	}

	// Maglumatlar binýady bilen baglanyşygy barlamak (Ping bilen synag etmek)
	err = DB.Ping()
	if err != nil {
		return err
	}

	// Maglumatlar binýadynda tablisa döretmek
	createTableSQL := `CREATE TABLE IF NOT EXISTS ticket (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		ticket INTEGER NOT NULL,
		type TEXT NOT NULL
	);`
	_, err = DB.Exec(createTableSQL)
	if err != nil {
		return err
	}

	// Ussatlyk bilen maglumatlar binýadyny açandygyny yglan etmek
	log.Println("Maglumatlar binýady üstünlikli açyldy we tablisa döredildi.")
	return nil
}
