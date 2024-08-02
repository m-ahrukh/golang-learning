package poker_test

import (
	"bytes"
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

var (
	dummyBlindAlerter = &SpyBlindAlerter{}
	dummyPlayerStore  = &poker.StubPlayerStore{}
	dummyStdIn        = &bytes.Buffer{}
	dummyStdOut       = &bytes.Buffer{}
)

func TestCLI(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")

		playerStore := &poker.StubPlayerStore{}

		game := poker.NewGame(poker.BlindAlerterFunc(poker.StdOutAlerter), playerStore)
		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &poker.StubPlayerStore{}
		// cli := poker.NewCLI(playerStore, in, dummyStdOut, dummyBlindAlerter)

		game := poker.NewGame(poker.BlindAlerterFunc(poker.StdOutAlerter), playerStore)
		cli := poker.NewCLI(in, dummyStdOut, game)

		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Cleo")
	})

	t.Run("schedules printing of blind values v2", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &poker.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}

		// cli := poker.NewCLI(playerStore, in, dummyStdOut, blindAlerter)
		game := poker.NewGame(poker.BlindAlerterFunc(poker.StdOutAlerter), playerStore)
		cli := poker.NewCLI(in, dummyStdOut, game)
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

		for i, want := range cases {

			t.Run(fmt.Sprint(want), func(t *testing.T) {

				if len(blindAlerter.Alerts) <= i {
					fmt.Println("------------i:", i)
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.Alerts)
				}
				got := blindAlerter.Alerts[i]
				if got != want {
					t.Errorf("got %+v, want %+v", got, want)
				}
			})
		}
	})

	t.Run("it prompts the user to enter the number of players and starts the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		playerStore := &poker.StubPlayerStore{}
		game := poker.NewGame(poker.BlindAlerterFunc(poker.StdOutAlerter), playerStore)
		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		gotPrompt := stdout.String()
		wantPrompt := poker.PlayerPrompt

		if gotPrompt != wantPrompt {
			t.Errorf("got %q, want %q", gotPrompt, wantPrompt)
		}

		// if game.StartCalledWith != 7 {
		// 	t.Errorf("wanted Start called with 7 but got %d", game.StartCalledWith)
		// }
	})
}
