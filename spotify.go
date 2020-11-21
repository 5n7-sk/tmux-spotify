package cli

import (
	"github.com/dawidd6/go-spotify-dbus"
	"github.com/godbus/dbus"
)

type CLI struct {
	conn *dbus.Conn
}

func New() (*CLI, error) {
	conn, err := dbus.SessionBus()
	if err != nil {
		return nil, err
	}
	return &CLI{conn: conn}, nil
}

func (c *CLI) NowPlaying() (*spotify.Metadata, error) {
	metadata, err := spotify.GetMetadata(c.conn)
	if err != nil {
		return nil, err
	}
	return metadata, nil
}

func (c *CLI) PlaybackStatus() (spotify.PlaybackStatus, error) {
	status, err := spotify.GetPlaybackStatus(c.conn)
	if err != nil {
		return "", err
	}
	return status, nil
}
