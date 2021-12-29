package telegram

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	invalidUrlError        = errors.New("Url is invalid")
	unableToSaveError      = errors.New("Unable to save link to Pocket")
	emptyRequestTokenError = errors.New("Empty request token")
	emptyAccessTokenError  = errors.New("Empty accesss token")
)

func (bot *Bot) handleError(chatID int64, err error) {
	messageText := err.Error()

	switch err {
	case invalidUrlError:
		messageText = bot.messages.Errors.InvalidURL
	case unableToSaveError:
		messageText = bot.messages.Errors.UnableToSave
	case emptyAccessTokenError:
		messageText = bot.messages.EmptyAccessToken
	case emptyRequestTokenError:
		messageText = bot.messages.EmptyRequestToken
	default:
		messageText = bot.messages.Errors.Default
	}

	msg := tgbotapi.NewMessage(chatID, messageText)
	bot.client.Send(msg)
}
