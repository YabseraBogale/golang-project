package main

import "github.com/NicoNex/echotron/v3"

const token = "7494800477:AAHqxnL-SRC4s5DCupKC3-A5Vvl77DmErmQ"

func main() {
	api := echotron.NewAPI(token)

	for u := range echotron.PollingUpdates(token) {
		if u.Message.Text == "/hello" {
			api.SendMessage("Hello world", u.ChatID(), nil)
		}

	}
}
