package main

import "sync"

type Store struct {
	data map[int]bool
	mu   sync.Mutex
}

func (s *Store) Add(i int) {
	s.mu.Lock()
	s.data[i] = true
	s.mu.Unlock()
}

func (s *Store) Get(i int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.data[i]

}

func main() {

}
