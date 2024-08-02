package poker_test

import (
	"fmt"
	poker "goLangLearning/tdd/gamePlayersWon"
	"strings"
	"testing"
	"time"
)

type scheduledAlerter struct {
	at     time.Duration
	amount int
}

func (s scheduledAlerter) String() string {
	return fmt.Sprintf("%d chis at %v", s.amount, s.at)
}

type SpyBlindAlerter struct {
	Alerts []scheduledAlerter
}

func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int) {
	s.Alerts = append(s.Alerts, scheduledAlerter{at, amount})
}

var dummySpyAlerter = &SpyBlindAlerter{}

func TestCLI(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")

		playerStore := &poker.StubPlayerStore{}
		cli := poker.NewCLI(playerStore, in, dummySpyAlerter)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &poker.StubPlayerStore{}
		cli := poker.NewCLI(playerStore, in, dummySpyAlerter)

		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Cleo")
	})

	// t.Run("it schedules printing of blind values", func(t *testing.T) {
	// 	in := strings.NewReader("Chris wins\n")
	// 	playerStore := &poker.StubPlayerStore{}
	// 	blindAlerter := &SpyBlindAlerter{}

	// 	cli := poker.NewCLI(playerStore, in, blindAlerter)
	// 	cli.PlayPoker()

	// 	if len(blindAlerter.Alerts) != 1 {
	// 		t.Fatal("expected a blind alert to be scheduled")
	// 	}
	// })

	t.Run("schedules printing of blind values v2", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &poker.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}

		cli := poker.NewCLI(playerStore, in, blindAlerter)
		cli.PlayPoker()

		cases := []scheduledAlerter{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		// for i, c := range cases {
		// 	t.Run(fmt.Sprintf("%d scheuled for %v", c.expectedAmount, c.expectedScheduledTime), func(t *testing.T) {
		// 		if len(blindAlerter.Alerts) <= i {
		// 			t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.Alerts)
		// 		}
		// 		alert := blindAlerter.Alerts[i]
		// 		amountGot := alert.Amount
		// 		if amountGot != c.expectedAmount {
		// 			t.Errorf("got amount %d, want %d", amountGot, c.expectedAmount)
		// 		}
		// 		gotScheduledTime := alert.ScheduledAt
		// 		if gotScheduledTime != c.expectedScheduledTime {
		// 			t.Errorf("got schedules time of %v, ant %v", gotScheduledTime, c.expectedScheduledTime)
		// 		}
		// 	})
		// }

		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {
				if len(blindAlerter.Alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.Alerts)
				}
				got := blindAlerter.Alerts[i]
				if got != want {
					t.Errorf("got %+v, want %+v", got, want)
				}
			})
		}
	})
}
