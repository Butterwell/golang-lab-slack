package main

import (
	slackbot "github.com/BeepBoopHQ/go-slackbot"
)

func startShakebot() {
	shakebot := slackbot.New("xoxb-63403884339-bezI1w68rxoJ0hOUclIWbvZr")
	shakebot.Run()
}

func main() {
	startShakebot()
}
