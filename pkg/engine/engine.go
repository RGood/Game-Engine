package engine

import (
	"github.com/RGood/go-collection-functions/pkg/set"
)

type Interceptor interface {
	Intercept(Event) Event
	Reference() *Interceptor
}

type Listener interface {
	Notify(Event)
	Reference() *Listener
}

type Event interface {
	Execute(*Engine)
}

type Engine struct {
	queue        []Event
	interceptors *set.OrderedSet[*Interceptor]
	listeners    *set.OrderedSet[*Listener]
}

func NewEngine() *Engine {
	return &Engine{
		queue:        []Event{},
		interceptors: set.NewOrderedSet[*Interceptor](),
		listeners:    set.NewOrderedSet[*Listener](),
	}
}

func (engine *Engine) Listen(listener *Listener) {
	engine.listeners.Add(listener)
}

func (engine *Engine) Unlisten(listener *Listener) {
	engine.listeners.Remove(listener)
}

func (engine *Engine) Intercept(interceptor *Interceptor) {
	engine.interceptors.Add(interceptor)
}

func (engine *Engine) Unintercept(interceptor *Interceptor) {
	engine.interceptors.Remove(interceptor)
}

func (engine *Engine) Queue(e Event) {
	engine.queue = append(engine.queue, e)
}

type TickEvent struct {
	Iterations int
}

func (tickEvent *TickEvent) Execute(_ *Engine) {}

func (engine *Engine) RunCycles(numCycles int) {
	engine.queue = append(engine.queue, &TickEvent{1})

	cyclesComplete := 0

	for len(engine.queue) > 0 && cyclesComplete < numCycles {

		currentEvent := engine.queue[0]
		engine.queue = engine.queue[1:]

		tickEvent, isTickEvent := currentEvent.(*TickEvent)

		// Interceptors shouldn't modify tick events
		if !isTickEvent {
			engine.interceptors.ForEach(func(_ int, i *Interceptor) {
				currentEvent = (*i).Intercept(currentEvent)
			})
		}

		currentEvent.Execute(engine)

		engine.listeners.ForEach(func(_ int, l *Listener) {
			(*l).Notify(currentEvent)
		})

		// If there's nothing after the final tick event, don't re-queue it
		if isTickEvent && len(engine.queue) > 0 {
			cyclesComplete++
			if cyclesComplete < numCycles {
				engine.queue = append(engine.queue, &TickEvent{
					tickEvent.Iterations + 1,
				})
			}

		}
	}
}

func (engine *Engine) Run() {
	engine.queue = append(engine.queue, &TickEvent{1})

	for len(engine.queue) > 0 {
		currentEvent := engine.queue[0]
		engine.queue = engine.queue[1:]

		tickEvent, isTickEvent := currentEvent.(*TickEvent)

		// Interceptors shouldn't modify tick events
		if !isTickEvent {
			engine.interceptors.ForEach(func(_ int, i *Interceptor) {
				currentEvent = (*i).Intercept(currentEvent)
			})
		}

		currentEvent.Execute(engine)

		engine.listeners.ForEach(func(_ int, l *Listener) {
			(*l).Notify(currentEvent)
		})

		// If there's nothing after the final tick event, don't re-queue it
		if isTickEvent && len(engine.queue) > 0 {
			engine.queue = append(engine.queue, &TickEvent{
				tickEvent.Iterations + 1,
			})
		}
	}
}
