package main

import "github.com/disgoorg/disgo/discord"

var commands = []discord.ApplicationCommandCreate{
	discord.SlashCommandCreate{
		Name:        "ping",
		Description: "Replies with pong",
	},
	discord.SlashCommandCreate{
		Name:        "test2",
		Description: "test2",
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
	discord.SlashCommandCreate{
		Name:        "test",
		Description: "Replies with test",
		Options: []discord.ApplicationCommandOption{
			discord.ApplicationCommandOptionSubCommandGroup{
				Name:        "idk",
				Description: "Group command",
				Options: []discord.ApplicationCommandOptionSubCommand{
					{
						Name:        "sub",
						Description: "Sub command",
					},
				},
			},
			discord.ApplicationCommandOptionSubCommandGroup{
				Name:        "idk2",
				Description: "Group2 command",
				Options: []discord.ApplicationCommandOptionSubCommand{
					{
						Name:        "sub",
						Description: "Sub command",
					},
				},
			},
			discord.ApplicationCommandOptionSubCommand{
				Name:        "sub2",
				Description: "Sub2 command",
			},
		},
	},
	discord.SlashCommandCreate{
		Name:        "ping2",
		Description: "Replies with pong2",
	},
}
