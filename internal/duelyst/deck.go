package duelyst

import (
	"math/rand"
	"time"
)

type Deck struct {
	General General
	cards   []*Card
}

func (deck *Deck) Draw() *Card {
	if len(deck.cards) > 0 {
		card := deck.cards[0]
		deck.cards = deck.cards[1:]
		return card
	}

	return nil
}

// Put a card in and get a different card out
// - If the deck is empty, return the same card
func (deck *Deck) Replace(card *Card) *Card {
	newCard := deck.Draw()
	if newCard != nil {
		deck.cards = append(deck.cards, card)
		deck.Shuffle()
		return newCard
	}
	return card
}

// Re-order the deck randomly
func (deck *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i := len(deck.cards) - 1; i > 0; i-- { // Fisher–Yates shuffle
		j := rand.Intn(i + 1)
		deck.cards[i], deck.cards[j] = deck.cards[j], deck.cards[i]
	}
}

type Hand struct {
	MaxSize int
	Cards   []*Card
}

func NewHand(size int) *Hand {
	cards := make([]*Card, size)
	return &Hand{
		size,
		cards,
	}
}

func (hand *Hand) FindEmptySlot() (int, bool) {
	for index, slot := range hand.Cards {
		if slot == nil {
			return index, true
		}
	}

	return -1, false
}
