package telegram

import (
	"github.com/ainurqa95/telegram-bot-pocket/pkg/config"
	"github.com/ainurqa95/telegram-bot-pocket/pkg/storage"

	pocket "github.com/ainurqa95/pocket/v3"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	client          *tgbotapi.BotAPI
	pocketClient    *pocket.PocketClient
	tokenRepository storage.TokenRepository
	redirectUrl     string
	messages        config.Messages
}

func NewBot(
	client *tgbotapi.BotAPI,
	pocketClient *pocket.PocketClient,
	tokenRepository storage.TokenRepository,
	redirectUrl string,
	messages config.Messages,
) *Bot {
	return &Bot{
		client:          client,
		pocketClient:    pocketClient,
		tokenRepository: tokenRepository,
		redirectUrl:     redirectUrl,
		messages:        messages,
	}
}

func (bot *Bot) Start() {
	updates := bot.initUpdates()

	bot.handleUpdates(updates)
}

func (bot *Bot) initUpdates() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return bot.client.GetUpdatesChan(u)
}

func (bot *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			err := bot.handleCommand(update.Message)
			if err != nil {
				bot.handleError(update.Message.Chat.ID, err)
			}
			continue
		}

		err := bot.handleMessage(update.Message)
		if err != nil {
			bot.handleError(update.Message.Chat.ID, err)
		}
	}
}
