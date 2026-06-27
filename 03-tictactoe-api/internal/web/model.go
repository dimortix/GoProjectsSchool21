package web

type Game struct {
	UUID  string    `json:"UUID"`
	Field [3][3]int `json:"Field"`
}
