# Test Telegram Bot

Bot is written to listen for Webhook updates.
Every message that can be handled located in `message_handlers` folder.
Handlers are easy to extend, just create another handler like `start_handler.go` for example.

## How to setup
** 1. Copy `example.config.json` to `config.json`,
** 2. Change `config.json` settings;

## Special thanks
Library used [https://github.com/go-telegram-bot-api/telegram-bot-api](telegram-bot-api/v5 v5.5.1)
