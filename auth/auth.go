package auth

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

func IsValidUserName(username string, validUsername string) bool {
	return strings.EqualFold(username, validUsername)
}

func IsAuthorisedUser(message *tgbotapi.Message, validUsername string) bool {
	return IsValidUserName(message.From.UserName, validUsername)
}
