package main

import (
	"context"
	"log"
	"net"

	"gitlab.com/amyasnikov/movies/common"
	mg "gitlab.com/amyasnikov/movies/grpc"
	"gitlab.com/amyasnikov/movies/storage"
	"google.golang.org/grpc"
)

type MoviesStorage struct {
	db *storage.DB
}

var (
	host = ":8080"
)

func (s MoviesStorage) GetStats(ctx context.Context, r *mg.Void) (*mg.Stats, error) {
	ret := &mg.Stats{
		MoviesCount: int32(s.db.Count()),
	}
	return ret, nil
}

func (s MoviesStorage) GetMovie(ctx context.Context, r *mg.MovieKey) (*mg.Movie, error) {
	movie := &common.Movie{
		Id: r.Id,
	}

	err := s.db.Select(movie)

	return &mg.Movie{
		Id:      movie.Id,
		Name:    movie.Name,
		Genres:  movie.Genres,
		Similar: movie.Similar,
		Photos:  movie.Photos,
	}, err
}

func (s MoviesStorage) GetMovieRandom(ctx context.Context, r *mg.Void) (*mg.Movie, error) {
	// TODO
	// _, v := s.db.RandomKV(MOVIES)
	// response := &mg.Movie{
	// 	Json: v,
	// }
	// return response, nil
	return nil, nil
}

func (s MoviesStorage) UpdateMovie(ctx context.Context, r *mg.Movie) (*mg.Void, error) {
	// TODO
	// s.db.InsertKV(MOVIES, r.Id, r.Json)
	// response := &mg.Void{}
	// return response, nil
	return nil, nil
}

func main() {
	listen, err := net.Listen("tcp", host)
	if err != nil {
		log.Panic(err)
	}

	db, err := storage.NewDB()
	if err != nil {
		log.Panic(err)
	}

	server := grpc.NewServer()
	moviesStorage := MoviesStorage{
		db: db,
	}

	moviesStorage.db.CreateSchema()

	mg.RegisterMoviesStorageServer(server, moviesStorage)
	server.Serve(listen)
}
