package config

import (
	"errors"
	"log/slog"
	"os"
	"strconv"
)

type Config struct {
	Debug    bool
	BotToken string
	ChatId   int64
}

func LoadConfig(logger *slog.Logger) (*Config, error) {
	env := os.Getenv("ENV")
	botToken := os.Getenv("BOT_TOKEN")
	chatId := os.Getenv("CHAT_ID")

	if botToken == "" || chatId == "" {
		logger.Error("Required environmental variables are missing", "BOT_TOKEN", botToken, "CHAT_ID", chatId)
		return nil, errors.New("required environmental variables are missing")
	}

	chatIdInt, err := strconv.ParseInt(chatId, 10, 64)
	if err != nil {
		logger.Error("Cannot parse CHAT_ID into int64", chatId, err)
		return nil, errors.New("cannot parse CHAT_ID into int64")
	}

	return &Config{BotToken: botToken, ChatId: chatIdInt, Debug: env == "DEBUG"}, nil
}
