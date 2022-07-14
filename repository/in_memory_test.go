package repository

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInMemoryRepo_Get_Err(t *testing.T) {

	repo := NewInMemory()

	sol, err := repo.Get("doesn't exist")
	assert.Error(t, err)
	assert.Nil(t, sol)

}

func TestInMemoryRepo_PutGet(t *testing.T) {
	wantSolution := []string{"test", "test"}

	repo := NewInMemory()

	sol, err := repo.Get("test")
	assert.Error(t, err)
	assert.Nil(t, sol)

	repo.Put("test", wantSolution)
	sol, err = repo.Get("test")
	assert.NoError(t, err)
	assert.Equal(t, wantSolution, sol)
}

func TestInMemoryRepo_Delete_DoesntExist(t *testing.T) {
	repo := NewInMemory()

	err := repo.Delete("doesn't exist")
	assert.NoError(t, err)
}

func TestInMemory_Concurrency(t *testing.T) {
	var wg sync.WaitGroup
	wantSolution := []string{"test"}
	repo := NewInMemory()

	for i := 0; i < 100; i++ {
		go func(t *testing.T, iter int) {
			wg.Add(1)
			id := fmt.Sprintf("test%v", iter)
			repo.Put(id, wantSolution)

			got, err := repo.Get(id)
			assert.NoError(t, err)
			assert.Equal(t, wantSolution, got)

			err = repo.Delete(id)
			assert.NoError(t, err)

			got, err = repo.Get(id)
			assert.ErrorIs(t, err, InvalidIdErr)
			assert.Nil(t, got)

			wg.Done()
		}(t, i)
	}
	wg.Wait()
}
