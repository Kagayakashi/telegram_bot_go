// handlers.go
package message_handlers

import (
	"log"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Глобальная переменная для хранения обработчиков
var messageHandlers = make(map[string]MessageHandler)
var mu sync.RWMutex // Мьютекс для безопасности доступа к коллекции, механизм локов как в БД

// Интерфейс для реализаций обработчиков
type MessageHandler interface {
	HandleMessage(update *tgbotapi.Update) tgbotapi.Chattable
	Message() string
}

// Глобальный метод для регистрации обработчика в общий пул
func RegisterHandler(handler MessageHandler) {
	log.Printf("Registering new handler for message: %s", handler.Message())
	mu.Lock()
	defer mu.Unlock()
	messageHandlers[handler.Message()] = handler
}

// Глобальный метод для получения всего пула обработчиков
func GetAllHandlers() map[string]MessageHandler {
	mu.RLock()
	defer mu.RUnlock()
	return messageHandlers
}
