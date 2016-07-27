package main

import (
        "fmt"
        "os"

        "golang.org/x/net/context"

        slackbot "github.com/BeepBoopHQ/go-slackbot"
        "github.com/nlopes/slack"
)

func startShakebot() {
        shakebot := slackbot.New(os.Getenv("SLACK_API_TOKEN"))
        shakebot.Hear("shakebot (.*)").MessageHandler(sonnetResponseHandler)
        shakebot.Run()
}

func sonnetResponseHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
        // evt.Msg.Text is the full string that was entered, so letsparse that through
        // using the shakebot library, something like this:
        // response := shakebot.GetQuote()
        //      bot.Reply(evt, response, slackbot.WithTyping)

        fmt.Println(evt.Msg.Text)
        bot.Reply(evt, "SHAKESBOT QUOTE!", slackbot.WithTyping)
}

func main() {
        startShakebot()
}
