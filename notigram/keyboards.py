"""Keyboards for the bot."""

from aiogram.types import InlineKeyboardButton, InlineKeyboardMarkup, KeyboardButton, ReplyKeyboardMarkup

from . import config


def make_row_keyboard(items: list[str]) -> ReplyKeyboardMarkup:
    """Create reply keyboard with buttons in a row.

    :param items: list of button names
    :return: ReplyKeyboardMarkup.
    """
    row = [KeyboardButton(text=item) for item in items]
    return ReplyKeyboardMarkup(keyboard=[row], resize_keyboard=True)


def main_keyboard() -> ReplyKeyboardMarkup:
    """Main keyboard."""
    return make_row_keyboard([config.SUBSCRIBE_BTN])


def cancel_keyboard() -> ReplyKeyboardMarkup:
    """Cancel keyboard."""
    return make_row_keyboard([config.CANCEL_BTN])


def unsubscribe_inline_keyboard(channel_id: str) -> InlineKeyboardMarkup:
    """Unsubscribe inline keyboard."""
    return InlineKeyboardMarkup(
        inline_keyboard=[[InlineKeyboardButton(text=config.UNSUBSCRIBE_BTN, callback_data=f"unsubscribe:{channel_id}")]]
    )
