package session

import (
	"sync"

	"github.com/snanovskyi/ooooooh/protocol/server"
	"github.com/snanovskyi/ooooooh/socket"
)

type Registry struct {
	mu       sync.RWMutex
	sessions map[socket.Socket]*Session
}

func NewRegistry() *Registry {
	return &Registry{
		sessions: make(map[socket.Socket]*Session),
	}
}

func (r *Registry) Add(s *Session) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.sessions[s.Socket()] = s
}

func (r *Registry) Get(s socket.Socket) *Session {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.sessions[s]
}

func (r *Registry) Delete(s socket.Socket) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.sessions, s)
}

func (r *Registry) Broadcast(m server.Message) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var wg sync.WaitGroup
	for _, session := range r.sessions {
		s := session
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.Send(m)
		}()
	}
	wg.Wait()
}
