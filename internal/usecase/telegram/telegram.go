package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

type tgUsecase struct {
	bot tgbotapi.BotAPI
}

func New(botApi tgbotapi.BotAPI) *tgUsecase {
	return &tgUsecase{bot: botApi}
}

func (tuc *tgUsecase) SendMessage() error {
	updt := tgbotapi.NewUpdate(0)
	updt.Timeout = 60
	updates, err := tuc.bot.GetUpdatesChan(updt)
	if err != nil {
		logrus.Error(err)
	}
	for update := range updates {
		switch update.Message.Text {
		case "Start":
			logrus.Infof(update.Message.Text)
		default:
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "http://localhost:5500/")
		msg.ReplyToMessageID = update.Message.MessageID
		tuc.bot.Send(msg)
	}
	return nil
}
