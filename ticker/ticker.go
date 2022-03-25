package ticker

import (
	"sync"
	"time"
)

type Ticker struct {
	addMu     sync.Mutex
	tickMu    sync.Mutex
	nextTick  []func()
	everyTick []func()
}

func (t *Ticker) NextTick(fn func()) {
	t.addMu.Lock()
	t.tickMu.Lock()
	defer t.tickMu.Unlock()
	defer t.addMu.Unlock()
	t.nextTick = append(t.nextTick, fn)
}

func (t *Ticker) EveryTick(fn func()) {
	t.addMu.Lock()
	t.tickMu.Lock()
	defer t.tickMu.Unlock()
	defer t.addMu.Unlock()
	t.everyTick = append(t.everyTick, fn)
}

func (t *Ticker) Run() {
	for range time.Tick(1000 / 20 * time.Millisecond) {
		t.tick()
	}
}

func (t *Ticker) tick() {
	t.tickMu.Lock()
	defer t.tickMu.Unlock()
	for _, fn := range t.nextTick {
		fn()
	}
	for _, fn := range t.everyTick {
		fn()
	}
	t.nextTick = nil
}
