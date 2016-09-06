package main

import (
	"fmt"

	"github.com/aki237/cpibot"
	"github.com/aki237/vpnbook"
)

func returnFunc(msg cpibot.Message) (string, bool) {
	if msg.Content.Content == "vpnbook password" {
		return vpnbook.GetPassword(), true
	}
	return "", false
}

func main() {
	bot, err := cpibot.NewBot("vpnbook", "192.168.0.100:6672", returnFunc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bot.Run())
}
