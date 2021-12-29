package main

import (
	"fmt"
	"log"

	pocket "github.com/ainurqa95/pocket/v3"
	"github.com/ainurqa95/telegram-bot-pocket/pkg/config"
	"github.com/ainurqa95/telegram-bot-pocket/pkg/storage/boltdb"
	"github.com/ainurqa95/telegram-bot-pocket/pkg/telegram"
	"github.com/boltdb/bolt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	config, err := config.Init()

	if err != nil {
		log.Fatal(err)
	}

	bot, err := tgbotapi.NewBotAPI(config.TelegramToken)
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true

	pocketClient := pocket.NewPocketClient(config.PocketConsumerKey)
	db, err := initBolt(config)
	if err != nil {
		log.Fatal(err)
	}
	tokenRepository := boltdb.NewDbTokenRepository(db)

	telegramBot := telegram.NewBot(bot, pocketClient, tokenRepository, config.BotURL, config.Messages)
	telegramBot.Start()
	fmt.Println(telegramBot)
}

func initBolt(config *config.Config) (*bolt.DB, error) {
	db, err := bolt.Open(config.BoltDBFile, 0600, nil)
	if err != nil {
		return nil, err
	}

	if err := db.Batch(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(boltdb.AccessTokenKey))
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists([]byte(boltdb.RequestTokenKey))
		return err
	}); err != nil {
		return nil, err
	}

	return db, nil
}
