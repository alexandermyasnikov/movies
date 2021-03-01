package common

type Movie struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Genres  []string `json:"genres"`
	Similar []string `json:"similar"`
	Photos  []string `json:"photos"`
}

type StorageCountAPI struct {
	Count int `json:"count"`
}

type StorageMovieAPI = Movie
