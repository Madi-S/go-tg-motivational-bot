package usecases

import (
	"context"
	"fmt"
	"tg-motivational-bot/internal/interfaces"
)

type TranslatorService struct {
	translator interfaces.Translator
}

func (s *TranslatorService) Translate(ctx context.Context, text string) (string, error) {
	translatedText, err := s.translator.Translate(ctx, text, "en", "kk")

	if err != nil {
		return "", fmt.Errorf("failed to translate text: %w", err)

	}

	return translatedText, nil
}

// Accepts Translator interface implementation and binds it to the returning TranslatorService  
func NewTranslatorService(translator interfaces.Translator) *TranslatorService {
	return &TranslatorService{translator: translator}
}
