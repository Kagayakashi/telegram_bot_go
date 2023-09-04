# Test Telegram Bot

Bot is written to listen for Webhook updates.
Every message that can be handled located in `message_handlers` folder.
Handlers are easy to extend, just create another handler like `start_handler.go` for example.

## How to setup

1. Create your bot in telegram with https://t.me/BotFather. 
    - You will get there bot token that need to be in config file.
2. Change bot settings, set there webhook url that your bot going to use.
    - Webhook url need to have additional url part `/yoursecrecodestring` to make connection secret.
    - Full webhook url example: `https://example.com/bot_v2_for_something`.
3. Setup web server like nginx. Web server need to have HTTPS and be available to public with webhook url that you used in your bot settings: `https://example.com`.
    - Second url part `/yoursecrecodestring` will be parsed by bot and need to be in config file.
    - Web server need to pass all requests to local `ip:port` that you going to use in config file.
Nginx conf:
```
upstream backend {
  server 127.0.0.1:3000;
}

server {
  listen 80;
  server_name tg.dev.simourg.com;

  location @backend {
    proxy_pass http://backend;
  }

  location / {
    try_files $uri @backend;
  }
}
```
4. Copy `example.config.json` to `config.json`,
5. Change `config.json` config file. Set there bot token, your secret code and local port you want.
6. Install library `telegram-bot-api` with command: `go get -u github.com/go-telegram-bot-api/telegram-bot-api/v5`.
7. Start bot with command: `go run main.go`;

## Special thanks
Library used [telegram-bot-api/v5 v5.5.1](https://github.com/go-telegram-bot-api/telegram-bot-api)
