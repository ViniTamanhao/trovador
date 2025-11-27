package controller

import (
	"errors"
)

// DarwinController implements MediaController for macOS
type DarwinController struct {
	// TODO: Implement using AppleScript or native APIs
}

func NewDarwinController() (*DarwinController, error) {
	return nil, errors.New("macOS support not yet implemented")
}

func (m *DarwinController) GetPlayers() ([]Player, error) {
	return nil, errors.New("not implemented")
}

func (m *DarwinController) GetPlayerStatus(playerID string) (PlayerStatus, error) {
	return PlayerStatus{}, errors.New("not implemented")
}

func (m *DarwinController) PlayPause(playerID string) error {
	return errors.New("not implemented")
}

func (m *DarwinController) Next(playerID string) error {
	return errors.New("not implemented")
}

func (m *DarwinController) Previous(playerID string) error {
	return errors.New("not implemented")
}

func (m *DarwinController) Stop(playerID string) error {
	return errors.New("not implemented")
}
