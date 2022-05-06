package engine

import "fmt"

type mockInterceptor struct {
	events []Event
	ref    Interceptor
}

func (mi *mockInterceptor) Intercept(e Event) Event {
	mi.events = append(mi.events, e)
	return e
}

func (mi *mockInterceptor) Reference() *Interceptor {
	if mi.ref == nil {
		mi.ref = mi
	}
	return &mi.ref
}

type mockListener struct {
	events []Event
	ref    Listener
}

func (ml *mockListener) Notify(e Event) {
	ml.events = append(ml.events, e)
}

func (ml *mockListener) Reference() *Listener {
	if ml.ref == nil {
		ml.ref = ml
	}
	return &ml.ref
}

type NoopEvent struct {
}

func (noop *NoopEvent) Execute(_ *Engine) {

}

type RepeaterEvent struct {
	count     int
	queueFunc func(Event)
}

func (re *RepeaterEvent) Execute(_ *Engine) {
	fmt.Printf("Count: %d\n", re.count)
	if re.count > 1 {
		re.count--
		re.queueFunc(re)
	}
}
