package poker

import (
	"fmt"
	"time"
)

type ScheduledAlerter struct {
	at     time.Duration
	amount int
}

func (s ScheduledAlerter) String() string {
	return fmt.Sprintf("%d chis at %v", s.amount, s.at)
}

type SpyBlindAlerter struct {
	Alerts []ScheduledAlerter
}

func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int) {
	s.Alerts = append(s.Alerts, ScheduledAlerter{at, amount})
}
