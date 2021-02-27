package storage

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"gitlab.com/amyasnikov/movies/common"
)

type DB struct {
	handle *pg.DB
}

type Movie = common.Movie

func NewDB(url string) (*DB, error) {
	options, err := pg.ParseURL(url)

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

func (db *DB) DropSchema() error {
	models := []interface{}{
		(*Movie)(nil),
	}

	for _, model := range models {
		err := db.handle.Model(model).DropTable(&orm.DropTableOptions{
			IfExists: true,
		})

		if err != nil {
			return err
		}
	}

	return nil
}

func (db *DB) Insert(movie *Movie) error {
	_, err := db.handle.Model(movie).OnConflict("(id) DO UPDATE").Insert()

	return err
}

func (db *DB) Delete(movie *Movie) error {
	_, err := db.handle.Model(movie).WherePK().Delete()

	return err
}

func (db *DB) Select(movie *Movie) error {
	return db.handle.Model(movie).WherePK().First()
}

func (db DB) Count() (int, error) {
	movie := &Movie{}
	return db.handle.Model(movie).Count()
}

func (db *DB) Random(movie *Movie) error {
	return db.handle.Model(movie).OrderExpr("RANDOM()").First()
}
