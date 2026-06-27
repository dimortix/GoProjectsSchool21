package datasource

import (
	"errors"
	"math"
	"tictactoe/internal/domain"
)

type serviceImpl struct {
	repo Repository
}

func NewService(repo Repository) domain.Service {
	return &serviceImpl{repo: repo}
}

func (s *serviceImpl) CheckEndGame(game *domain.Game) (bool, int) {
	winner := checkWinner(game.Field)
	if winner != 0 {
		return true, winner
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if game.Field[i][j] == 0 {
				return false, 0
			}
		}
	}
	return true, 0
}

func (s *serviceImpl) ValidateField(oldGame, newGame *domain.Game) error {
	ended, _ := s.CheckEndGame(oldGame)
	if ended {
		return errors.New("cannot make a move in an already ended game")
	}

	diffCount := 0
	var userMove int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if oldGame.Field[i][j] != newGame.Field[i][j] {
				if oldGame.Field[i][j] != 0 {
					return errors.New("cannot overwrite an existing move")
				}
				diffCount++
				userMove = newGame.Field[i][j]
			}
		}
	}

	if diffCount != 1 {
		return errors.New("strictly one move is allowed")
	}

	xCount, oCount := countPieces(oldGame.Field)
	expectedMove := 1
	if xCount > oCount {
		expectedMove = -1
	}

	if userMove != expectedMove {
		return errors.New("invalid player turn")
	}

	return nil
}

func (s *serviceImpl) NextMove(game *domain.Game) error {
	ended, _ := s.CheckEndGame(game)
	if ended {
		return nil
	}

	xCount, oCount := countPieces(game.Field)
	computerPlayer := 1
	if xCount > oCount {
		computerPlayer = -1
	}

	bestScore := math.Inf(-1)
	if computerPlayer == -1 {
		bestScore = math.Inf(1)
	}

	bestMove := [2]int{-1, -1}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if game.Field[i][j] == 0 {
				game.Field[i][j] = computerPlayer

				nextIsMax := true
				if computerPlayer == 1 {
					nextIsMax = false
				}
				score := minimax(game.Field, 0, nextIsMax, computerPlayer)
				game.Field[i][j] = 0

				if computerPlayer == 1 {
					if score > bestScore {
						bestScore = score
						bestMove = [2]int{i, j}
					}
				} else {
					if score < bestScore {
						bestScore = score
						bestMove = [2]int{i, j}
					}
				}
			}
		}
	}

	if bestMove[0] != -1 && bestMove[1] != -1 {
		game.Field[bestMove[0]][bestMove[1]] = computerPlayer
	}

	return nil
}

func minimax(board domain.Field, depth int, isMaximizing bool, computerPlayer int) float64 {
	winner := checkWinner(board)
	if winner == 1 {
		return 10 - float64(depth)
	} else if winner == -1 {
		return -10 + float64(depth)
	}

	isDraw := true
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == 0 {
				isDraw = false
			}
		}
	}
	if isDraw {
		return 0
	}

	if isMaximizing {
		bestScore := math.Inf(-1)
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board[i][j] == 0 {
					board[i][j] = 1
					score := minimax(board, depth+1, false, computerPlayer)
					board[i][j] = 0
					if score > bestScore {
						bestScore = score
					}
				}
			}
		}
		return bestScore
	} else {
		bestScore := math.Inf(1)
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board[i][j] == 0 {
					board[i][j] = -1
					score := minimax(board, depth+1, true, computerPlayer)
					board[i][j] = 0
					if score < bestScore {
						bestScore = score
					}
				}
			}
		}
		return bestScore
	}
}

func checkWinner(b domain.Field) int {
	for i := 0; i < 3; i++ {
		if b[i][0] != 0 && b[i][0] == b[i][1] && b[i][1] == b[i][2] {
			return b[i][0]
		}
		if b[0][i] != 0 && b[0][i] == b[1][i] && b[1][i] == b[2][i] {
			return b[0][i]
		}
	}
	if b[0][0] != 0 && b[0][0] == b[1][1] && b[1][1] == b[2][2] {
		return b[0][0]
	}
	if b[0][2] != 0 && b[0][2] == b[1][1] && b[1][1] == b[2][0] {
		return b[0][2]
	}
	return 0
}

func countPieces(b domain.Field) (xCount, oCount int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b[i][j] == 1 {
				xCount++
			} else if b[i][j] == -1 {
				oCount++
			}
		}
	}
	return
}
