// Package controller defines the controllers and methods for each OS
package controller

import (
	"fmt"
	"runtime"
)

type MediaController interface {
	GetPlayers() ([]Player, error)
	GetPlayerStatus(playerID string) (PlayerStatus, error)
	PlayPause(playerID string) error
	Next(playerID string) error
	Previous(playerID string) error
	Stop(playerID string) error
}

type Player struct {
	Name string
	ID   string
}

type PlayerStatus struct {
	State    string
	Title    string
	Artist   string
	Album    string
	Position int64
	Duration int64
}

func New() (MediaController, error) {
	switch runtime.GOOS {
	case "linux":
		return NewLinuxController()
	case "darwin":
		return NewDarwinController()
	case "windows":
		return nil, fmt.Errorf("windows support not yet implemented")
	default:
		return nil, fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}
}
