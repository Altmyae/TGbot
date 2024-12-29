package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

// func sticker(){

// }

func main() {
	botToken := "your-token-here"
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	updates, _ := bot.UpdatesViaLongPolling(nil)
	bh, _ := th.NewBotHandler(bot, updates)
	defer bh.Stop()
	defer bot.StopLongPolling()

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		//Обработчик команды /start
		chatID := tu.ID(update.Message.Chat.ID)
		keyboard := tu.Keyboard(
			tu.KeyboardRow(
				tu.KeyboardButton("Старт"),
				tu.KeyboardButton("Помощь"),
				tu.KeyboardButton("Тех. поддержка"),
			),
			tu.KeyboardRow(
				tu.KeyboardButton("Локация").WithRequestLocation(),
				tu.KeyboardButton("Контакт").WithRequestContact(),
				tu.KeyboardButton("Отмена"),
			),
		)
		message := tu.Message(
			chatID,
			"С этим сообщением пользователю придет клавиатура",
		).WithReplyMarkup(keyboard)
		_, _ = bot.SendMessage(message)

		for update := range updates {
			chatID := tu.ID(update.Message.Chat.ID)
			_, _ = bot.SendSticker(
				tu.Sticker(
					chatID,
					tu.FileFromID(
						"CAACAgIAAxkBAAENaG9ncUl9anT_V1bRuqi2YZwnFVbAWgACSRQAAo9ZqEjCNkSEChgZmDYE"),
				),
			)
		}

	}, th.CommandEqual("start"))
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		//Обработчик любого другого сообщения
	}, th.AnyMessage())
	bh.Start()
}
