package adapters

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"tg-motivational-bot/internal/entities"
)

const ZenQuotesURL string = "https://zenquotes.io/api/random"

type ZenQuotesAPI struct{}

func (z *ZenQuotesAPI) GetRandomQuote(ctx context.Context) (*entities.Quote, error) {
	resp, err := http.Get(ZenQuotesURL)
	if err != nil {
		return nil, errors.New("failed to fetch response from zenquotes api")
	}
	defer resp.Body.Close()

	var quotes []struct {
		Quote  string `json:"q"`
		Author string `json:"a"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&quotes); err != nil {
		return nil, errors.New("failed to decode response body from zenquotes api")
	}

	if len(quotes) == 0 {
		return nil, errors.New("empty list of quotes from zenquotes api")
	}

	return &entities.Quote{
		Text:   quotes[0].Quote,
		Author: quotes[0].Author,
	}, nil
}

func NewZenQuotesAPI() *ZenQuotesAPI {
	return &ZenQuotesAPI{}
}