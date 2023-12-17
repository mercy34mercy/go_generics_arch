package usecase

import (
	"context"
	"generics_pra/domain/firestore"
	"generics_pra/repository"
)

type GetCardByCardIDSupabase interface {
	Execute(ctx context.Context, cardID string) (*repository.CardRepository[firestore.Card], error)
}

var _ GetCardByCardIDSupabase = (*getCardByCardIDSupabaseImpl)(nil)

type getCardByCardIDSupabaseImpl struct {
	cardRepo repository.Repository[repository.CardRepository[firestore.Card]]
}

func (uc *getCardByCardIDSupabaseImpl) Execute(ctx context.Context, cardID string) (*repository.CardRepository[firestore.Card], error) {
	card, err := uc.cardRepo.Get(cardID)
	if err != nil {
		return nil, err
	}
	return &card, nil
}

func NewGetCardByCardIDSupabase(cardRepo repository.Repository[repository.CardRepository[firestore.Card]]) GetCardByCardIDSupabase {
	return &getCardByCardIDSupabaseImpl{
		cardRepo: cardRepo,
	}
}
