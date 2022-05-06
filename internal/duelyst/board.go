package duelyst

import (
	"github.com/RGood/go-collection-functions/pkg/set"
)

type Board struct {
	units      *set.OrderedSet[*Unit]
	tiles      []Tile
	maxX, maxY int
}

type Position struct {
	X, Y int
}

func (pos Position) Equals(otherPos Position) bool {
	return otherPos.X == pos.X && otherPos.Y == pos.Y
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func (pos Position) NextTo(otherPos Position) bool {
	return max(abs(otherPos.X-pos.X), abs(otherPos.Y-pos.Y)) == 1
}

func NewBoard(maxX, maxY int) *Board {
	return &Board{
		set.NewOrderedSet[*Unit](),
		[]Tile{},
		maxX,
		maxY,
	}
}

func (board *Board) At(pos Position) (Unit, Tile) {
	var u Unit
	board.units.ForEach(func(_ int, unit *Unit) {
		if (*unit).GetPosition().Equals(pos) {
			u = *unit
		}
	})

	var t Tile
	for _, tile := range board.tiles {
		if tile.Position().Equals(pos) {
			t = tile
			break
		}
	}

	return u, t
}

func (board *Board) Place(unit Unit, pos Position) {
	unit.SetPosition(&pos)
	board.units.Add(unit.Reference())
}

func (board *Board) Remove(unit Unit) {
	unit.SetPosition(nil)
	board.units.Remove(unit.Reference())
}
