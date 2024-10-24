package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // SQLite driver import
)

// DB yapısı, veritabanı bağlantısını temsil eder
var DB *sql.DB

// InitDatabase fonksiyonu veritabanı bağlantısını başlatır ve tabloyu oluşturur
func InitDatabase() error {
	var err error
	DB, err = sql.Open("sqlite3", "./test.db")
	if err != nil {
		return err
	}

	// Veritabanı tablosunu oluştur
	createTableSQL := `CREATE TABLE IF NOT EXISTS ticket (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		ticket INTEGER NOT NULL,
		type TEXT NOT NULL
	);`
	_, err = DB.Exec(createTableSQL)
	if err != nil {
		return err
	}

	return nil
}
