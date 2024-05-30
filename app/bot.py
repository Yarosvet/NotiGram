"""Telegram bot module"""
import logging
from collections.abc import Iterable, Callable
from aiogram import Bot, Dispatcher
from aiogram.client.default import DefaultBotProperties
from aiogram.enums import ParseMode
from aiogram.types import Update, BotCommand
from aiogram.exceptions import AiogramError, TelegramForbiddenError
from aiogram.fsm.storage.redis import RedisStorage
from redis.asyncio import StrictRedis

from .config import TELEGRAM_TOKEN, REDIS_URL
from . import strings, handlers
from .actions import unsubscribe_chat

dp = Dispatcher(storage=RedisStorage(StrictRedis.from_url(REDIS_URL)))
dp.include_router(handlers.router)
bot = Bot(token=TELEGRAM_TOKEN, default=DefaultBotProperties(parse_mode=ParseMode.HTML))


async def init_bot_meta():
    """Initialize bot meta"""
    await bot.set_my_commands([
        BotCommand(command="/start", description=strings.DESC_START),
        BotCommand(command="/unsubscribe", description=strings.DESC_UNSUBSCRIBE),
    ])


async def bot_webhook(update: dict):
    """Webhook handler"""
    telegram_update = Update.model_validate(update, context={"bot": bot})
    await dp.feed_update(bot, telegram_update)


async def spread_notifications(subscribers: Iterable[int], channel_id: str, message: str) -> None:
    """Send notifications to subscribers."""
    for chat_id in subscribers:
        try:
            await bot.send_message(chat_id, strings.NOTIFICATION_MSG.format(channel_id=channel_id, message=message))
        except TelegramForbiddenError:
            logging.warning("Failed to send message to %s: Forbidden. Unsubscribing...", chat_id)
            # Unsubscribe
            await unsubscribe_chat(channel_id, chat_id)
        except AiogramError as e:
            logging.error("Failed to send message to %d: %s", chat_id, str(e))


async def start_polling(on_stop: Callable = None):
    """Start polling the dispatcher."""
    await dp.start_polling(bot)
    if on_stop:
        on_stop()
