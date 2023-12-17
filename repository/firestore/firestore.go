package firestore

import (
	"errors"
	"generics_pra/domain/firestore"
	"generics_pra/repository"
	"sync"
)

type MemoryCardRepository[T any] struct {
	mu    sync.Mutex
	cards map[string]repository.CardRepository[firestore.Card]
}

func NewMemoryCardRepository[T any]() *MemoryCardRepository[firestore.Card] {
	return &MemoryCardRepository[firestore.Card]{
		cards: make(map[string]repository.CardRepository[firestore.Card]),
	}
}

func (r *MemoryCardRepository[T]) Get(id string) (repository.CardRepository[firestore.Card], error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	card, exists := r.cards[id]
	if !exists {
		return repository.CardRepository[firestore.Card]{}, errors.New("card not found")
	}
	return card, nil
}

func (r *MemoryCardRepository[T]) Save(card repository.CardRepository[firestore.Card]) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.cards[card.CardID] = card
	return nil
}
