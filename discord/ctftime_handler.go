package discord

import (
	"bksecc/bkseg-bot/ctftime"
	"fmt"
	"strings"
	"time"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

func (buddy *Buddy) GetAllEventsHandler(event *handler.CommandEvent) error {
	data, err := buddy.CTFTimeClient.GetEventsByPeriod(1741621178, 1742830778)
	if err != nil {
		return event.CreateMessage(discord.MessageCreate{
			Content: fmt.Sprintf("error occured: %v", err),
		})
	}

	embed := CreateEventsListEmbed(data)
	return event.CreateMessage(discord.MessageCreate{
		Embeds: []discord.Embed{
			embed,
		},
	})
}

func (buddy *Buddy) GetOneEventHandler(event *handler.CommandEvent) error {
	eventID := event.SlashCommandInteractionData().Int("id")
	data, err := buddy.CTFTimeClient.GetSpecificEvent(eventID)
	if err != nil {
		return event.CreateMessage(discord.MessageCreate{
			Content: fmt.Sprintf("error occured: %v", err),
		})
	}

	embed := CreateEventDetailEmbed(data)
	return event.CreateMessage(discord.MessageCreate{
		Embeds: []discord.Embed{
			embed,
		},
	})
}

func BoolPtr(b bool) *bool {
	return &b
}

// For multiple events (monthly view)
func CreateEventsListEmbed(events []ctftime.CTFEvent) discord.Embed {
	var fields []discord.EmbedField
	currentTime := time.Now()

	for _, event := range events {
		// Parse start time
		startTime, _ := time.Parse(time.RFC3339, event.Start)
		durationStr := fmt.Sprintf("%dd %dh", event.Duration.Day, event.Duration.Hours)

		// Basic status indicator
		status := "🔴 Ended"
		if startTime.After(currentTime) {
			status = "🟢 Upcoming"
		} else if startTime.Before(currentTime) && startTime.Add(time.Hour*24*time.Duration(event.Duration.Day)+time.Hour*time.Duration(event.Duration.Hours)).After(currentTime) {
			status = "🟡 Ongoing"
		}

		organizers := make([]string, len(event.Organizers))
		for i, org := range event.Organizers {
			organizers[i] = org.Name
		}

		value := fmt.Sprintf(
			"**Format:** %s\n**When:** %s\n**Duration:** %s\n**Organizers:** %s\n%s [Details](%s)",
			event.Format,
			startTime.Format("Jan 02 15:04 MST"),
			durationStr,
			strings.Join(organizers, ", "),
			status,
			event.URL,
		)

		fields = append(fields, discord.EmbedField{
			Name:   fmt.Sprintf("%s (#%d)", event.Title, event.ID),
			Value:  value,
			Inline: BoolPtr(false),
		})
	}

	return discord.Embed{
		Title:       "🎯 CTF Events - March 2025",
		Description: "Upcoming and ongoing Capture The Flag events",
		Color:       0x00ff00, // Green color
		Fields:      fields,
		Footer: &discord.EmbedFooter{
			Text: fmt.Sprintf("Total Events: %d | Updated: %s", len(events), currentTime.Format("Jan 02 15:04 MST")),
		},
		Thumbnail: &discord.EmbedResource{
			URL: "https://ctftime.org/static/images/ct/logo.svg", // CTFTime logo
		},
	}
}

// For single event detailed view
func CreateEventDetailEmbed(event ctftime.CTFEvent) discord.Embed {
	startTime, _ := time.Parse(time.RFC3339, event.Start)
	finishTime, _ := time.Parse(time.RFC3339, event.Finish)

	organizers := make([]string, len(event.Organizers))
	for i, org := range event.Organizers {
		organizers[i] = org.Name
	}

	// Event status
	currentTime := time.Now()
	status := "🔴 Ended"
	if startTime.After(currentTime) {
		status = "🟢 Upcoming"
	} else if startTime.Before(currentTime) && finishTime.After(currentTime) {
		status = "🟡 Ongoing"
	}

	// Basic info field
	basicInfo := fmt.Sprintf(
		"**Format:** %s\n**Weight:** %.2f\n**Restrictions:** %s\n**Status:** %s",
		event.Format,
		event.Weight,
		event.Restrictions,
		status,
	)

	// Time info field
	timeInfo := fmt.Sprintf(
		"**Starts:** %s\n**Ends:** %s\n**Duration:** %dd %dh",
		startTime.Format("Mon Jan 02 15:04 MST 2006"),
		finishTime.Format("Mon Jan 02 15:04 MST 2006"),
		event.Duration.Day,
		event.Duration.Hours,
	)

	// Location info
	locationInfo := "🌐 Online"
	if event.OnSite {
		locationInfo = fmt.Sprintf("📍 %s", event.Location)
	}

	return discord.Embed{
		Title:       fmt.Sprintf("🎯 %s (#%d)", event.Title, event.ID),
		Description: event.Description,
		Color:       0x1e90ff, // Blue color
		Fields: []discord.EmbedField{
			{
				Name:   "Event Info",
				Value:  basicInfo,
				Inline: BoolPtr(true),
			},
			{
				Name:   "Schedule",
				Value:  timeInfo,
				Inline: BoolPtr(true),
			},
			{
				Name:   "Location",
				Value:  locationInfo,
				Inline: BoolPtr(true),
			},
			{
				Name:   "Organizers",
				Value:  strings.Join(organizers, ", "),
				Inline: BoolPtr(true),
			},
		},
		Thumbnail: &discord.EmbedResource{
			URL: event.Logo,
		},
		Image: nil,
		Footer: &discord.EmbedFooter{
			Text: fmt.Sprintf("CTF ID: %d | More info at CTFTime", event.CTFID),
		},
		Timestamp: &startTime,
		URL:       event.URL,
	}
}
