package usecases

import (
	"context"
	"fmt"
	"tg-motivational-bot/internal/entities"
	"tg-motivational-bot/internal/interfaces"
)

type SendQuoteService struct {
	telegram interfaces.TelegramSender
}

func (s *SendQuoteService) SendQuote(ctx context.Context, quote *entities.Quote) error {
	message := fmt.Sprintf("📖 %s\n\n _%s_ ✍️", quote.Text, quote.Author)

	if err := s.telegram.SendMessage(ctx, message); err != nil {
		return fmt.Errorf("failed to send the message: %w", err)
	}

	return nil
}

// Accepts TelegramSender interface implementation and binds it to the returning SendQuoteService  
func NewSendQuoteService(telegram interfaces.TelegramSender) *SendQuoteService {
	return &SendQuoteService{telegram: telegram}
}
