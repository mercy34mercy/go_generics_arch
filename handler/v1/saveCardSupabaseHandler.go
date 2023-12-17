package handler

import (
	"encoding/json"
	"generics_pra/domain/supabase"
	"generics_pra/repository"
	"generics_pra/usecase"
	"net/http"
)

type SaveSupabaseCardRequest struct {
	CardID       string        `json:"card_id"`
	Name         string        `json:"name"`
	Organization string        `json:"organization"`
	CardData     supabase.Card `json:"card_data"`
}

type SaveSupabaseCardHandler struct {
	saveCardUseCase usecase.SaveCardSupabase
}

func NewSaveSupabaseCardHandler(saveCardUseCase usecase.SaveCardSupabase) *SaveSupabaseCardHandler {
	return &SaveSupabaseCardHandler{
		saveCardUseCase: saveCardUseCase,
	}
}

func (h *SaveSupabaseCardHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SaveSupabaseCardRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cardRepo := repository.CardRepository[supabase.Card]{
		CardID:       req.CardID,
		Name:         req.Name,
		Organization: req.Organization,
		CustomFields: req.CardData,
	}

	if err := h.saveCardUseCase.Execute(r.Context(), cardRepo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
