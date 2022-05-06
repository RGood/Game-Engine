package duelyst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinionDamage(t *testing.T) {
	m := &Minion{
		Name:     "Foo",
		Attack:   1,
		Health:   5,
		damage:   0,
		position: &Position{0, 0},
	}

	m.Damage(2)

	assert.Equal(t, 2, m.damage)
}

func TestGeneralDamage(t *testing.T) {
	g := &General{
		Minion: Minion{
			Name:     "Foo",
			Attack:   2,
			Health:   25,
			damage:   0,
			position: &Position{0, 0},
		},
		Artifacts: []*Artifact{},
	}

	g.Damage(10)

	assert.Equal(t, 10, g.damage)
}
