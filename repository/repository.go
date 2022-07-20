package repository

type Repository interface {
	Get(id string) ([]string, error)
	Put(id string, solutions []string)
	Delete(id string) error
}
