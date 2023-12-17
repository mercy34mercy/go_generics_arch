package usecase

import (
	"context"
	"generics_pra/domain/supabase"
	"generics_pra/repository"
)

type SaveCardSupabase interface {
	Execute(ctx context.Context, card repository.CardRepository[supabase.Card]) error
}

var _ SaveCardSupabase = (*saveCardSupabaseImpl)(nil)

type saveCardSupabaseImpl struct {
	cardRepo repository.Repository[repository.CardRepository[supabase.Card]]
}

func (uc *saveCardSupabaseImpl) Execute(ctx context.Context, card repository.CardRepository[supabase.Card]) error {
	return uc.cardRepo.Save(card)
}

func NewSaveSupabaseCard(cardRepo repository.Repository[repository.CardRepository[supabase.Card]]) SaveCardSupabase {
	return &saveCardSupabaseImpl{
		cardRepo: cardRepo,
	}
}
