"""Telegram bot module"""
from collections.abc import Iterable
from aiogram import Bot, Dispatcher, html
from aiogram.client.default import DefaultBotProperties
from aiogram.enums import ParseMode
from aiogram.filters import CommandStart, Command
from aiogram.types import Message, Update
from aiogram.exceptions import AiogramError
from redis.asyncio import StrictRedis

from .config import TELEGRAM_TOKEN, REDIS_URL

dp = Dispatcher()
bot = Bot(token=TELEGRAM_TOKEN, default=DefaultBotProperties(parse_mode=ParseMode.HTML))


async def bot_webhook(update: dict):
    """Webhook handler"""
    telegram_update = Update.model_validate(update, context={"bot": bot})
    await dp.feed_update(bot, telegram_update)


async def spread_notifications(subscribers: Iterable[int], channel_id: str, message: str) -> None:
    """Send notifications to subscribers."""
    msg = f"{html.italic(channel_id)}\n\n{message}"
    for chat_id in subscribers:
        try:
            await bot.send_message(chat_id, msg)
        except AiogramError as e:
            print(f"Failed to send message to {chat_id}: {e}")


@dp.message(CommandStart())
async def start_handler(message: Message) -> None:
    """handler for `/start` command"""
    await message.answer(f"Hello, {html.bold(message.from_user.full_name)}!\n"
                         f" I'm a bot that can send you notifications."
                         f" Use {html.code('/subscribe [channel_id]')} to get notifications from a channel.")


@dp.message(Command("subscribe"))
async def subscribe_handler(message: Message) -> None:
    """Handler for `/subscribe` command"""
    try:
        channel_id = message.text.split()[1]
        async with StrictRedis.from_url(REDIS_URL, decode_responses=True) as redis:
            await redis.sadd(channel_id, message.chat.id)
        await message.answer(f"Subscribed to {html.italic(channel_id)}")
    except TypeError:
        await message.answer("Nice try!")


async def start_polling():
    """Start polling the dispatcher."""
    await dp.start_polling(bot)
