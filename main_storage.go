package main

import (
	"context"
	"log"
	"net"

	mg "gitlab.com/amyasnikov/movies/grpc"
	"gitlab.com/amyasnikov/movies/storage"
	"google.golang.org/grpc"
)

type MoviesStorage struct {
	db *storage.DB
}

var (
	host   = ":8080"
	MOVIES = "movies"
)

func (s MoviesStorage) GetStats(ctx context.Context, r *mg.Void) (*mg.Stats, error) {
	ret := &mg.Stats{
		MoviesCount: int32(s.db.Count(MOVIES)),
	}
	return ret, nil
}

func (s MoviesStorage) GetMovie(ctx context.Context, r *mg.MovieKey) (*mg.Movie, error) {
	response := &mg.Movie{
		Json: s.db.SelectKV(MOVIES, r.Id),
	}
	return response, nil
}

func (s MoviesStorage) GetMovieRandom(ctx context.Context, r *mg.Void) (*mg.Movie, error) {
	_, v := s.db.RandomKV(MOVIES)
	response := &mg.Movie{
		Json: v,
	}
	return response, nil
}

func (s MoviesStorage) UpdateMovie(ctx context.Context, r *mg.Movie) (*mg.Void, error) {
	s.db.InsertKV(MOVIES, r.Id, r.Json)
	response := &mg.Void{}
	return response, nil
}

func main() {
	listen, err := net.Listen("tcp", host)
	if err != nil {
		log.Panic(err)
	}

	server := grpc.NewServer()
	moviesStorage := MoviesStorage{
		db: storage.NewDB("file:/tmp/imdb.db?cache=shared"),
	}

	moviesStorage.db.CreateTableKV(MOVIES)

	mg.RegisterMoviesStorageServer(server, moviesStorage)
	server.Serve(listen)
}
