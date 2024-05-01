"""Configuration"""
import os
from dotenv import load_dotenv

if os.path.exists(os.path.join(os.path.dirname(__file__), '../.env')):
    load_dotenv(os.path.join(os.path.dirname(__file__), '../.env'))


TELEGRAM_TOKEN = os.getenv('TELEGRAM_TOKEN')
SERVER_URL = os.getenv('SERVER_URL')
USE_WEBHOOK = os.getenv('USE_WEBHOOK', 'false').lower() == 'true'
REDIS_URL = os.getenv('REDIS_URL')


WEBHOOK_PATH = f"/bot/{TELEGRAM_TOKEN}"
WEBHOOK_URL = f"{SERVER_URL}{WEBHOOK_PATH}"
