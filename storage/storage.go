package storage

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	Handle *sql.DB
}

func NewDB(connectionString string) *DB {
	handle, err := sql.Open("sqlite3", connectionString)

	if err != nil {
		panic(err)
	}

	db := &DB{
		Handle: handle,
	}

	return db
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
		log.Println(err)
		return
	}

	_, err = stmt.Exec()

	if err != nil {
		log.Println(err)
		return
	}
}

func (db DB) InsertKV(name, key, value string) {
	log.Println("insert: ", name, key, value)

	sql := `INSERT OR REPLACE INTO ` + name + ` (key, value) VALUES (?, ?);`

	stmt, err := db.Handle.Prepare(sql)

	if err != nil {
		log.Println(err)
		return
	}

	_, err = stmt.Exec(key, value)

	if err != nil {
		log.Println(err)
		return
	}
}

func (db DB) SelectKV(name, key string) (value string) {
	log.Println("Select:", name, key)

	sql := `SELECT value from ` + name + ` WHERE key = ? LIMIT 1;`

	stmt, err := db.Handle.Prepare(sql)

	if err != nil {
		log.Println(err)
		return
	}

	defer stmt.Close()

	rows, err := stmt.Query(key)

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&value)
		if err != nil {
			log.Println(err)
		}
	}

	return
}

func (db DB) Count(name string) (count int) {
	sql := `SELECT count(*) from ` + name + ` LIMIT 1;`

	stmt, err := db.Handle.Prepare(sql)

	if err != nil {
		log.Println(err)
		return
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		log.Println(err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			log.Println(err)
		}
	}

	return
}

func (db DB) RandomKV(name string) (key, value string) {
	log.Println("Random:", name, key)

	sql := `SELECT key, value from ` + name + ` ORDER BY RANDOM() LIMIT 1;`

	stmt, err := db.Handle.Prepare(sql)

	if err != nil {
		log.Println(err)
		return
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&key, value)
		if err != nil {
			log.Println(err)
		}
	}

	return
}
