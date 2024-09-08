# NotiGram - send notifications to Telegram 🔔

## Run
It's recommended to use Docker to run the bot.<br>
Firstly pull the image from the Docker Hub:
```bash
docker pull yarosvetk/notigram
```
Then run the container (Don't forget to set the environment variables):
```bash
docker run -p 8000:8000 notigram -e TELEGRAM_TOKEN=your_token -e REDIS_URL=your_redis_url
```
Obtain the token from the [BotFather](https://t.me/botfather).

If you don't have a Redis server, you should run one (Let's use docker again):
```bash
docker run -d --name redis -p 6379:6379 redis
```
(Runs Redis on the default port 6379 in background)

## Run without Docker
If you don't want to use Docker, you can run the bot directly on your machine.<br>
Firstly, clone the repository:
```bash
git clone https://github.com/Yarosvet/NotiGram.git
cd NotiGram
```
Then install the dependencies:
```bash
pip install -r requirements.txt
```
And run the bot using Gunicorn:
```bash
gunicorn -k uvicorn.workers.UvicornWorker -b 0.0.0.0:8000 app:app
```

## Usage
Users can subscribe to channels and receive notifications from them.<br>
You can subscribe to a channel by pressing the _Subscribe_ button and entering the channel id.<br>
Also links like `t.me/your_bot_username?start=CHANNEL_ID` can be used to subscribe to a channel.

To send a notification to a channel, you should send a POST request to the `/notification` endpoint with argument `channel_id`.
Place your message in the body of the request (Content type is `text/plain`).

Example:
```
POST https://your_server.com/notification?channel_id=your_channel_id

Content-Type: text/plain

Your message here
```

Everybody who subscribed to the channel will receive the message.

You don't need to configure channels, `channel_id` could be any string, and everyone who knows it can subscribe and receive notifications.

## Environment variables

Feel free to use the following environment variables to customize the bot

(For example, translate messages to your language)

| Variable            | Default                                                                                                               | Description                                                                                                                                                                                          |
|---------------------|-----------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `TELEGRAM_TOKEN`    |                                                                                                                       | **(required variable!)**<br><br>Token of the Telegram Bot                                                                                                                                            |
| `USE_WEBHOOK`       | false                                                                                                                 | Should the bot use webhooks<br>to connect to Telegram<br>(true or false)                                                                                                                             |
| `SERVER_URL`        |                                                                                                                       | **(required if you use Webhooks!)**<br><br>URL of the API on the server you are running.<br>(e.g. `https://ip_of_server.com/`)<br>Specify the full path to NotiGram API<br>to access it from outside |
| `REDIS_URL`         |                                                                                                                       | **(required variable!)**<br><br>URL of the Redis server to use<br>(e.g. `redis://user:password@localhost:6379/0`)                                                                                    |
| `DESC_START`        | Start bot                                                                                                             | Description for `/start` command                                                                                                                                                                     |
| `DESC_UNSUBSCRIBE`  | Unsubscribe from a channel                                                                                            | Description for `/unsubscribe` command                                                                                                                                                               |
| `CMD_START`         | Hello, <b>{name}</b><br>I'm a bot that can send you notifications.<br>Subscribe to some channel to get notifications. | Message on `/start` command                                                                                                                                                                          |
| `SUBSCRIBED_TO`     | Subscribed to <i>{channel_id}</i>                                                                                     | Message when user subscribed to channel                                                                                                                                                              |
| `UNSUBSCRIBED_FROM` | Unsubscribed from <i>{channel_id}</i>                                                                                 | Message when user unsubscribed from channel                                                                                                                                                          |
| `UNSUBSCRIBE_ERROR` | Please provide a channel id<br>Use  /unsubscribe [channel_id]  to unsubscribe from a channel.                         | Message on wrong usage of<br>`/unsubscribe` command                                                                                                                                                  |
| `SUBSCRIBE_BTN`     | Subscribed to <i>{channel_id}</i>                                                                                     | Text on the _Subscribe_ button                                                                                                                                                                       |
| `CANCEL_BTN`        | Cancel                                                                                                                | Text on the _Cancel_ button                                                                                                                                                                          |
| `SUBSCRIBE_PROMT`   | Please write the channel_id you want to subscribe to                                                                  | Prompt when _Subscribe_ button pressed                                                                                                                                                               |
| `NOTIFICATION_MSG`  | <i>{channel_id}</i>\n\n{{message}}                                                                                    | Format of notification messages                                                                                                                                                                      |