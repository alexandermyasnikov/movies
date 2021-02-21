package storage

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var (
	MOVIES = "movies"
)

type DB struct {
	Handle *sql.DB
}

func NewDB(connectionString string) DB {
	handle, err := sql.Open("sqlite3", connectionString)

	if err != nil {
		panic(err)
	}

	return DB{
		Handle: handle,
	}
}

func (db DB) Close() {
	db.Handle.Close()
}

func (db DB) CreateTableKV(name string) {
	sql := `CREATE TABLE IF NOT EXISTS ` + name + ` (
    key TEXT NOT NULL PRIMARY KEY,
		value TEXT
	  );`

	stmt, err := db.Handle.Prepare(sql)

	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = stmt.Exec()

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Create table", name)
}

func (db DB) InsertKV(name, key, value string) {
	sql := `INSERT OR REPLACE INTO ` + name + ` (key, value) VALUES (?, ?);`

	stmt, err := db.Handle.Prepare(sql)

	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = stmt.Exec(key, value)

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Insert", name, key, value)
}

func (db DB) SelectKV(name, key string) (value string) {
	sql := `SELECT value from ` + name + ` WHERE key = ? LIMIT 1;`

	stmt, err := db.Handle.Prepare(sql)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer stmt.Close()

	rows, err := stmt.Query(key)

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&value)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("value:", value)
	}

	log.Println("Select", name, key)
	return
}

func (db DB) Init() {
	db.CreateTableKV(MOVIES)
	db.InsertKV(MOVIES, "tt123456", "new_json")
	db.InsertKV(MOVIES, "tt123457", "text")
	db.SelectKV(MOVIES, "tt123457")
	db.SelectKV(MOVIES, "tt123458")
}
