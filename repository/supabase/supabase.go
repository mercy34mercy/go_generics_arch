package supabase

import (
	"errors"
	"generics_pra/domain/supabase"
	"generics_pra/repository"
	"sync"
)

type MemoryCardRepository[T supabase.Card] struct {
	mu    sync.Mutex
	cards map[string]repository.CardRepository[supabase.Card]
}

func NewMemoryCardRepository[T supabase.Card]() *MemoryCardRepository[supabase.Card] {
	return &MemoryCardRepository[supabase.Card]{
		cards: make(map[string]repository.CardRepository[supabase.Card]),
	}
}

func (r *MemoryCardRepository[T]) Get(id string) (repository.CardRepository[supabase.Card], error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	card, exists := r.cards[id]
	if !exists {
		return repository.CardRepository[supabase.Card]{}, errors.New("card not found")
	}
	return card, nil
}

func (r *MemoryCardRepository[T]) Save(card repository.CardRepository[supabase.Card]) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.cards[card.CardID] = card
	return nil
}
