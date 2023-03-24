package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"testing"
)

func TestGetMessageOrEditedMessageWithMessage(t *testing.T) {
	update := tgbotapi.Update{
		Message: &tgbotapi.Message{},
	}
	got, err := GetMessageOrEditedMessage(update)
	if got != update.Message {
		t.Errorf("expected update.Message")
	}
	if err != nil {
		t.Errorf("got error %q, wanted nil", err)
	}
	if got == nil {
		t.Errorf("got nil, wanted message")
	}
}

func TestGetMessageOrEditedMessageWithErrorMessage(t *testing.T) {
	update := tgbotapi.Update{
		EditedMessage: &tgbotapi.Message{},
	}
	got, err := GetMessageOrEditedMessage(update)
	if got != update.EditedMessage {
		t.Errorf("expected update.EditedMessage")
	}
	if err != nil {
		t.Errorf("got error %q, wanted nil", err)
	}
	if got == nil {
		t.Errorf("got nil, wanted message")
	}
}

func TestGetMessageOrEditedMessageWithNoMessages(t *testing.T) {
	update := tgbotapi.Update{}
	got, err := GetMessageOrEditedMessage(update)
	if got != nil {
		t.Errorf("expected no message")
	}
	if err == nil {
		t.Errorf("got nil, wanted error")
	}
}
