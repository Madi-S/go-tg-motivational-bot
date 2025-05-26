package interfaces

import (
	"context"
	"tg-motivational-bot/internal/entities"
)

type QuoteAPI interface {
	GetRandomQuote(ctx context.Context) (*entities.Quote, error)
}

type Translator interface {
	Translate(ctx context.Context, text, origLang, targetLang string) (string, error)
}

type TelegramSender interface {
	SendMessage(ctx context.Context, message string) error
}

type CronScheduler interface {
	Start()
	AddJob(spec string, job func())
}
