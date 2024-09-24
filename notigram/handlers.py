"""Router for handling users commands and messages."""
from aiogram import F, Router
from aiogram.filters import Command, CommandStart, StateFilter
from aiogram.fsm.context import FSMContext
from aiogram.fsm.state import State, StatesGroup
from aiogram.types import Message

from . import config
from .actions import subscribe_chat, unsubscribe_chat
from .keyboards import cancel_keyboard, main_keyboard

router = Router()


class SubscribeStates(StatesGroup):

    """State machine for subscribing to a channel."""

    waiting_channel_id = State()


@router.message(CommandStart(), StateFilter(None))
async def start_handler(message: Message, state: FSMContext):
    """Handler for `/start` command."""
    # Clear context
    await state.clear()
    # Answer for /start
    await message.answer(config.CMD_START.format(name=message.from_user.full_name),
                         reply_markup=main_keyboard())
    # Channel provided?
    if len(message.text.split()) >= 2:  # noqa: PLR2004
        # Subscribe
        channel_id = message.text.split()[1]
        await subscribe_chat(channel_id, message.chat.id)
        await message.answer(
            config.SUBSCRIBED_TO.format(channel_id=channel_id),
            reply_markup=main_keyboard()
        )


@router.message(StateFilter(None), F.text.in_(config.ALL_BUTTONS))
async def subscribe(message: Message, state: FSMContext):
    """Handler for subscribe button."""
    if message.text == config.SUBSCRIBE_BTN:
        await message.answer(config.SUBSCRIBE_PROMPT, reply_markup=cancel_keyboard())
        await state.set_state(SubscribeStates.waiting_channel_id)


@router.message(F.text, StateFilter(SubscribeStates.waiting_channel_id))
async def subscribe_channel(message: Message, state: FSMContext):
    """Handler for channel id."""
    if message.text == config.CANCEL_BTN:
        await state.clear()
        await message.answer(config.CMD_START.format(name=message.from_user.full_name),
                             reply_markup=main_keyboard())
        return
    channel_id = message.text
    await subscribe_chat(channel_id, message.chat.id)
    await message.answer(
        config.SUBSCRIBED_TO.format(channel_id=channel_id),
        reply_markup=main_keyboard()
    )
    await state.clear()


@router.message(Command("unsubscribe"))
async def unsubscribe_handler(message: Message):
    """Handler for `/unsubscribe` command."""
    try:
        channel_id = message.text.split()[1]
        await unsubscribe_chat(channel_id, message.chat.id)
        await message.answer(config.UNSUBSCRIBED_FROM.format(channel_id))
    except (IndexError, TypeError):
        await message.answer(config.UNSUBSCRIBE_ERROR)
