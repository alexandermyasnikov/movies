package common

type Movie struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Genres  []string `json:"genres"`
	Similar []string `json:"similar"`
	Photos  []string `json:"photos"`
}

type Quiz struct {
	Question  string   `json:"question"`
	Photo     string   `json:"photo"`
	Options   []string `json:"options"`
	CorrectId int      `json:"correct_option_id"`
}

type StorageCountAPI struct {
	Count int `json:"count"`
}

type StorageMovieAPI = Movie
