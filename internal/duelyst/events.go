package duelyst

import "github.com/RGood/game-engine/pkg/engine"

type DrawCard struct {
	player *Player
}

func (draw *DrawCard) Execute(engine *engine.Engine) {
	// Draw a card from the deck
	card := draw.player.deck.Draw()

	// If a card was drawn successfully
	if card != nil {
		// Find an empty spot in the hand
		slotIndex, slotFound := draw.player.hand.FindEmptySlot()

		// If an empty spot was found
		if slotFound {
			// Put the card there
			draw.player.hand.Cards[slotIndex] = card
		} else {
			// Otherwise exile the card
			draw.player.exile = append(draw.player.exile, card)
		}
	} else {

		// If a card couldn't be drawn, deal 2 damage to that player's general
		engine.Queue(&Damage{
			unit:  draw.player.general,
			value: 2,
		})
	}
}

type Damage struct {
	unit  Unit
	value int
}

func (damage *Damage) Execute(engine *engine.Engine) {
	damage.unit.Damage(damage.value)
}

type Kill struct {
	unit  Unit
	board *Board
}

func (kill *Kill) Execute(engine *engine.Engine) {
	kill.board.Remove(kill.unit)
}

type EndTurn struct {
	game *Game
}

func (endTurn *EndTurn) Execute(engine *engine.Engine) {
	endTurn.game.RotateActive()
}

type Move struct {
	unit Unit
	pos  *Position
}

func (move *Move) Execute(engine *engine.Engine) {
	move.unit.SetPosition(move.pos)
}
