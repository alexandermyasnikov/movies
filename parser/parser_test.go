package parser

import (
	"testing"
)

func TestIdsCount(t *testing.T) {
	p := NewParser(DefaultOptions)

	count := 0
	for range p.Ids(50) {
		count++
	}

	if count != 50 {
		t.Errorf("db.Ids() = %d; want 50", count)
	}
}

func TestMovie(t *testing.T) {
	p := NewParser(DefaultOptions)

	movie := p.Movie("tt0111161")

	if movie == nil {
		t.Errorf("p.Movie() = nil")
	}

	id := "tt0111161"
	if movie.Id != id {
		t.Errorf("p.Movie().Id = %v; want %v", movie.Id, id)
	}

	name := "The Shawshank Redemption"
	if movie.Name != name {
		t.Errorf("p.Movie().Name = %v; want %v", movie.Name, name)
	}

	if len(movie.Genres) != 1 || movie.Genres[0] != "Drama" {
		t.Errorf("p.Movie().Genres = %v; want %v", movie.Genres, []string{"Drama"})
	}

	if len(movie.Similar) != 12 {
		t.Errorf("p.Movie().Similar = %v; want %v", len(movie.Similar), 12)
	}

	if len(movie.Photos) < 200 {
		t.Errorf("p.Movie().Photos = %v", len(movie.Photos))
	}

	url := "https://m.media-amazon.com/images/M/MV5BMTM0NjUxMDk5MF5BMl5BanBnXkFtZTcwNDMxNDY3Mw@@._V1_UY100_CR25,0,100,100_AL_"
	if movie.Photos[0] != url {
		t.Errorf("p.Movie().Photos[0] = %v; want %v", movie.Photos[0], url)
	}
}

func TestMovieRu(t *testing.T) {
	opts := DefaultOptions
	opts.Lang = "ru"
	p := NewParser(opts)

	movie := p.Movie("tt0111161")

	if movie == nil {
		t.Errorf("p.Movie() = nil")
	}

	id := "tt0111161"
	if movie.Id != id {
		t.Errorf("p.Movie().Id = %v; want %v", movie.Id, id)
	}

	name := "Побег из Шоушенка"
	if movie.Name != name {
		t.Errorf("p.Movie().Name = %v; want %v", movie.Name, name)
	}
}
