package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEngineProcessing(t *testing.T) {
	e := NewEngine()

	listener := &mockListener{}
	e.Listen(listener.Reference())

	interceptor := &mockInterceptor{}
	e.Intercept(interceptor.Reference())

	e.Queue(&NoopEvent{})
	e.Run()

	assert.Equal(t, 2, len(listener.events))

	_, ok := listener.events[0].(*NoopEvent)
	assert.True(t, ok)

	te, ok := listener.events[1].(*TickEvent)
	assert.True(t, ok)
	assert.Equal(t, 1, te.Iterations)

	assert.Equal(t, 1, len(interceptor.events))

	_, ok = interceptor.events[0].(*NoopEvent)
	assert.True(t, ok)
}

func TestEngineUnsubscribe(t *testing.T) {

	e := NewEngine()

	listener := &mockListener{}
	e.Listen(listener.Reference())

	interceptor := &mockInterceptor{}
	e.Intercept(interceptor.Reference())

	e.Queue(&NoopEvent{})
	e.Run()

	assert.Equal(t, 2, len(listener.events))

	_, ok := listener.events[0].(*NoopEvent)
	assert.True(t, ok)

	te, ok := listener.events[1].(*TickEvent)
	assert.True(t, ok)
	assert.Equal(t, 1, te.Iterations)

	assert.Equal(t, 1, len(interceptor.events))

	_, ok = interceptor.events[0].(*NoopEvent)
	assert.True(t, ok)

	e.Unlisten(listener.Reference())
	e.Unintercept(interceptor.Reference())

	e.Queue(&NoopEvent{})
	e.Run()

	assert.Equal(t, 1, len(interceptor.events))
	assert.Equal(t, 2, len(listener.events))
}

func TestCycleCount(t *testing.T) {
	e := NewEngine()

	listener := &mockListener{}
	e.Listen(listener.Reference())

	interceptor := &mockInterceptor{}
	e.Intercept(interceptor.Reference())

	re := &RepeaterEvent{
		5,
		func(self Event) {
			e.Queue(self)
		},
	}

	e.Queue(re)
	e.RunCycles(3)

	assert.Equal(t, 6, len(listener.events))
	assert.Equal(t, 3, len(interceptor.events))

	e.Run()

	assert.Equal(t, 10, len(listener.events))
	assert.Equal(t, 5, len(interceptor.events))
}
