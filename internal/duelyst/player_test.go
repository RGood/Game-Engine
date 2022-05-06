package duelyst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayerLiveness(t *testing.T) {
	p := &Player{
		User: User{
			Id:       "1234",
			Username: "Foo",
		},
		general: &General{
			Minion: Minion{
				Name:     "Bar",
				Attack:   2,
				Health:   25,
				damage:   0,
				position: nil,
			},
		},
	}

	b := NewBoard(9, 5)

	assert.False(t, p.IsAlive())

	b.Place(p.general, Position{0, 2})

	assert.True(t, p.IsAlive())
}
