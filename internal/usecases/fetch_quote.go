package usecases

import (
	"context"
	"errors"
	"tg-motivational-bot/internal/entities"
	"tg-motivational-bot/internal/interfaces"
)

type FetchQuoteService struct {
	api interfaces.QuoteAPI
}

func (s *FetchQuoteService) FetchQuote(ctx context.Context) (*entities.Quote, error) {
	quote, err := s.api.GetRandomQuote(ctx)
	if err != nil {
		return nil, errors.New("failed to fetch a random quote")
	}
	return quote, nil
}

// Accepts QuoteAPI interface implementation and binds it to the returning FetchQuoteService
func NewFetchQuoteService(api interfaces.QuoteAPI) *FetchQuoteService {
	return &FetchQuoteService{api: api}
}
