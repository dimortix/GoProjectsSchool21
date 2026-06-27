package domain

type Service interface {
	NextMove(game *Game) error
	ValidateField(oldGame, newGame *Game) error
	CheckEndGame(game *Game) (bool, int)
}
