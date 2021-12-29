package telegram

import (
	"fmt"
	"net/url"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	startCommand = "start"
)

func (bot *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case startCommand:
		return bot.handleStartCommand(message)
	default:
		return bot.handleUnknownCommand(message)
	}
}

func (bot *Bot) handleStartCommand(message *tgbotapi.Message) error {
	authUrl, err := bot.InitAuthorization(message.Chat.ID)

	authUrlMessage := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(bot.messages.Start, authUrl))

	_, err = bot.client.Send(authUrlMessage)

	return err
}

func (bot *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, bot.messages.UnknownCommand)
	_, err := bot.client.Send(msg)

	return err
}

func (bot *Bot) handleMessage(message *tgbotapi.Message) error {
	link := message.Text
	err := bot.validateURL(link)
	if err != nil {
		return invalidUrlError
	}

	err = bot.AddLinkToPocket(message.Chat.ID, link)

	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, bot.messages.LinkSaved)

	_, err = bot.client.Send(msg)

	return err
}

func (b *Bot) validateURL(text string) error {
	_, err := url.ParseRequestURI(text)
	return err
}
