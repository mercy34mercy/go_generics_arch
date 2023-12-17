package usecase

import (
	"context"
	"generics_pra/domain/firestore"
	"generics_pra/repository"
)

type GetCardByCardIDFirestore interface {
	Execute(ctx context.Context, cardID string) (*firestore.Card, error)
}

var _ GetCardByCardIDFirestore = (*getCardByCardIDFirestoreImpl)(nil)

type getCardByCardIDFirestoreImpl struct {
	cardRepo repository.Repository[firestore.Card]
}

func (uc *getCardByCardIDFirestoreImpl) Execute(ctx context.Context, cardID string) (*firestore.Card, error) {
	card, err := uc.cardRepo.Get(cardID)
	if err != nil {
		return nil, err
	}
	return &card, nil
}

func NewGetCardByCardIDFirestore(cardRepo repository.Repository[firestore.Card]) GetCardByCardIDFirestore {
	return &getCardByCardIDFirestoreImpl{
		cardRepo: cardRepo,
	}
}
