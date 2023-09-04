// Button1_handler.go
package message_handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Button1MessageHandler struct {
	message string // Приватное поле для хранения сообщения
}

func NewButton1MessageHandler() *Button1MessageHandler {
	handler := &Button1MessageHandler{}
	handler.message = "Reply" // Устанавливаем сообщение, которое обрабатывается
	RegisterHandler(handler)
	return handler
}

// Геттер сообщения
func (h *Button1MessageHandler) Message() string {
	return h.message
}

// Реализация обработчика для сообщения
func (h *Button1MessageHandler) HandleMessage(update *tgbotapi.Update) []tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы нажали на первую кнопку.\r\nЗдесь вы можете увидеть сообщение с ответом на ваше.")
	msg.ReplyToMessageID = update.Message.MessageID

	response := []tgbotapi.Chattable{msg}
	return response
}

// Инициализация обработчика при импорте
func init() {
	NewButton1MessageHandler()
}
