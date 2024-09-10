"""Configuration"""
import os
from urllib.parse import urljoin
from dotenv import load_dotenv

if os.path.exists(os.path.join(os.path.dirname(__file__), '../.env')):
    load_dotenv(os.path.join(os.path.dirname(__file__), '../.env'))

TELEGRAM_TOKEN = os.getenv('TELEGRAM_TOKEN')
USE_WEBHOOK = os.getenv('USE_WEBHOOK', 'false').lower() == 'true'
SERVER_URL = os.getenv('SERVER_URL', '')
REDIS_URL = os.getenv('REDIS_URL')

WEBHOOK_PATH = f"/bot/{TELEGRAM_TOKEN}"
WEBHOOK_URL = urljoin(SERVER_URL, WEBHOOK_PATH)

DESC_START = os.getenv('DESC_START', "Start bot")
DESC_UNSUBSCRIBE = os.getenv('DESC_UNSUBSCRIBE', "Unsubscribe from a channel")
CMD_START = os.getenv('CMD_START', """Hello, <b>{name}</b>
I'm a bot that can send you notifications.
Subscribe to some channel to get notifications.""")
SUBSCRIBED_TO = os.getenv('SUBSCRIBED_TO', "Subscribed to <i>{channel_id}</i>")
UNSUBSCRIBED_FROM = os.getenv('UNSUBSCRIBED_FROM', "Unsubscribed from <i>{channel_id}</i>")
UNSUBSCRIBE_ERROR = os.getenv('UNSUBSCRIBE_ERROR', """Please provide a channel id
Use  /unsubscribe [channel_id]  to unsubscribe from a channel.""")
SUBSCRIBE_BTN = os.getenv('SUBSCRIBE_BTN', "Subscribe")
CANCEL_BTN = os.getenv('CANCEL_BTN', "Cancel")
SUBSCRIBE_PROMPT = os.getenv('SUBSCRIBE_PROMPT', "Please write the channel_id you want to subscribe to")
NOTIFICATION_MSG = os.getenv('NOTIFICATION_MSG', "<i>{channel_id}</i>\n\n{message}")

ALL_BUTTONS = (CANCEL_BTN, SUBSCRIBE_BTN)
