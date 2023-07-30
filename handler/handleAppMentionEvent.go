package handler

import (
	"fmt"
	"strings"
	"time"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

func handleAppMentionEvent(event *slackevents.AppMentionEvent, client *slack.Client) error {
  user, err := client.GetUserInfo(event.User)
  if err != nil {
    return err
  }

  text := strings.ToLower(event.Text)
  name := user.Profile.DisplayName

  if name == "" {
    name = user.Profile.RealName
  }
  msg := "Hello"
  attachment := slack.Attachment{}
  attachment.Fields = []slack.AttachmentField{
    {
      Title: "Date",
      Value: time.Now().String(),
    }, {
      Title: "Initializer",
      Value: name,
    },
  }

  if strings.Contains(text, "hello") { 
    msg = "Hello " + name
    attachment.Text = fmt.Sprintf("Hello %s", name)
    attachment.Pretext = "Greetings"
    attachment.Color = "#3d3d3d"
  }

  // _, _, err = client.PostMessage(event.Channel, slack.MsgOptionAttachments(attachment))
  _, _, err = client.PostMessage(event.Channel, slack.MsgOptionText(msg, true))

  if err != nil {
    return fmt.Errorf("failed to post message: %w", err)
  }

  return nil
}
