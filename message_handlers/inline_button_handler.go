// InlineButtonCb_handler.go
package message_handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type InlineButtonCbMessageHandler struct {
	message string // Приватное поле для хранения сообщения
}

func NewInlineButtonCbMessageHandler() *InlineButtonCbMessageHandler {
	handler := &InlineButtonCbMessageHandler{}
	// Callback Query обработчики надо регистрировать со специальным префиксом
	handler.message = CALLBACK_QUERY_PREFIX + "inline_button" // Устанавливаем сообщение, которое обрабатывается
	RegisterHandler(handler)
	return handler
}

// Геттер сообщения
func (h *InlineButtonCbMessageHandler) Message() string {
	return h.message
}

// Реализация обработчика для сообщения
func (h *InlineButtonCbMessageHandler) HandleMessage(update *tgbotapi.Update) []tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(update.FromChat().ID, "Ответ на нажатие кнопки без отправки сообщения.")

	response := []tgbotapi.Chattable{msg}
	return response
}

// Инициализация обработчика при импорте
func init() {
	NewInlineButtonCbMessageHandler()
}
