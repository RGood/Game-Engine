package duelyst

import "github.com/RGood/game-engine/pkg/engine"

type Game struct {
	board        *Board
	players      []*Player
	engine       *engine.Engine
	activePlayer int
}

func NewGame(board *Board, players []*Player) *Game {
	return &Game{
		board,
		players,
		engine.NewEngine(),
		0,
	}
}

func (game *Game) GetActivePlayer() *Player {
	return game.players[game.activePlayer]
}

func (game *Game) RotateActive() *Player {
	startIndex := game.activePlayer
	curIndex := (startIndex + 1) % len(game.players)

	for !game.players[curIndex].IsAlive() && curIndex != startIndex {
		curIndex++
		curIndex %= len(game.players)
	}

	game.activePlayer = curIndex

	return game.GetActivePlayer()
}

func (game *Game) HasEnded() (bool, *Player) {
	livePlayers := []*Player{}
	for _, player := range game.players {
		if player.IsAlive() {
			livePlayers = append(livePlayers, player)
		}
	}

	if len(livePlayers) == 1 {
		return true, livePlayers[0]
	} else if len(livePlayers) == 0 {
		return true, nil
	} else {
		return false, nil
	}
}
