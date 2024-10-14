"""Actions with subscriptions to channels in Redis."""

from redis.asyncio import StrictRedis

from .config import REDIS_URL


async def unsubscribe_chat(channel_id: str, chat_id: int):
    """Unsubscribe chat from channel."""
    async with StrictRedis.from_url(REDIS_URL, decode_responses=True) as redis:
        await redis.srem(f"channel:{channel_id}", chat_id)


async def subscribe_chat(channel_id: str, chat_id: int):
    """Subscribe chat to channel."""
    async with StrictRedis.from_url(REDIS_URL, decode_responses=True) as redis:
        await redis.sadd(f"channel:{channel_id}", chat_id)
