package adapters

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

const TranslatorURL string = "https://api.mymemory.translated.net/get"

type MyMemoryTranslator struct{}

func (t *MyMemoryTranslator) Translate(ctx context.Context, text, origLang, targetLang string) (string, error) {
	params := url.Values{}
	params.Set("q", text)
	params.Set("langpair", origLang+"|"+targetLang)

	resp, err := http.Get(TranslatorURL + "?" + params.Encode())
	if err != nil {
		return "", errors.New("failed to fetch response from translate api")
	}
	defer resp.Body.Close()

	var result struct {
		ResponseData struct {
			TranslatedText string `json:"translatedText"`
		} `json:"responseData"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", errors.New("failed to decode response body from translate api")
	}

	if result.ResponseData.TranslatedText == "" {
		return "", errors.New("empty response from translate api")
	}

	return result.ResponseData.TranslatedText, nil
}

func NewMyMemoryTranslator() *MyMemoryTranslator {
	return &MyMemoryTranslator{}
}
