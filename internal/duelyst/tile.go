package duelyst

type Tile interface {
	Position() *Position
}

type NormalTile struct {
	pos *Position
}

func (normalTile *NormalTile) Position() *Position {
	return normalTile.pos
}
