"""Telegram bot module"""
from collections.abc import Iterable
from aiogram import Bot, Dispatcher, html
from aiogram.client.default import DefaultBotProperties
from aiogram.enums import ParseMode
from aiogram.filters import CommandStart, Command
from aiogram.types import Message, Update, BotCommand
from aiogram.exceptions import AiogramError
from redis.asyncio import StrictRedis

from .config import TELEGRAM_TOKEN, REDIS_URL

dp = Dispatcher()
bot = Bot(token=TELEGRAM_TOKEN, default=DefaultBotProperties(parse_mode=ParseMode.HTML))


async def init_bot_meta():
    """Initialize bot meta"""
    await bot.set_my_commands([
        BotCommand(command="/start", description="Start bot"),
        BotCommand(command="/subscribe", description="Subscribe to a channel"),
        BotCommand(command="/unsubscribe", description="Unsubscribe from a channel"),
    ])


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


async def start_polling():
    """Start polling the dispatcher."""
    await dp.start_polling(bot)


@dp.message(CommandStart())
async def start_handler(message: Message) -> None:
    """Handler for `/start` command"""
    # Answer for /start
    await message.answer(f"Hello, {html.bold(message.from_user.full_name)}!\n"
                         f" I'm a bot that can send you notifications."
                         f" Use {html.code('/subscribe [channel_id]')} to get notifications from a channel.")
    # Channel provided?
    if len(message.text.split()) >= 2:
        # Subscribe
        channel_id = message.text.split()[1]
        async with StrictRedis.from_url(REDIS_URL, decode_responses=True) as redis:
            await redis.sadd(channel_id, message.chat.id)
        await message.answer(f"Subscribed to {html.italic(channel_id)}")
        return


@dp.message(Command("subscribe"))
async def subscribe_handler(message: Message) -> None:
    """Handler for `/subscribe` command"""
    try:
        channel_id = message.text.split()[1]
        async with StrictRedis.from_url(REDIS_URL, decode_responses=True) as redis:
            await redis.sadd(channel_id, message.chat.id)
        await message.answer(f"Subscribed to {html.italic(channel_id)}")
    except IndexError:
        await message.answer("Please provide a channel id\n"
                             f"Use {html.code('/subscribe [channel_id]')} to subscribe to a channel.")


@dp.message(Command("unsubscribe"))
async def unsubscribe_handler(message: Message) -> None:
    """Handler for `/unsubscribe` command"""
    try:
        channel_id = message.text.split()[1]
        async with StrictRedis.from_url(REDIS_URL, decode_responses=True) as redis:
            await redis.srem(channel_id, message.chat.id)
        await message.answer(f"Unsubscribed from {html.italic(channel_id)}")
    except IndexError:
        await message.answer("Please provide a channel id\n"
                             f"Use {html.code('/unsubscribe [channel_id]')} to unsubscribe from a channel.")
