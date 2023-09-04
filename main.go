// main.go
package main

import (
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"tgbotgolang/config"
	"tgbotgolang/message_handlers"
)

func main() {
	config, err := config.ReadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	// Использование настроек из конфигурации
	botToken := config.BotToken
	botHost := config.BotHost
	webhookPath := config.WebhookPath

	// Инициализация с ботом по токену, проверка что токен правильный
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("TelegramBot API object initialized.")

	bot.Debug = true

	// Прослушка новых обновлений получаемых на вебхук
	updates := bot.ListenForWebhook(webhookPath)
	go http.ListenAndServe(botHost, nil)
	log.Println("Listen and serve on " + botHost)

	handlers := message_handlers.GetAllHandlers()

	for update := range updates {
		log.Println("Received new message from TelegramBot.")

		if update.Message == nil {
			log.Println("Message is null. Skip message.")
			continue
		}

		msg := update.Message.Text
		handler, exists := handlers[msg]
		if !exists {
			handler = message_handlers.NewDefaultHandler()
		}

		response := handler.HandleMessage(&update)
		if response == nil {
			log.Println("Handler returned null. Expected response.")
			continue
		}

		_, err := bot.Send(response)
		if err != nil {
			log.Println(err)
		}
	}
}

/*
func сhoosePath() tgbotapi.ReplyKeyboardMarkup {
	good := tgbotapi.NewKeyboardButton(TXT_GOOD)
	evil := tgbotapi.NewKeyboardButton(TXT_EVIL)

	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(good, evil),
	)
}

func evilPath() tgbotapi.InlineKeyboardMarkup {
	var1 := tgbotapi.NewInlineKeyboardButtonData("Извинится перед ней", "sorry")
	var2 := tgbotapi.NewInlineKeyboardButtonData("Раздавить до конца", "destroy")

	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(var1, var2),
	)
}

func goodPath() tgbotapi.InlineKeyboardMarkup {
	var1 := tgbotapi.NewInlineKeyboardButtonData("Перевести обратно", "walk_back")
	var2 := tgbotapi.NewInlineKeyboardButtonData("Забыть отдать сумочку", "forget_bag")

	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(var1, var2),
	)
}
*/
