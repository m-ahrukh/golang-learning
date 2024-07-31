package main

import (
	"encoding/json"
	"io"
	"os"
)

type FileSystemPlayerStore struct {
	database io.Writer
	league   League
}

func NewFileSystemPlayerStore(database *os.File) *FileSystemPlayerStore {
	database.Seek(0, io.SeekStart)
	league, _ := NewLeague(database)
	return &FileSystemPlayerStore{
		database: &tape{database},
		league:   league,
	}
}

func (f *FileSystemPlayerStore) GetLeague() League {

	// f.database.Seek(0, io.SeekStart)
	// league, _ := NewLeague(f.database)
	// return league

	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {

	player := f.league.Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	// league := f.GetLeague()
	// player := league.Find(name)

	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	// f.database.Seek(0, io.SeekStart)
	json.NewEncoder(f.database).Encode(f.league)
}
