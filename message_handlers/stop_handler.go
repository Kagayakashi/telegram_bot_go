package message_handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type StopMessageHandler struct {
	message string // Приватное поле для хранения сообщения
}

func NewStopMessageHandler() *StopMessageHandler {
	handler := &StopMessageHandler{}
	handler.message = "/stop" // Устанавливаем сообщение, которое обрабатывается
	RegisterHandler(handler)
	return handler
}

// Геттер сообщения
func (h *StopMessageHandler) Message() string {
	return h.message
}

// Реализация обработчика для сообщения
func (h *StopMessageHandler) HandleMessage(update *tgbotapi.Update) tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы остановили бота и убрали кнопки.")
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	return msg
}

func init() {
	NewStopMessageHandler()
}
