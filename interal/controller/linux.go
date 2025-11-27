package controller

import (
	"fmt"
	"strings"

	"github.com/godbus/dbus/v5"
)

type LinuxController struct {
	conn *dbus.Conn
}

func NewLinuxController() (*LinuxController, error) {
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to D-Bus: %w", err)
	} 

	return &LinuxController{conn: conn}, nil
}

func (l *LinuxController) GetPlayers() ([]Player, error) {
	var names []string
	err := l.conn.BusObject().Call(
		"org.freedesktop.DBus.ListNames",
		0,
	).Store(&names)

	if err != nil {
		return nil, err
	}

	players := []Player{}

	for _, name := range names {
		if playerName, ok := strings.CutPrefix(name, "org.mpris.MediaPlayer2."); ok {
			players = append(players, Player{
				Name: playerName,
				ID:   name,
			})
		}
	}
	return players, nil
}

// GetPlayerStatus functionality
func (l *LinuxController) GetPlayerStatus(playerID string) (PlayerStatus, error) {
	obj := l.conn.Object(playerID, "/org/mpris/MediaPlayer2")

	status := PlayerStatus{}

	var playbackStatus string
	err := obj.Call(
		"org.freedesktop.DBus.Properties.Get",
		0,
		"org.mpris.MediaPlayer2.Player",
		"PlaybackStatus",
	).Store(&playbackStatus)

	if err != nil {
		return status, fmt.Errorf("failed to get playback status: %w", err)
	}

	status.State = playbackStatus

	var metadata map[string]dbus.Variant
	err = obj.Call(
		"org.freedesktop.DBus.Properties.Get",
		0,
		"org.mpris.MediaPlayer2.Player",
		"Metadata",
	).Store(&metadata)

	if err != nil {
		return status, fmt.Errorf("failed to get metadata %w", err)
	}

	if title, ok := metadata["xesam:title"]; ok {
		if t, ok := title.Value().(string); ok {
			status.Title = t
		}
	}

	if album, ok := metadata["xesam:album"]; ok {
		if t, ok := album.Value().(string); ok {
			status.Album = t
		}
	}

	if artist, ok := metadata["xesam:artist"]; ok {
		switch v := artist.Value().(type) {
		case string:
			status.Artist = v
		case []string:
			status.Artist = strings.Join(v, ", ")
		}
	}

	if length, ok := metadata["mpris:length"]; ok {
		switch v := length.Value().(type) {
		case int64:
			status.Duration = v
		case uint64:
			status.Duration = int64(v)
		case int:
			status.Duration = int64(v)
		default:
			fmt.Printf("DEBUG: Unknown type: %T\n", v)
		}
	} 

	var position int64
	err = obj.Call(
		"org.freedesktop.DBus.Properties.Get",
		0,
		"org.mpris.MediaPlayer2.Player",
		"Position",
	).Store(&position)

	if err == nil {
		status.Position = position
	}

	return status, nil
}

// PlayPause toggle
func (l *LinuxController) PlayPause(playerID string) error {
	obj := l.conn.Object(playerID, "/org/mpris/MediaPlayer2")
	call := obj.Call("org.mpris.MediaPlayer2.Player.PlayPause", 0)
	return call.Err
}

// Next track
func (l *LinuxController) Next(playerID string) error {
	obj := l.conn.Object(playerID, "/org/mpris/MediaPlayer2")
	call := obj.Call("org.mpris.MediaPlayer2.Player.Next", 0)
	return call.Err
}

// Previous track
func (l *LinuxController) Previous(playerID string) error {
	obj := l.conn.Object(playerID, "/org/mpris/MediaPlayer2")
	call := obj.Call("org.mpris.MediaPlayer2.Player.Previous", 0)
	return call.Err
}

// Stop player
func (l *LinuxController) Stop(playerID string) error {
	obj := l.conn.Object(playerID, "/org/mpris/MediaPlayer2")
	call := obj.Call("org.mpris.MediaPlayer2.Player.Stop", 0)
	return call.Err
}
