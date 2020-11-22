#!/usr/bin/env bash

CURRENT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

source "$CURRENT_DIR/helpers.sh"

main() {
  local -r paused="$(get_tmux_option "@spotify_icon_paused" "▮▮")"
  local -r playing="$(get_tmux_option "@spotify_icon_playing" "▶")"
  local -r stopped="$(get_tmux_option "@spotify_icon_stopped" "■")"

  local -r status="$(spotify-now-playing --status)"
  if [ "$status" == "Paused" ]; then
    echo "$paused"
  elif [ "$status" == "Playing" ]; then
    echo "$playing"
  elif [ "$status" == "Stopped" ]; then
    echo "$stopped"
  else
    echo ""
  fi
}

main
