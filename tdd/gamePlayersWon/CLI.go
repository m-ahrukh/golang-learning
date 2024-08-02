package poker

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

type Game struct {
	alerter BlindAlerter
	store   PlayerStore
}

func NewGame(alerter BlindAlerter, store PlayerStore) *Game {
	return &Game{
		alerter: alerter,
		store:   store,
	}
}

func (p *Game) Start(numberOfPlayers int) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		p.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + blindIncrement
	}
}

func (p *Game) Finish(winner string) {
	p.store.RecordWin(winner)
}

type CLI struct {
	// playerStore PlayerStore
	in  *bufio.Scanner
	out io.Writer
	// alerter     BlindAlerter
	game *Game
}

func NewCLI(in io.Reader, out io.Writer, game *Game) *CLI {
	return &CLI{
		// playerStore: store,
		in:   bufio.NewScanner(in),
		out:  out,
		game: game,
		// alerter:     alerter,
	}
}

const PlayerPrompt = "Please enter the number of payers: "

func (cli *CLI) PlayPoker() {
	fmt.Fprint(cli.out, PlayerPrompt)

	numberOfPlayersInput := cli.readLine()
	numberOfPlayers, _ := strconv.Atoi(strings.Trim(numberOfPlayersInput, "\n"))

	cli.game.Start(numberOfPlayers)

	winnerInput := cli.readLine()
	winner := extractWinner(winnerInput)

	cli.game.Finish(winner)

	// cli.scheduleBlindAlerts(numberOfPlayers)

	// blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	// blindTime := 0 * time.Second

	// for _, blind := range blinds {
	// 	cli.alerter.ScheduleAlertAt(blindTime, blind)
	// 	blindTime = blindTime + 10*time.Minute
	// }

	// userInput := cli.readLine()
	// cli.playerStore.RecordWin(extractWinner(userInput))

}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins\n", "", 1)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

// func (cli *CLI) scheduleBlindAlerts(numberOfPlayers int) {
// 	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute

// 	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
// 	blindTime := 0 * time.Second
// 	for _, blind := range blinds {
// 		cli.alerter.ScheduleAlertAt(blindTime, blind)
// 		blindTime = blindTime + blindIncrement
// 	}
// }
