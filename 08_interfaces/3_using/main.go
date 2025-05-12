package main

type Model interface {
	ID() int
}

type Data map[int]Model

type Store struct {
	data Data
}

func (s *Store) Insert(m Model) error {
	s.data[m.ID()] = m
	return nil
}

func main() {}
