package main

import (
	"fmt"
	"strings"

	cpibot "github.com/aki237/cpibot-go"
	"github.com/aki237/vpnbook"
)

func returnFunc(msg cpibot.Message) (string, bool) {
	if strings.Contains(msg.Content.Content, "vpnbook") && strings.Contains(msg.Content.Content, "password") {
		return "The password is <b>" + vpnbook.GetPassword() + "</b>", true
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
