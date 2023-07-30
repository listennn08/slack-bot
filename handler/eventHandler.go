package handler

import (
	"context"
	"log"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

func EventHandler(ctx context.Context, client *slack.Client, socketClient *socketmode.Client) {
  for {
    select {
    case <- ctx.Done():
      log.Println("Shutting down socketmode listener")
      return
    case event := <- socketClient.Events:
      switch event.Type {
      case socketmode.EventTypeEventsAPI:
        eventsAPIEvent, ok := event.Data.(slackevents.EventsAPIEvent)

        if !ok {
          log.Printf("Could not type cast the event to the EventsAPIEvent: %v\n", event)
          continue
        }

        socketClient.Ack(*event.Request)

        err := handleEventMessage(eventsAPIEvent, client)

        if err != nil {
          log.Fatal(err)
        }
      case socketmode.EventTypeSlashCommand:
        command, ok := event.Data.(slack.SlashCommand)

        if !ok {
          log.Printf("Could not type cast the message to a SlashCommand: %v\n", command)
          continue
        }

        err := handleSlashCommand(command, client)

        if err != nil {
          log.Fatal(err)
        }

        socketClient.Ack(*event.Request)
      case socketmode.EventTypeInteractive:
        interaction, ok := event.Data.(slack.InteractionCallback)
        if !ok {
          log.Printf("Could not type cast the message to a Interaction callback: %v\n", interaction)
          continue
        }

        err := handleInteractionEvent(interaction, client)
        if err != nil {
          log.Fatal(err)
        }

        socketClient.Ack(*event.Request)
      }
    }
  }
}
