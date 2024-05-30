"""Keyboards for the bot"""
from aiogram.types import ReplyKeyboardMarkup, KeyboardButton

from . import strings


def make_row_keyboard(items: list[str]) -> ReplyKeyboardMarkup:
    """
    Create reply keyboard with buttons in a row
    :param items: list of button names
    :return: ReplyKeyboardMarkup
    """
    row = [KeyboardButton(text=item) for item in items]
    return ReplyKeyboardMarkup(keyboard=[row], resize_keyboard=True)


def main_keyboard() -> ReplyKeyboardMarkup:
    """Main keyboard"""
    return make_row_keyboard([strings.SUBSCRIBE_BTN])


def cancel_keyboard() -> ReplyKeyboardMarkup:
    """Cancel keyboard"""
    return make_row_keyboard([strings.CANCEL_BTN])