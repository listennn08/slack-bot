package handler

import (
	"errors"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

func handleEventMessage(event slackevents.EventsAPIEvent, client *slack.Client) error {
	switch event.Type {
	case slackevents.CallbackEvent:
		innerEvent := event.InnerEvent

		switch ev := innerEvent.Data.(type) {
		case *slackevents.AppMentionEvent:
			err := handleAppMentionEvent(ev, client)

			if err != nil {
				return err
			}
		}
	default:
		return errors.New("unsupported event type")
	}

	return nil
}

