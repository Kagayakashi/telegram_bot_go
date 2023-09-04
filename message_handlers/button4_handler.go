// Button4_handler.go
package message_handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Button4MessageHandler struct {
	message string // Приватное поле для хранения сообщения
}

func NewButton4MessageHandler() *Button4MessageHandler {
	handler := &Button4MessageHandler{}
	handler.message = "Change buttons" // Устанавливаем сообщение, которое обрабатывается
	RegisterHandler(handler)
	return handler
}

// Геттер сообщения
func (h *Button4MessageHandler) Message() string {
	return h.message
}

// Реализация обработчика для сообщения
func (h *Button4MessageHandler) HandleMessage(update *tgbotapi.Update) []tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Кнопки снизу поменялись.\r\nНажмите на одну из кнопок.")
	msg.ReplyMarkup = Button4Buttons()

	response := []tgbotapi.Chattable{msg}
	return response
}

// Стандартные кнопки
func Button4Buttons() tgbotapi.ReplyKeyboardMarkup {
	button1 := tgbotapi.NewKeyboardButton("Reply")
	button2 := tgbotapi.NewKeyboardButton("Forward")
	button3 := tgbotapi.NewKeyboardButton("Inline buttons")
	button4 := tgbotapi.NewKeyboardButton("Recover buttons")

	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(button3, button4),
		tgbotapi.NewKeyboardButtonRow(button2, button1),
	)
}

// Инициализация обработчика при импорте
func init() {
	NewButton4MessageHandler()
}
