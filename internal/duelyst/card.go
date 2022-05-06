package duelyst

type Target interface {
	GetTargets() []*Position
}

type UnitList []Unit

type TargetFunc[T any] func([]T) []T

func (targetFunc TargetFunc[T]) And(next TargetFunc[T]) TargetFunc[T] {
	return func(targetList []T) []T {
		return next(targetFunc(targetList))
	}
}

func (unitList UnitList) GetTargets() []Position {
	positions := []Position{}

	for _, unit := range unitList {
		positions = append(positions, *unit.GetPosition())
	}

	return positions
}

func Friendly(players []*Player) TargetFunc[Unit] {
	return func(unitList []Unit) []Unit {
		var newList UnitList = []Unit{}

		for _, unit := range unitList {
			if unit.Friendly(players) {
				newList = append(newList, unit)
			}
		}

		return newList
	}
}

func OfType(t string) TargetFunc[Unit] {
	return func(unitList []Unit) []Unit {
		var newList UnitList = []Unit{}

		for _, unit := range unitList {
			if unit.Type() == t {
				newList = append(newList, unit)
			}
		}

		return newList
	}
}

func OfClass(c string) TargetFunc[Unit] {
	return func(unitList []Unit) []Unit {
		var newList UnitList = []Unit{}

		for _, unit := range unitList {
			for _, class := range unit.Classes() {
				if class == c {
					newList = append(newList, unit)
					break
				}
			}
		}

		return newList
	}
}

func Minions() TargetFunc[Unit] {
	return OfType("minion")
}

func Generals() TargetFunc[Unit] {
	return OfType("general")
}

type PositionList []Position

func (positionList PositionList) GetTargets() []Position {
	return positionList
}

func Nearby(target Target) TargetFunc[Position] {
	return func(oldPositions []Position) []Position {
		newPositions := []Position{}

		for _, pos := range target.GetTargets() {
			for _, oldPos := range oldPositions {
				if pos.NextTo(oldPos) {
					newPositions = append(newPositions, *pos)
					break
				}
			}
		}

		return newPositions
	}

}

func Unoccupied(game *Game) TargetFunc[Position] {
	return func(oldPositions []Position) []Position {
		newPositions := []Position{}

		for _, oldPos := range oldPositions {
			occupied := false
			game.board.units.ForEach(func(_ int, unit *Unit) {
				if (*unit).GetPosition().Equals(oldPos) {
					occupied = true
				}
			})

			if !occupied {
				newPositions = append(newPositions, oldPos)
			}
		}

		return newPositions
	}
}

type Card struct {
	targetFunc func(game *Game) Target
}

func (card *Card) GetValidTargets(game *Game) []*Position {
	return card.targetFunc(game).GetTargets()
}
