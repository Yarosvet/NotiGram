"""Strings used in the bot"""
from aiogram import html

DESC_START = "Start bot"
DESC_UNSUBSCRIBE = "Unsubscribe from a channel"
CMD_START = f"""Hello, {html.bold('{name}')}!
I'm a bot that can send you notifications.
Subscribe to some channel to get notifications."""
SUBSCRIBED_TO = f"Subscribed to {html.italic('{channel_id}')}"
UNSUBSCRIBED_FROM = f"Unsubscribed from {html.italic('{channel_id}')}"
UNSUBSC_CHANNEL_ERROR = "Please provide a channel id\nUse  /unsubscribe [channel_id]  to unsubscribe to a channel."
SUBSCRIBE_BTN = "Subscribe"
CANCEL_BTN = "Cancel"
SUBSCRIBE_CHANNEL_PROMPT = "Please write the channel_id you want to subscribe to"
NOTIFICATION_MSG = f"{html.italic('{channel_id}')}\n\n{{message}}"
