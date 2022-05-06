package duelyst

type User struct {
	Id       string
	Username string
}

type Player struct {
	User
	general *General
	hand    *Hand
	deck    *Deck
	exile   []*Card
}

func NewPlayer(user *User, hand *Hand, deck *Deck) *Player {
	return &Player{
		User: *user,
		hand: hand,
		deck: deck,
	}
}

func (player *Player) IsAlive() bool {
	return player.general.IsAlive()
}
