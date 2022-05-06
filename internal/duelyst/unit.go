package duelyst

type Unit interface {
	GetPosition() *Position
	SetPosition(*Position)
	Damage(int)
	Reference() *Unit
	Friendly([]*Player) bool
	Type() string
	Classes() []string
}

type Minion struct {
	Name     string
	Attack   int
	Health   int
	damage   int
	owner    string
	classes  []string
	position *Position
	ref      *Unit
}

func (minion Minion) GetPosition() *Position {
	return minion.position
}

func (minion *Minion) SetPosition(pos *Position) {
	minion.position = pos
}

func (minion Minion) IsAlive() bool {
	return minion.position != nil
}

func (minion *Minion) Damage(value int) {
	minion.damage += value
}

func (minion *Minion) Friendly(players []*Player) bool {
	for _, player := range players {
		if player.Username == minion.owner {
			return true
		}
	}

	return false
}

func (minion Minion) Type() string {
	return "minion"
}

func (minion Minion) Classes() []string {
	return minion.classes
}

func (minion Minion) Reference() *Unit {
	if minion.ref == nil {
		var unit Unit = &minion
		minion.ref = &unit
	}

	return minion.ref
}

type General struct {
	Minion
	Artifacts []*Artifact
}

func (general General) Type() string {
	return "general"
}
