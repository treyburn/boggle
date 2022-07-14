package repository

import (
	"errors"
	"sync"
)

var InvalidIdErr = errors.New("invalid id")

type InMemoryRepo struct {
	mu    sync.RWMutex
	store map[string][]string
}

func (r *InMemoryRepo) Get(id string) ([]string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	solution, ok := r.store[id]
	if !ok {
		return nil, InvalidIdErr
	}
	return solution, nil
}

func (r *InMemoryRepo) Put(id string, solutions []string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.store[id] = solutions
}

func (r *InMemoryRepo) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.store, id)
	return nil
	// always returning nil here because the error return isn't necessary
	// but would be useful in other repo implementations (such as redis or a db)
}

func NewInMemory() *InMemoryRepo {
	s := make(map[string][]string)

	return &InMemoryRepo{
		store: s,
	}
}
