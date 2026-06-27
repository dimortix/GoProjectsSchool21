package datasource

import "tictactoe/internal/domain"

func toDomainGame(g *Game) *domain.Game {
	if g == nil {
		return nil
	}
	var f domain.Field
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			f[i][j] = g.Field[i][j]
		}
	}
	return &domain.Game{
		UUID:  g.UUID,
		Field: f,
	}
}

func toDsGame(g *domain.Game) *Game {
	if g == nil {
		return nil
	}
	var f Field
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			f[i][j] = g.Field[i][j]
		}
	}
	return &Game{
		UUID:  g.UUID,
		Field: f,
	}
}
