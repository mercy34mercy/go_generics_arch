package repository

type CardRepository[T any] struct {
	CardID       string
	Name         string
	Organization string
	CustomFields T
}

type Repository[T any] interface {
	Get(id string) (T, error)
	Save(entity T) error
}
