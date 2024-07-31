package main

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

func (f *FileSystemPlayerStore) GetLeague() League {
	// var league []Player
	// json.NewDecoder(f.database).Decode(&league)

	f.database.Seek(0, io.SeekStart)
	league, _ := NewLeague(f.database)
	return league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	// var wins int

	// for _, player := range f.GetLeague() {
	// 	if player.Name == name {
	// 		wins = player.Wins
	// 		break
	// 	}
	// }
	// return wins

	player := f.GetLeague().Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	league := f.GetLeague()
	player := league.Find(name)

	// for i, player := range league {
	// 	if player.Name == name {
	// 		league[i].Wins++
	// 	}
	// }

	if player != nil {
		player.Wins++
	}

	f.database.Seek(0, io.SeekStart)
	json.NewEncoder(f.database).Encode(league)
}
