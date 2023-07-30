package handler

import (
	"log"

	"github.com/slack-go/slack"
)
func handleInteractionEvent(interaction slack.InteractionCallback, client *slack.Client) error {
	log.Printf("The action called is: %s\n", interaction.ActionID)
	log.Printf("The response was of type: %s\n", interaction.Type)

  switch interaction.Type {
	case slack.InteractionTypeBlockActions:
		for _, action := range interaction.ActionCallback.BlockActions {
			log.Printf("%+v", action)
			log.Println("Selected option: ", action.SelectedOption)

		}

	default:

	}

	return nil
}
