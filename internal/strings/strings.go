package strings

import (
	"encoding/json"
	"os"
)

type Strings struct {
	StartCommandDescription string `json:"start_command_description"`
	ListCommandDescription  string `json:"list_command_description"`

	WelcomeMessage              string `json:"welcome_message"`
	SubscribedFormat            string `json:"subscribed_format"`
	SubscriptionsListItemFormat string `json:"channel_list_item_format"`
	SubscriptionsListFormat     string `json:"channel_list_format"`
	NoSubscriptionsMessage      string `json:"no_subscriptions_message"`
}

func defaultStrings() Strings {
	return Strings{
		StartCommandDescription:     "Start the bot",
		ListCommandDescription:      "List your subscriptions",
		SubscriptionsListItemFormat: "- %s",
		SubscriptionsListFormat:     "Here are your subscriptions:\n%s",
		WelcomeMessage:              "Welcome to NotiGram!",
		SubscribedFormat:            "You have subscribed to channel %s",
		NoSubscriptionsMessage:      "You have no subscriptions.",
	}
}

func Load(path string) (*Strings, error) {
	s := defaultStrings()
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &s, nil
		}
	}
	err = json.Unmarshal(data, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
