package datasource

type Field [3][3]int

type Game struct {
	UUID  string
	Field Field
}
