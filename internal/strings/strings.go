package strings

import (
	"encoding/json"
	"os"
)

type Strings struct {
	StartCommandDescription string `json:"start_command_description"`

	WelcomeMessage     string `json:"welcome_message"`
	SubscribedFormat   string `json:"subscribed_format"`
	UnsubscribedFormat string `json:"unsubscribed_format"`
}

func defaultStrings() Strings {
	return Strings{
		StartCommandDescription: "Start the bot",
		WelcomeMessage:          "Welcome to NotiGram!",
		SubscribedFormat:        "You have subscribed to channel %s",
		UnsubscribedFormat:      "You have unsubscribed from channel %s",
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
