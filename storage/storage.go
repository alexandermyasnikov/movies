package storage

import (
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"gitlab.com/amyasnikov/movies/common"
)

var (
	connection string = "XXX"
)

type DB struct {
	handle *pg.DB
}

type Movie = common.Movie

func NewDB() (*DB, error) {
	options, err := pg.ParseURL(connection)

	if err != nil {
		return nil, err
	}

	return &DB{
		handle: pg.Connect(options),
	}, nil
}

func (db *DB) Close() {
	db.handle.Close()
}

func (db *DB) CreateSchema() error {
	models := []interface{}{
		(*Movie)(nil),
	}

	for _, model := range models {
		err := db.handle.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})

		if err != nil {
			return err
		}
	}

	return nil
}

func (db *DB) Insert(movie *Movie) error {
	log.Println("insert: ", movie)

	_, err := db.handle.Model(movie).OnConflict("(id) DO UPDATE").Insert()

	return err
}

func (db *DB) Select(movie *Movie) error {
	log.Println("Select:", movie)

	return db.handle.Model(movie).WherePK().First()
}

func (db DB) Count() int {

	return 0
}

func (db *DB) Random(movie *Movie) error {
	log.Println("Random:", movie)

	return db.handle.Model(movie).OrderExpr("RANDOM()").First()
}
