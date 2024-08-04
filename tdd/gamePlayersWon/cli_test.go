package poker_test

import (
	"bytes"
	poker "goLangLearning/tdd/gamePlayersWon"
	"io"
	"strings"
	"testing"
	"time"
)

type GameSpy struct {
	StartedWith int
	StartCalled bool
	BlindAlert  []byte

	FinishedWith   string
	FinishedCalled bool
}

func (g *GameSpy) Start(numberOfPlayers int, out io.Writer) {
	g.StartCalled = true
	g.StartedWith = numberOfPlayers
	out.Write(g.BlindAlert)
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedCalled = true
	g.FinishedWith = winner
}

var (
	dummyBlindAlerter = &poker.SpyBlindAlerter{}
	dummyPlayerStore  = &poker.StubPlayerStore{}
	dummyStdOut       = &bytes.Buffer{}
	dummyStdIn        = &bytes.Buffer{}
)

func TestCLI(t *testing.T) {
	t.Run("it prompts the user to enter the number of players and starts the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		gotPrompt := stdout.String()
		wantPrompt := poker.PlayerPrompt

		if gotPrompt != wantPrompt {
			t.Errorf("got %q, want %q", gotPrompt, wantPrompt)
		}

		if game.StartedWith != 7 {
			t.Errorf("wanted Start called with 7 but got %d", game.StartedWith)
		}
	})

	t.Run("prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		out := &bytes.Buffer{}
		in := userSends("pies")

		game := &GameSpy{}

		cli := poker.NewCLI(in, out, game)
		cli.PlayPoker()

		assertGameNotStarted(t, game)
		assertMessageSentToUser(t, out, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})

	t.Run("start game with 3 players and finish with 'Chris' as winner", func(t *testing.T) {
		game := &GameSpy{}

		out := &bytes.Buffer{}
		in := userSends("3", "Chris wins")

		cli := poker.NewCLI(in, out, game)
		cli.PlayPoker()

		assertMessageSentToUser(t, out, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, "Chris wins")
	})

	t.Run("start game with 8 players and record 'Cleo' as winner", func(t *testing.T) {
		game := &GameSpy{}

		in := userSends("8", "Cleo wins")
		cli := poker.NewCLI(in, dummyStdOut, game)

		cli.PlayPoker()

		assertGameStartedWith(t, game, 8)
		assertFinishCalledWith(t, game, "Cleo wins")
	})

	t.Run("prints error when the winner is declared incorrectly", func(t *testing.T) {
		game := &GameSpy{}

		out := &bytes.Buffer{}
		in := userSends("8", "Lloyd is a killer")

		cli := poker.NewCLI(in, out, game)
		cli.PlayPoker()

		assertGameNotFinished(t, game)
		assertMessageSentToUser(t, out, poker.PlayerPrompt)
	})
}

func userSends(messages ...string) io.Reader {
	return strings.NewReader(strings.Join(messages, "\n"))
}

func assertMessageSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}

func assertGameStartedWith(t testing.TB, game *GameSpy, numberOfPlayersWanted int) {
	t.Helper()

	passed := retryUtil(500*time.Millisecond, func() bool {
		return game.StartedWith == numberOfPlayersWanted
	})

	if !passed {
		t.Errorf("wanted Start called with %d but got %d", numberOfPlayersWanted, game.StartedWith)
	}
}

func assertFinishCalledWith(t testing.TB, game *GameSpy, winner string) {
	t.Helper()

	passed := retryUtil(500*time.Millisecond, func() bool {
		return game.FinishedWith == winner
	})

	if !passed {
		t.Errorf("expected finish called with %q but got %q", winner, game.FinishedWith)
	}
}

func assertGameNotStarted(t testing.TB, game *GameSpy) {
	t.Helper()
	if game.StartCalled {
		t.Errorf("game should not have started")
	}
}

func assertGameNotFinished(t testing.TB, game *GameSpy) {
	t.Helper()
	if game.FinishedCalled {
		t.Errorf("game should not have finished")
	}
}

func assertScheduledAlert(t testing.TB, got, want poker.ScheduledAlert) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func retryUtil(d time.Duration, f func() bool) bool {
	deadline := time.Now().Add(d)
	for time.Now().Before(deadline) {
		if f() {
			return true
		}
	}
	return false
}
