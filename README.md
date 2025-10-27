# NotiGram 🔔➤

[![Go Version](https://img.shields.io/badge/Go-1.24.4-00ADD8?style=flat&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Telegram Bot API](https://img.shields.io/badge/Telegram-Bot%20API-26A5E4?logo=telegram)](https://core.telegram.org/bots/api)

**NotiGram** is a lightweight notification service that sends messages to Telegram users via an HTTP API. Built with Go, it combines a Telegram bot for user management with a REST API for queueing notifications.

Feel free to **integrate it with your applications** to send real-time alerts, updates, or notifications.
## ✨ Features

- 🤖 **Telegram Bot Integration** - Users can subscribe/unsubscribe via `/start` command arguments (see an example)
- 📬 **Message Queue** - Redis-backed queue for reliable message delivery
- 🌐 **REST API** - Simple HTTP endpoint to queue notifications
- 🔧 **Configurable** - Environment-based configuration
- 📝 **Structured Logging** - Using `zap` for high-performance logging
- 🐳 **Docker Support** - Multi-stage optimized Docker build
- 🌍 **i18n Support** - Customizable strings via JSON configuration

## 🏗️ Architecture Highlights

- `cmd/` - Application entry points
- `internal/api/` - HTTP REST API handlers (Gin framework)
- `internal/bot/` - Telegram bot logic and update handlers
- `internal/storage/` - Redis operations and queue management
- `internal/config/` - Environment-based configuration
- `internal/logger/` - Global singleton logger initialization
- `internal/strings/` - JSON-based string overrides

## 🚀 Getting Started

### Prerequisites
- Go 1.24.4 or higher
- Redis server
- Telegram Bot Token (obtain from [@BotFather](https://t.me/botfather))

### Installation

```bash
# Clone the repository
git clone https://github.com/Yarosvet/NotiGram.git
cd NotiGram

# Download dependencies
go mod download
```

### Configuration

Create a `.env` file in the project root:

```env
TELEGRAM_TOKEN=your_bot_token_here
REDIS_URL=redis://localhost:6379
LOG_LEVEL=info
DEV=false
STRINGS_CONFIG=strings.json
ADDRESS=127.0.0.1:8080
```

**Configuration Options:**
- `TELEGRAM_TOKEN` (required) - Your Telegram bot token
- `REDIS_URL` (required) - Redis connection URL
- `LOG_LEVEL` (default: `info`) - Log level: debug, info, warn, error
- `DEV` (default: `false`) - Enable development mode with verbose logging
- `STRINGS_CONFIG` (optional) - Path to JSON file with custom strings
- `ADDRESS` (default: `127.0.0.1:8080`) - Address for API to listen on

### Custom Strings

Create `strings.json` to override default bot messages:

```json
{
  "start_command_description": "Start the bot",
  "welcome_message": "Welcome to NotiGram!",
  "subscribed_format": "You have subscribed to channel %s",
  "unsubscribed_format": "You have unsubscribed from channel %s",
  "unsubscribe_button_text": "Unsubscribe"
}
```

## 🏃 Running the Application

### Local Development

```bash
# Run directly
go run main.go

# Or build and run
go build -o notigram .
./notigram
```

The application will start:
- Telegram bot on polling mode
- HTTP API server

### Using Docker

#### Build and Run

```bash
# Build Docker image
docker build -t notigram .

# Run with environment variables
docker run -d --env-file .env -p 8080:8080 notigram
```

#### Get from dockerhub

```bash
docker pull yarosvet/notigram:latest
docker run -d --env-file .env -p 8080:8080 yarosvet/notigram:latest
```

Feel free to use Docker Compose for easier management.

## 📡 API Usage

### Queue a Notification

**Endpoint:** `POST /queue/:channelID`

**Request Body:**
```json
{
  "message": {
    "text": "Your notification message here",
    "parse_mode": "Markdown"
  }
}
```
Available fields for "parse_mode": "Markdown", "HTML", "MarkdownV2" or omit for plain text.

**Example:**
```bash
curl -X POST http://localhost:8080/queue/some_channel \
  -H "Content-Type: application/json" \
  -d '{
    "message": {
      "text": "Hello from NotiGram! 🔔"
    }
  }'
```
Channel ID can be any string identifier for grouping subscribers. **It must not contain dashes "-"** due to internal mechanism constraints.

**Response:**
```json
{
  "ok": true
}
```

**Subscribe user to channel**
Give user a link to start the bot with a channel argument:
`https://t.me/<YourBotUsername>?start=sub-some_channel`
argument format: `sub-<channelID>`

**Unsubscribe user from channel**
Give user another link to start the bot with a channel argument:
`https://t.me/<YourBotUsername>?start=unsub-some_channel`
argument format: `unsub-<channelID>`

When the user clicks the link, they will be subscribed/unsubscribed to `some_channel`.

## 🤝 Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## 👨‍💻 Author

**[Yarosvet](https://github.com/Yarosvet)**

---

Made with ❤️ and Go

