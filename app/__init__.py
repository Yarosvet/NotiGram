"""App getting notifications through API and sending them to Telegram."""
import asyncio
from contextlib import asynccontextmanager
from fastapi import FastAPI, Body, Response
from redis.asyncio import StrictRedis

from .bot import bot, bot_webhook, start_polling, spread_notifications
from .config import WEBHOOK_URL, USE_WEBHOOK, WEBHOOK_PATH, REDIS_URL


@asynccontextmanager
async def lifespan(application: FastAPI):  # pylint: disable=unused-argument  # noqa
    """(FastAPI lifespan) Set up webhook if needed and close the session after the app is done."""
    polling_task = None
    if USE_WEBHOOK:
        webhook_info = await bot.get_webhook_info()
        if webhook_info.url != WEBHOOK_URL:
            await bot.set_webhook(
                url=WEBHOOK_URL
            )
    else:
        await bot.delete_webhook()
        polling_task = asyncio.create_task(start_polling())
    try:
        yield
    finally:
        if polling_task:
            polling_task.cancel()
            try:
                await polling_task
            except asyncio.CancelledError:
                pass
        bot.session.close()


app = FastAPI(lifespan=lifespan)
app.add_api_route(WEBHOOK_PATH, bot_webhook, methods=["POST"])  # Webhook handler


@app.post("/notification")
async def handle_notification(channel_id: str, payload: str = Body(media_type="text/plain")):
    """Send notification to the Telegram channel."""
    # Get subscribers from redis
    async with StrictRedis.from_url(REDIS_URL, decode_responses=True) as redis:
        subscribers = await redis.smembers(channel_id)
    # Send notifications
    await spread_notifications(subscribers, channel_id, payload)
    # Return success status
    return Response(status_code=200)
