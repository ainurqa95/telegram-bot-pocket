package boltdb

import (
	"strconv"

	"github.com/boltdb/bolt"
)

const (
	RequestTokenKey = "request_token"
	AccessTokenKey  = "access_token"
)

type DbTokenRepository struct {
	db *bolt.DB
}

func NewDbTokenRepository(db *bolt.DB) *DbTokenRepository {
	return &DbTokenRepository{db: db}
}

func (tokenRepository *DbTokenRepository) SaveAccessToken(chatId int64, accessToken string) error {
	return saveToken(tokenRepository, AccessTokenKey, chatId, accessToken)
}

func (tokenRepository *DbTokenRepository) SaveRequestToken(chatId int64, requestToken string) error {
	return saveToken(tokenRepository, RequestTokenKey, chatId, requestToken)
}

func saveToken(tokenRepository *DbTokenRepository, tokenKey string, chatId int64, tokenValue string) error {
	return tokenRepository.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(tokenKey))
		key := intToBytes(chatId)
		return bucket.Put(key, []byte(tokenValue))
	})
}

func (tokenRepository *DbTokenRepository) GetRequestToken(chatId int64) (string, error) {
	return getToken(tokenRepository, chatId, RequestTokenKey)
}

func (tokenRepository *DbTokenRepository) GetAccessToken(chatId int64) (string, error) {
	return getToken(tokenRepository, chatId, AccessTokenKey)
}

func getToken(tokenRepository *DbTokenRepository, chatId int64, tokenKey string) (string, error) {
	var token string
	err := tokenRepository.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(tokenKey))
		key := intToBytes(chatId)
		data := bucket.Get(key)
		token = string(data)
		return nil
	})
	if err != nil {
		return "", err
	}

	return token, err
}

func intToBytes(v int64) []byte {
	return []byte(strconv.FormatInt(v, 10))
}
