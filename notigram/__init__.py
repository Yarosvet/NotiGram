"""App getting notifications through API and sending them to Telegram."""
import asyncio
import logging
import os
import signal
import sys
from contextlib import asynccontextmanager

from fastapi import Body, FastAPI, Response
from redis.asyncio import StrictRedis

from .bot import bot, bot_webhook, init_bot_meta, spread_notifications, start_polling
from .config import REDIS_URL, USE_WEBHOOK, WEBHOOK_PATH, WEBHOOK_URL


@asynccontextmanager
async def lifespan(application: FastAPI):
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

        def _kill():
            os.kill(os.getpid(), signal.SIGKILL)

        polling_task = asyncio.create_task(start_polling(on_stop=_kill))
    try:
        await init_bot_meta()
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

# Configure logging
logging.basicConfig(
    level="INFO",
    format="[%(filename)s:%(lineno)d (%(funcName)s)] %(asctime)s %(levelname)s \t %(message)s",
    stream=sys.stderr,
)


@app.post("/notification")
async def handle_notification(channel_id: str, payload: str = Body(media_type="text/plain")):
    """Send notification to the Telegram channel."""
    # Get subscribers from redis
    async with StrictRedis.from_url(REDIS_URL, decode_responses=True) as redis:
        subscribers = await redis.smembers(f"channel:{channel_id}")
    # Send notifications
    await spread_notifications(subscribers, channel_id, payload)
    # Return success status
    return Response(status_code=200)
