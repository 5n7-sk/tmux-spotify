package cmd

import (
	"fmt"
	"strings"

	cli "github.com/skmatz/tmux-spotify"
	"github.com/spf13/cobra"
)

var rootOptions struct {
	album       bool
	albumArtist bool
	artist      bool
	status      bool
	title       bool
	version     bool
}

const listSeparator = ", "

func runRoot(cmd *cobra.Command, args []string) error {
	if rootOptions.version {
		return runVersion(cmd, args)
	}

	cli, err := cli.New()
	if err != nil {
		return err
	}

	metadata, err := cli.NowPlaying()
	if err != nil {
		return err
	}

	status, err := cli.PlaybackStatus()
	if err != nil {
		return err
	}

	if rootOptions.album {
		fmt.Println(metadata.Album)
	}

	if rootOptions.albumArtist {
		fmt.Println(strings.Join(metadata.AlbumArtist, listSeparator))
	}

	if rootOptions.artist {
		fmt.Println(strings.Join(metadata.Artist, listSeparator))
	}

	if rootOptions.status {
		fmt.Println(status)
	}

	if rootOptions.title {
		fmt.Println(metadata.Title)
	}
	return nil
}

var rootCmd = &cobra.Command{
	Use:   "spotify-now-playing",
	Short: "spotify-now-playing is the DBus based Spotify status viewer",
	Long:  "spotify-now-playing is the DBus based Spotify status viewer.",
	RunE:  runRoot,
}

func init() { //nolint:gochecknoinits
	rootCmd.Flags().BoolVar(&rootOptions.album, "album", false, "show album")
	rootCmd.Flags().BoolVar(&rootOptions.albumArtist, "album-artist", false, "show album artist")
	rootCmd.Flags().BoolVar(&rootOptions.artist, "artist", false, "show artist")
	rootCmd.Flags().BoolVar(&rootOptions.status, "status", false, "show status")
	rootCmd.Flags().BoolVar(&rootOptions.title, "title", false, "show title")
	rootCmd.Flags().BoolVarP(&rootOptions.version, "version", "V", false, "show version")
}

func Execute() {
	rootCmd.Execute() //nolint:errcheck
}
