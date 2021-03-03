package storage

import (
	"fmt"
	"math/rand"
	"sort"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"gitlab.com/amyasnikov/movies/common"
)

type DB struct {
	handle *pg.DB
}

type (
	Movie = common.Movie
	Quiz  = common.Quiz
)

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

func (db *DB) RandomMany(count int) ([]Movie, error) {
	var movies []Movie
	err := db.handle.Model(&movies).OrderExpr("RANDOM()").Limit(count).Select()
	return movies, err
}

func (db *DB) Quiz(optionsCount, SimilarCount int) (Quiz, error) {
	var quiz Quiz
	var movie Movie

	err := db.Random(&movie)
	if err != nil {
		return quiz, err
	}

	if len(movie.Name) == 0 {
		return quiz, fmt.Errorf("movie %s has no name", movie.Id)
	}

	if len(movie.Photos) == 0 {
		return quiz, fmt.Errorf("movie %s(%s) has no photos", movie.Name, movie.Id)
	}

	quiz.Question = "What is name of a movie?"
	quiz.Photo = movie.Photos[rand.Intn(len(movie.Photos))]
	quiz.Options = make([]string, 0, optionsCount)
	quiz.Options = append(quiz.Options, movie.Name)

	if len(movie.Similar) > 0 {
		for i := 0; i < SimilarCount; i++ {
			id := movie.Similar[rand.Intn(len(movie.Similar))]
			m := Movie{Id: id}
			err := db.Select(&m)
			if err == nil && m.Name != "" && !common.ContainsString(quiz.Options, m.Name) {
				quiz.Options = append(quiz.Options, m.Name)
			}
		}
	}

	for i := 0; i < optionsCount; i++ {
		if len(quiz.Options) == optionsCount {
			break
		}

		var m Movie
		err := db.Random(&m)
		if err == nil && m.Name != "" && !common.ContainsString(quiz.Options, m.Name) {
			quiz.Options = append(quiz.Options, m.Name)
		}
	}

	if len(quiz.Options) != optionsCount {
		return quiz, fmt.Errorf("can not find movie")
	}

	sort.Strings(quiz.Options)
	quiz.CorrectId = sort.SearchStrings(quiz.Options, movie.Name)

	if quiz.CorrectId >= len(quiz.Options) {
		return quiz, fmt.Errorf("can not search correctId")
	}

	return quiz, nil
}
