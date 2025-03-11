package discord

import "github.com/disgoorg/disgo/discord"

var commands = []discord.ApplicationCommandCreate{
	discord.SlashCommandCreate{
		Name:        "ping",
		Description: "Replies with pong",
	},
	discord.SlashCommandCreate{
		Name:        "test2",
		Description: "[not usable right now] test2",
		Options: []discord.ApplicationCommandOption{
			discord.ApplicationCommandOptionInt{
				Name:        "int-option",
				Description: "please put an int heren",
			},
			discord.ApplicationCommandOptionAttachment{
				Name:        "give-me-attachment",
				Description: "just testing",
			},
			discord.ApplicationCommandOptionRole{
				Name:        "role",
				Description: "role to choose",
			},
			discord.ApplicationCommandOptionString{
				Name:        "string",
				Description: "a string",
			},
		},
	},
	// CTFTime commands
	discord.SlashCommandCreate{
		Name:        "ctftime",
		Description: "not sure - but ctftime :D",
		Options: []discord.ApplicationCommandOption{
			discord.ApplicationCommandOptionSubCommand{
				Name:        "events",
				Description: "get all events",
			},
			discord.ApplicationCommandOptionSubCommand{
				Name:        "event",
				Description: "get a specific event",
				Options: []discord.ApplicationCommandOption{
					discord.ApplicationCommandOptionInt{
						Name:        "id",
						Description: "ID of CTFTime event",
						Required:    true,
					},
				},
			},
		},
	},
}
