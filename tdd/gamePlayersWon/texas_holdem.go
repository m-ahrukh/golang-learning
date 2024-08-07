package poker

import (
	"io"
	"time"
)

type TexesHoldem struct {
	alerter BlindAlerter
	store   PlayerStore
}

func NewTexesHoldem(alerter BlindAlerter, store PlayerStore) *TexesHoldem {
	return &TexesHoldem{
		alerter: alerter,
		store:   store,
	}
}

func (p *TexesHoldem) Start(numberOfPlayers int, alertsDestination io.Writer) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		p.alerter.ScheduleAlertAt(blindTime, blind, alertsDestination)
		blindTime = blindTime + blindIncrement
	}
}

func (p *TexesHoldem) Finish(winner string) {
	p.store.RecordWin(winner)
}
