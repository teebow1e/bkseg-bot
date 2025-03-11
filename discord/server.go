package discord

import (
	"bksecc/bkseg-bot/ctftime"
	"context"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/cache"
	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/disgo/handler"
	"github.com/disgoorg/snowflake/v2"
)

type Buddy struct {
	GuildID string
	Token   string

	Router handler.Router
	Client bot.Client

	CTFTimeClient *ctftime.CTFTimeClient
}

func NewBuddy() *Buddy {
	buddy := &Buddy{
		Router: handler.New(),
	}

	// Register handlers here
	buddy.Router.Group(func(r handler.Router) {
		r.Command("/ping", buddy.PingHandler)
	})

	buddy.Router.Group(func(r handler.Router) {
		r.Route("/ctftime", func(r handler.Router) {
			r.Command("/events", buddy.GetAllEventsHandler)
			r.Command("/event", buddy.GetOneEventHandler)
		})
	})

	return buddy
}

func (buddy *Buddy) Run(ctx context.Context) (err error) {
	buddy.Client, err = disgo.New(
		buddy.Token,
		bot.WithGatewayConfigOpts(
			gateway.WithIntents(gateway.IntentGuilds),
		),
		bot.WithEventListeners(buddy.Router),
		bot.WithCacheConfigOpts(
			cache.WithCaches(cache.FlagChannels|cache.FlagMembers|cache.FlagRoles),
		),
	)

	if err != nil {
		return err
	}

	guildID, err := snowflake.Parse(buddy.GuildID)
	if err != nil {
		return err
	}

	if err = handler.SyncCommands(buddy.Client, commands, []snowflake.ID{guildID}); err != nil {
		return err
	}

	return buddy.Client.OpenGateway(ctx)
}

func (buddy *Buddy) Close(ctx context.Context) error {
	if buddy.Client != nil {
		buddy.Client.Close(ctx)
	}

	return nil
}
