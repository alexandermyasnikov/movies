package main

import (
	"log"

	"gitlab.com/amyasnikov/movies/parser"
)

func main() {
	p := parser.Parser{}
	for movie := range p.Movies(500) {
		log.Println("movie:", movie.Id, movie.Name, len(movie.Photos))
	}
}
