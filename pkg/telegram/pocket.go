package telegram

import (
	pocket "github.com/ainurqa95/pocket/v3"
)

func (bot *Bot) InitAuthorization(chatId int64) (string, error) {
	existedRequestToken, err := bot.tokenRepository.GetRequestToken(chatId)
	if existedRequestToken != "" {
		return bot.pocketClient.DefineAuthorizationUrl(existedRequestToken, bot.redirectUrl)
	}

	token, err := bot.pocketClient.GetRequestToken(bot.redirectUrl)
	if err != nil {
		return "", err
	}
	err = bot.tokenRepository.SaveRequestToken(chatId, token)
	if err != nil {
		return "", err
	}

	return bot.pocketClient.DefineAuthorizationUrl(token, bot.redirectUrl)
}

func (bot *Bot) AddLinkToPocket(chatId int64, link string) error {

	accessToken, err := bot.defineAccessToken(chatId)

	if err != nil {
		return err
	}

	err = bot.pocketClient.AddItem(pocket.AddInput{
		AccessToken: accessToken,
		Url:         link,
	})

	if err != nil {
		return unableToSaveError
	}

	return nil
}

func (bot *Bot) defineAccessToken(chatId int64) (string, error) {

	requestToken, err := bot.tokenRepository.GetRequestToken(chatId)
	if err != nil || requestToken == "" {
		return "", emptyRequestTokenError
	}

	existedToken, err := bot.tokenRepository.GetAccessToken(chatId)

	if existedToken != "" {
		return existedToken, nil
	}

	accessToken, err := bot.pocketClient.AuthAndGetAccessToken(requestToken)

	if err != nil {
		return "", emptyAccessTokenError
	}

	err = bot.tokenRepository.SaveAccessToken(chatId, accessToken.AccessToken)

	if err != nil {
		return "", err
	}

	return accessToken.AccessToken, nil
}
