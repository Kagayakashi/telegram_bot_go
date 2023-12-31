// start_handler.go
package message_handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type StartMessageHandler struct {
	message string // Приватное поле для хранения сообщения
}

func NewStartMessageHandler() *StartMessageHandler {
	handler := &StartMessageHandler{}
	handler.message = "/start" // Устанавливаем сообщение, которое обрабатывается
	RegisterHandler(handler)
	return handler
}

// Геттер сообщения
func (h *StartMessageHandler) Message() string {
	return h.message
}

// Реализация обработчика для сообщения
func (h *StartMessageHandler) HandleMessage(update *tgbotapi.Update) []tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Это стандартный набор кнопок снизу.\r\nНажмите на одну из кнопок.")
	msg.ReplyMarkup = startButtons()

	response := []tgbotapi.Chattable{msg}
	return response
}

// Стандартные кнопки
func startButtons() tgbotapi.ReplyKeyboardMarkup {
	button1 := tgbotapi.NewKeyboardButton("Reply")
	button2 := tgbotapi.NewKeyboardButton("Forward")
	button3 := tgbotapi.NewKeyboardButton("Inline buttons")
	button4 := tgbotapi.NewKeyboardButton("Change buttons")

	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(button1, button2),
		tgbotapi.NewKeyboardButtonRow(button4, button3),
	)
}

// Инициализация обработчика при импорте
func init() {
	NewStartMessageHandler()
}
