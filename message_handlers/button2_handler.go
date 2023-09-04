// Button2_handler.go
package message_handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Button2MessageHandler struct {
	message string // Приватное поле для хранения сообщения
}

func NewButton2MessageHandler() *Button2MessageHandler {
	handler := &Button2MessageHandler{}
	handler.message = "Forward" // Устанавливаем сообщение, которое обрабатывается
	RegisterHandler(handler)
	return handler
}

// Геттер сообщения
func (h *Button2MessageHandler) Message() string {
	return h.message
}

// Реализация обработчика для сообщения
func (h *Button2MessageHandler) HandleMessage(update *tgbotapi.Update) []tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы нажали на вторую кнопку.\r\nЗдесь вы можете увидеть ответ с пересланным сообщением.")

	fwd := tgbotapi.NewForward(update.Message.Chat.ID, update.Message.Chat.ID, update.Message.MessageID)

	response := []tgbotapi.Chattable{msg, fwd}
	return response
}

// Инициализация обработчика при импорте
func init() {
	NewButton2MessageHandler()
}
