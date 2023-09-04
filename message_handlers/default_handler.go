package message_handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type DefaultHandler struct {
	message string // Приватное поле для хранения сообщения
}

func NewDefaultHandler() *DefaultHandler {
	handler := &DefaultHandler{}
	handler.message = "/default" // Устанавливаем сообщение, которое обрабатывается
	// Регистрация обработчика не нужна, ему надо на любое сообщение триггериться
	return handler
}

// Геттер сообщения
func (h *DefaultHandler) Message() string {
	return h.message
}

// Реализация обработчика для сообщения по умолчанию
func (h *DefaultHandler) HandleMessage(update *tgbotapi.Update) tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неизвестная команда или сообщение.")
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	return msg
}
