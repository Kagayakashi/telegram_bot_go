// Button5_handler.go
package message_handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Button5MessageHandler struct {
	message string // Приватное поле для хранения сообщения
}

func NewButton5MessageHandler() *Button5MessageHandler {
	handler := &Button5MessageHandler{}
	handler.message = "Recover buttons" // Устанавливаем сообщение, которое обрабатывается
	RegisterHandler(handler)
	return handler
}

// Геттер сообщения
func (h *Button5MessageHandler) Message() string {
	return h.message
}

// Реализация обработчика для сообщения
func (h *Button5MessageHandler) HandleMessage(update *tgbotapi.Update) []tgbotapi.Chattable {
	// Переиспользовать имеющийся обработчик
	return GetAllHandlers()["/start"].HandleMessage(update)
}

// Инициализация обработчика при импорте
func init() {
	NewButton5MessageHandler()
}
