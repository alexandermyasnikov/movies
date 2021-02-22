package main

import (
	"log"

	"gitlab.com/amyasnikov/movies/parser"
)

func main() {
	p := parser.Parser{}
	for movie := range p.Search(100) {
		// log.Println("Movie.Id:", movie.Id)
		// log.Println("Movie.name:", movie.Name)
		log.Println("Movie:", movie.Name)
	}
}
