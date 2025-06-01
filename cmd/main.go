package main

import (
	"context"
	"log/slog"
	"os"
	"tg-motivational-bot/internal/adapters"
	"tg-motivational-bot/internal/config"
	"tg-motivational-bot/internal/usecases"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func setupLogger() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	return logger
}

func main() {
	logger := setupLogger()

	if err := godotenv.Load(); err != nil {
		logger.Warn(".env file was not found")
	}

	cfg, err := config.LoadConfig(logger)
	if err != nil {
		logger.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	quotesAPI := adapters.NewZenQuotesAPI()
	translator := adapters.NewMyMemoryTranslator()
	telegramAdapter, err := adapters.NewTelegramAdapter(cfg.BotToken, cfg.ChatId)
	if err != nil {
		logger.Error("failed to load telegram adapter", "error", err)
		os.Exit(1)
	}

	fetchQuotesService := usecases.NewFetchQuoteService(quotesAPI)
	translatorService := usecases.NewTranslatorService(translator)
	sendQuoteService := usecases.NewSendQuoteService(telegramAdapter)

	if cfg.Debug {
		logger.Debug("running instantly due to debug mode")
		doTheMagic(logger, *fetchQuotesService, *translatorService, *sendQuoteService)
	}

	c := cron.New()
	defer c.Stop()

	_, err = c.AddFunc("0 3,5,7,9,11,13,15,17 * * *", func() {
		doTheMagic(logger, *fetchQuotesService, *translatorService, *sendQuoteService)
	})
	if err != nil {
		logger.Error("failed to add cronjob", "error", err)
		os.Exit(1)
	}

	c.Start()
	logger.Debug("successfully started cronjob")

	// Infinite loop
	select {}
}

func doTheMagic(logger *slog.Logger, fetchQuotesService usecases.FetchQuoteService, translatorService usecases.TranslatorService, sendQuoteService usecases.SendQuoteService) {
	ctx := context.Background()

	quote, err := fetchQuotesService.FetchQuote(ctx)
	if err != nil {
		logger.Error("failed to fetch quotes", "error", err)
		os.Exit(1)
	}
	logger.Debug("successfully fetched a random quote", "text", quote.Text, "author", quote.Author)

	translatedQuoteText, err := translatorService.Translate(ctx, quote.Text)
	if err != nil {
		logger.Error("failed to translate quote text", "error", err)
		os.Exit(1)
	}
	translatedQuoteAuthor, err := translatorService.Translate(ctx, quote.Author)
	if err != nil {
		logger.Error("failed to translate quote author", "error", err)
		os.Exit(1)
	}

	quote.Text = translatedQuoteText
	quote.Author = translatedQuoteAuthor
	logger.Debug("successfully translated quote", "text", quote.Text, "author", quote.Author)

	if err := sendQuoteService.SendQuote(ctx, quote); err != nil {
		logger.Error("failed to send quote", "error", err)
		os.Exit(1)
	}
}
