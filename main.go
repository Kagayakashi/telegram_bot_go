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

		msg := ""

		if update.Message != nil {
			msg = update.Message.Text
		}

		callback := update.CallbackData()
		if callback != "" {
			msg = message_handlers.CALLBACK_QUERY_PREFIX + callback
		}

		// Если имеется callback, использовать как msg для обработчика
		handler, exists := handlers[msg]
		if !exists {
			handler = message_handlers.NewDefaultHandler()
		}

		// В ответе может быть одно или массив сообщений
		responses := handler.HandleMessage(&update)

		if responses == nil {
			log.Println("Handler returned null. Expected response.")
			continue
		}

		// Отправить массив сообщений
		for _, response := range responses {
			_, err := bot.Send(response)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
