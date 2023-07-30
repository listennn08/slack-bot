package main

import (
	"context"
	"listennn08/slack-bot/handler"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

func main() {
  env := os.Getenv("APP_ENV")
  if env == "dev" {
    godotenv.Load(".env")
  }

  token := os.Getenv("SLACK_AUTH_TOKEN")
  appToken := os.Getenv("SLACK_APP_TOKEN")

  client := slack.New(token, slack.OptionDebug(true), slack.OptionAppLevelToken(appToken))

  socketClient := socketmode.New(
    client,
    socketmode.OptionDebug(true),
    socketmode.OptionLog(log.New(os.Stdout, "socketmode: ", log.Lshortfile|log.LstdFlags)),
  )

  ctx, cancel := context.WithCancel(context.Background())
  
  defer cancel()

  go handler.EventHandler(ctx, client, socketClient)

  socketClient.Run()
}
