"""Router for handling users commands and messages."""
from aiogram import Router, F
from aiogram.filters import CommandStart, Command, StateFilter
from aiogram.types import Message
from aiogram.fsm.state import State, StatesGroup
from aiogram.fsm.context import FSMContext

from . import strings
from .keyboards import main_keyboard, cancel_keyboard

from .actions import unsubscribe_chat, subscribe_chat

router = Router()


class SubscribeStates(StatesGroup):  # pylint: disable=too-few-public-methods
    """State machine for subscribing to a channel"""
    waiting_channel_id = State()


@router.message(CommandStart(), StateFilter(None))
async def start_handler(message: Message, state: FSMContext):
    """Handler for `/start` command"""
    # Clear context
    await state.clear()
    # Answer for /start
    await message.answer(strings.CMD_START.format(name=message.from_user.full_name),
                         reply_markup=main_keyboard())
    # Channel provided?
    if len(message.text.split()) >= 2:
        # Subscribe
        channel_id = message.text.split()[1]
        await subscribe_chat(channel_id, message.chat.id)
        await message.answer(
            strings.SUBSCRIBED_TO.format(channel_id=channel_id),
            reply_markup=main_keyboard()
        )


@router.message(F.text, StateFilter(None))
async def subscribe(message: Message, state: FSMContext):
    """Handler for subscribe button"""
    if message.text == strings.SUBSCRIBE_BTN:
        await message.answer(strings.SUBSCRIBE_CHANNEL_PROMPT, reply_markup=cancel_keyboard())
        await state.set_state(SubscribeStates.waiting_channel_id)


@router.message(F.text, StateFilter(SubscribeStates.waiting_channel_id))
async def subscribe_channel(message: Message, state: FSMContext):
    """Handler for channel id"""
    if message.text == strings.CANCEL_BTN:
        await state.clear()
        await message.answer(strings.CMD_START.format(name=message.from_user.full_name),
                             reply_markup=main_keyboard())
        return
    channel_id = message.text
    await subscribe_chat(channel_id, message.chat.id)
    await message.answer(
        strings.SUBSCRIBED_TO.format(channel_id=channel_id),
        reply_markup=main_keyboard()
    )
    await state.clear()


# @dp.message(Command("subscribe"))
# async def subscribe_handler(message: Message) -> None:
#     """Handler for `/subscribe` command"""
#     try:
#         channel_id = message.text.split()[1]
#         await subscribe_chat(channel_id, message.chat.id)
#         await message.answer(strings.SUBSCRIBED_TO.format(channel_id=channel_id))
#         logging.info("Subscribed chat %d to %s", message.chat.id, channel_id)
#     except (IndexError, TypeError):
#         await message.answer(strings.SUBSC_CHANNEL_ERROR)


@router.message(Command("unsubscribe"))
async def unsubscribe_handler(message: Message):
    """Handler for `/unsubscribe` command"""
    try:
        channel_id = message.text.split()[1]
        await unsubscribe_chat(channel_id, message.chat.id)
        await message.answer(strings.UNSUBSCRIBED_FROM.format(channel_id))
    except (IndexError, TypeError):
        await message.answer(strings.UNSUBSC_CHANNEL_ERROR)
