package telegram

import (
	"encoding/json"
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io/ioutil"
	"log"
	"net/http"
)

func GetUpdate(r *http.Request) tgbotapi.Update {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var update tgbotapi.Update
	err = json.Unmarshal(body, &update)
	log.Printf("update: %+v\n", update)

	return update
}

func GetMessageOrEditedMessage(update tgbotapi.Update) (*tgbotapi.Message, error) {
	if update.Message != nil {
		return update.Message, nil
	} else if update.EditedMessage != nil {
		return update.EditedMessage, nil
	}

	return nil, errors.New("no message")
}
