package storage

import (
	"reflect"
	"testing"

	"gitlab.com/amyasnikov/movies/common"
	"gitlab.com/amyasnikov/movies/storage"
)

var (
	url = "postgresql://postgres:postgres@127.0.0.1/test?sslmode=disable"
)

func newClearDB(url string) (*storage.DB, error) {
	db, err := storage.NewDB(url)
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
