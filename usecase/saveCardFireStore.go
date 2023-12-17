package usecase

import (
	"context"
	"generics_pra/domain/firestore"
	"generics_pra/repository"
)

type SaveCardFirestore interface {
	Execute(ctx context.Context, card repository.CardRepository[firestore.Card]) error
}

var _ SaveCardFirestore = (*saveCardFirestoreImpl)(nil)

type saveCardFirestoreImpl struct {
	cardRepo repository.Repository[repository.CardRepository[firestore.Card]]
}

func (uc *saveCardFirestoreImpl) Execute(ctx context.Context, card repository.CardRepository[firestore.Card]) error {
	return uc.cardRepo.Save(card)
}

func NewSaveCard(cardRepo repository.Repository[repository.CardRepository[firestore.Card]]) SaveCardFirestore {
	return &saveCardFirestoreImpl{
		cardRepo: cardRepo,
	}
}
