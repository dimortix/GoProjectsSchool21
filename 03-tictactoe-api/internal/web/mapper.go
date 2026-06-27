package web

import "tictactoe/internal/domain"

func toDomainGame(g Game) *domain.Game {
	return &domain.Game{
		UUID:  g.UUID,
		Field: domain.Field(g.Field),
	}
}

func fromDomainGame(g *domain.Game) Game {
	return Game{
		UUID:  g.UUID,
		Field: [3][3]int(g.Field),
	}
}
