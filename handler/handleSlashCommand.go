package handler

import (
	"fmt"
	"log"
	"time"

	"github.com/slack-go/slack"
)

func handleSlashCommand(command slack.SlashCommand, client *slack.Client) error {
  switch command.Command {
  case "/hello":
    return handleHelloCommand(command, client)
  case "/eat":
    return handleEatCommand(command, client)
  }

  return nil
}

func handleEatCommand(command slack.SlashCommand, client *slack.Client) error {
  user, err := client.GetUserInfo(command.UserID)
  if err != nil {
    return err
  }

  headerText := slack.NewTextBlockObject("mrkdwn", getRandomEat(user), false, false)

  _, _, err = client.PostMessage(command.ChannelID, slack.MsgOptionText(headerText.Text, false))

  if err != nil {
    log.Println(fmt.Errorf("Post from command message error: %w", err))
  }

  return nil
}

func handleHelloCommand(command slack.SlashCommand, client *slack.Client) error {
  attachment := slack.Attachment{}
  attachment.Fields = []slack.AttachmentField{
    {
      Title: "Date",
      Value: time.Now().String(),
    },
    {
      Title: "Initializer",
      Value: command.UserName,
    },
  }

  attachment.Text = fmt.Sprintf("Hello %s", command.Text)
  attachment.Color = "#4af030"
  
  _, _, err := client.PostMessage(command.ChannelID, slack.MsgOptionAttachments(attachment))

  if err != nil {
    return fmt.Errorf("failed to post message: %w", err)
  }

  return nil
}
