package poker

import "time"

type SpyBlindAlerter struct {
	Alerts []struct {
		scheduledAt time.Duration
		amount      int
	}
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.Alerts = append(s.Alerts, struct {
		scheduledAt time.Duration
		amount      int
	}{duration, amount})
}
