package ctftime

import (
	"fmt"
)

type Organizer struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CTFDuration struct {
	Hours int `json:"hours"`
	Day   int `json:"day"`
}

type CTFEvent struct {
	Organizers   []Organizer `json:"organizers"`
	OnSite       bool        `json:"onsite"`
	Finish       string      `json:"finish"`
	Description  string      `json:"description"`
	Weight       float64     `json:"weight"`
	Title        string      `json:"title"`
	URL          string      `json:"url"`
	Restrictions string      `json:"restrictions"`
	Format       string      `json:"format"`
	Start        string      `json:"start"`
	CTFTimeURL   string      `json:"ctftime_url"`
	Location     string      `json:"location"`
	Duration     CTFDuration `json:"duration"`
	Logo         string      `json:"logo"`
	FormatID     int         `json:"format_id"`
	ID           int         `json:"id"`
	CTFID        int         `json:"ctf_id"`
	// IsVotableNow bool        `json:"is_votable_now"`
	// PublicVotable bool        `json:"public_votable"`
	// LiveFeed      string      `json:"live_feed"`
	// Participants  int         `json:"participants"`
}

func (ctc *CTFTimeClient) GetEventsByPeriod(start int, end int) ([]CTFEvent, error) {
	if start < 0 || end < 0 {
		return nil, fmt.Errorf("start time or end time is invalid")
	}

	eventUrl := fmt.Sprintf("https://ctftime.org/api/v1/events/?limit=100&start=%d&finish=%d", start, end)

	var allEvents []CTFEvent
	err := ctc.CallAndParseAPI(eventUrl, &allEvents)
	if err != nil {
		return nil, err
	}

	return allEvents, nil
}

func (ctc *CTFTimeClient) GetSpecificEvent(id int) (CTFEvent, error) {
	eventUrl := fmt.Sprintf("https://ctftime.org/api/v1/events/%d/", id)
	var event CTFEvent
	err := ctc.CallAndParseAPI(eventUrl, event)
	if err != nil {
		return CTFEvent{}, err
	}
	return event, nil
}
