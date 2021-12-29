package storage

type Bucket string

type TokenRepository interface {
	SaveAccessToken(chatId int64, accessToken string) error
	SaveRequestToken(chatId int64, requestToken string) error
	GetRequestToken(chatId int64) (string, error)
	GetAccessToken(chatId int64) (string, error)
}
