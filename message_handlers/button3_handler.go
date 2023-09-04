// Button3_handler.go
package message_handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Button3MessageHandler struct {
	message string // Приватное поле для хранения сообщения
}

func NewButton3MessageHandler() *Button3MessageHandler {
	handler := &Button3MessageHandler{}
	handler.message = "Inline buttons" // Устанавливаем сообщение, которое обрабатывается
	RegisterHandler(handler)
	return handler
}

// Геттер сообщения
func (h *Button3MessageHandler) Message() string {
	return h.message
}

// Реализация обработчика для сообщения
func (h *Button3MessageHandler) HandleMessage(update *tgbotapi.Update) []tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы нажали на третью кнопку.\r\nЗдесь вы можете увидеть ответ с приложенными кнопками.")
	msg.ReplyMarkup = button3Buttons()

	response := []tgbotapi.Chattable{msg}
	return response
}

// Приложенные к сообщению кнопки
func button3Buttons() tgbotapi.InlineKeyboardMarkup {
	button1 := tgbotapi.NewInlineKeyboardButtonData("Кнопка с командой без отправки текста", "inline_button")
	button2 := tgbotapi.NewInlineKeyboardButtonURL("Кнопка ссылка на телеграм", "https://telegram.org/")
	button3 := tgbotapi.NewInlineKeyboardButtonSwitch("Кнопка поделится ботом", "Смотри какой крутой бот!")

	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(button1),
		tgbotapi.NewInlineKeyboardRow(button2),
		tgbotapi.NewInlineKeyboardRow(button3),
	)
}

// Инициализация обработчика при импорте
func init() {
	NewButton3MessageHandler()
}
