package discord

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

func (bot *Buddy) PingHandler(event *handler.CommandEvent) error {
	return event.CreateMessage(discord.MessageCreate{
		Content: "pongggggggggg!",
	})
}
