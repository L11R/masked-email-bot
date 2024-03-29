package domain

import (
	"context"
	"golang.org/x/oauth2"
	"net/url"
)

type Database interface {
	CreateUser(telegramID int64, languageCode string) error
	UpdateToken(telegramID int64, fastmailToken string) error
	UpdateLanguageCode(telegramID int64, languageCode string) error
	GetUser(telegramID int64) (*User, error)

	CreateOAuth2State(state, codeVerifier string, telegramID int64) error
	GetOAuth2State(state string) (*OAuth2State, error)

	Close() error
	NewTokenSource(baseTokenSource oauth2.TokenSource, telegramID int64) oauth2.TokenSource
}

type MaskingEmail interface {
	CreateMaskedEmailFromURL(ctx context.Context, tokenSrc oauth2.TokenSource, url *url.URL) (*MaskedEmail, error)
	CreateMaskedEmailWithPrefix(ctx context.Context, tokenSrc oauth2.TokenSource, prefix string) (*MaskedEmail, error)
	EnableMaskedEmail(ctx context.Context, tokenSrc oauth2.TokenSource, id string) error
	GetOAuth2Config() *oauth2.Config
}

type Delivery interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}

type Telegram interface {
	SendMessage(telegramID int64, languageCode, messageID string) error
}
