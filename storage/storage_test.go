package storage

import (
	"reflect"
	"testing"

	"gitlab.com/amyasnikov/movies/common"
)

var (
	url = "postgresql://postgres:postgres@127.0.0.1:5433/test?sslmode=disable"
)

func newClearDB(url string) (*DB, error) {
	db, err := NewDB(url)
	if err != nil {
		return nil, err
	}

	if db == nil {
		return nil, err
	}

	err = db.DropSchema()
	if err != nil {
		return nil, err
	}

	err = db.CreateSchema()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestInitDB(t *testing.T) {
	db, err := newClearDB(url)
	if err != nil {
		t.Error(err)
	}

	defer db.Close()
}

func TestInsertCount(t *testing.T) {
	db, err := newClearDB(url)
	if err != nil {
		t.Error(err)
	}

	defer db.Close()

	count, err := db.Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Errorf("db.Count() = %d; want 0", count)
	}

	// insert
	err = db.Insert(&common.Movie{Id: "id_1", Name: "name_1"})
	if err != nil {
		t.Error(err)
	}

	count, err = db.Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Errorf("db.Count() = %d; want 1", count)
	}

	// replace
	err = db.Insert(&common.Movie{Id: "id_1", Name: "name_1"})
	if err != nil {
		t.Error(err)
	}

	count, err = db.Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Errorf("db.Count() = %d; want 1", count)
	}

	// insert
	err = db.Insert(&common.Movie{Id: "id_2", Name: "name_2"})
	if err != nil {
		t.Error(err)
	}

	count, err = db.Count()
	if err != nil {
		t.Error(err)
	}
	if count != 2 {
		t.Errorf("db.Count() = %d; want 2", count)
	}

	// delete
	err = db.Delete(&common.Movie{Id: "id_1"})
	if err != nil {
		t.Error(err)
	}

	err = db.Delete(&common.Movie{Id: "id_2"})
	if err != nil {
		t.Error(err)
	}

	count, err = db.Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Errorf("db.Count() = %d; want 0", count)
	}
}

func TestSelect(t *testing.T) {
	db, err := newClearDB(url)
	if err != nil {
		t.Error(err)
	}

	defer db.Close()

	movie := &common.Movie{
		Id:     "id_1",
		Name:   "name_1",
		Genres: []string{"drama", "fantasy", "music"},
		Photos: []string{"photo1"},
	}

	err = db.Insert(movie)
	if err != nil {
		t.Error(err)
	}

	movieNew := &common.Movie{Id: "id_1"}
	err = db.Select(movieNew)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(movie, movieNew) {
		t.Errorf("db.Select() = %v; want %v", movieNew, movie)
	}
}

func TestSelectRandom(t *testing.T) {
	db, err := newClearDB(url)
	if err != nil {
		t.Error(err)
	}

	defer db.Close()

	movie := &common.Movie{
		Id:     "id_1",
		Name:   "name_1",
		Genres: []string{"drama", "fantasy", "music"},
		Photos: []string{"photo1"},
	}

	err = db.Insert(movie)
	if err != nil {
		t.Error(err)
	}

	movieNew := &common.Movie{Id: "id_1"}
	err = db.Random(movieNew)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(movie, movieNew) {
		t.Errorf("db.Random() = %v; want %v", movieNew, movie)
	}
}

func TestSelectRandomMany(t *testing.T) {
	db, err := newClearDB(url)
	if err != nil {
		t.Error(err)
	}

	defer db.Close()

	movie1 := &common.Movie{
		Id: "id_1",
	}

	movie2 := &common.Movie{
		Id: "id_2",
	}

	err = db.Insert(movie1)
	if err != nil {
		t.Error(err)
	}

	err = db.Insert(movie2)
	if err != nil {
		t.Error(err)
	}

	movies, err := db.RandomMany(2)
	if err != nil {
		t.Error(err)
	}

	if len(movies) != 2 {
		t.Errorf("db.RandomMany().len = %v; want 2", len(movies))
	}

	if movies[0].Id == movies[1].Id {
		t.Errorf("db.RandomMany(): movies[0] == movies[1]")
	}
}

func TestUpdate(t *testing.T) {
	db, err := newClearDB(url)
	if err != nil {
		t.Error(err)
	}

	defer db.Close()

	movie1 := &common.Movie{
		Id:     "id_1",
		Name:   "name_1",
		Genres: []string{"drama", "fantasy", "music"},
		Photos: []string{"photo1"},
	}

	err = db.Insert(movie1)
	if err != nil {
		t.Error(err)
	}

	movie := &common.Movie{Id: "id_1"}
	err = db.Select(movie)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(movie, movie1) {
		t.Errorf("db.Select() = %v; want %v", movie, movie1)
	}

	movie2 := &common.Movie{
		Id:     "id_1",
		Name:   "name_2",
		Genres: []string{"drama", "music"},
		Photos: []string{"photo1", "photo2"},
	}

	err = db.Insert(movie2)
	if err != nil {
		t.Error(err)
	}

	movie = &common.Movie{Id: "id_1"}
	err = db.Select(movie)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(movie, movie2) {
		t.Errorf("db.Select() = %v; want %v", movie, movie2)
	}
}

func TestDeleteIfNotExists(t *testing.T) {
	db, err := newClearDB(url)
	if err != nil {
		t.Error(err)
	}

	defer db.Close()

	err = db.Delete(&common.Movie{Id: "id_1"})
	if err != nil {
		t.Error(err)
	}
}

func TestSelectIfNotExists(t *testing.T) {
	db, err := newClearDB(url)
	if err != nil {
		t.Error(err)
	}

	defer db.Close()

	err = db.Select(&common.Movie{Id: "id_1"})
	if err == nil { // we expect error
		t.Error(err)
	}
}

func TestQuiz(t *testing.T) {
	db, err := newClearDB(url)
	if err != nil {
		t.Error(err)
	}

	defer db.Close()

	movie1 := &common.Movie{
		Id:      "id_1",
		Name:    "movie_1",
		Similar: []string{"id_3"},
		Photos:  []string{"photo_1"},
	}

	movie2 := &common.Movie{
		Id:      "id_2",
		Name:    "movie_2",
		Similar: []string{"id_3"},
		Photos:  []string{"photo_2"},
	}

	movie3 := &common.Movie{
		Id:     "id_3",
		Name:   "movie_3",
		Photos: []string{"photo_3"},
	}

	movie4 := &common.Movie{
		Id:     "id_4",
		Name:   "movie_4",
		Photos: []string{"photo_4"},
	}

	err = db.Insert(movie1)
	if err != nil {
		t.Error(err)
	}

	err = db.Insert(movie2)
	if err != nil {
		t.Error(err)
	}

	err = db.Insert(movie3)
	if err != nil {
		t.Error(err)
	}

	err = db.Insert(movie4)
	if err != nil {
		t.Error(err)
	}

	quiz, err := db.Quiz(2, 1)
	if err != nil {
		t.Error(err)
	}

	if quiz.Question == "" {
		t.Errorf("quiz.Question is empty")
	}

	if len(quiz.Options) != 2 {
		t.Errorf("quiz.Options.len = %d; want 2", len(quiz.Options))
	}

	if quiz.Photo == "" {
		t.Errorf("quiz.Photo is empty")
	}
}
