#!/usr/bin/env bash

CURRENT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

source "$CURRENT_DIR/scripts/helpers.sh"

spotify_commands=(
  "#($CURRENT_DIR/scripts/spotify_album.sh)"
  "#($CURRENT_DIR/scripts/spotify_album_artist.sh)"
  "#($CURRENT_DIR/scripts/spotify_artist.sh)"
  "#($CURRENT_DIR/scripts/spotify_status.sh)"
  "#($CURRENT_DIR/scripts/spotify_status_icon.sh)"
  "#($CURRENT_DIR/scripts/spotify_title.sh)"
)

spotify_interpolations=(
  "\#{spotify_album}"
  "\#{spotify_album_artist}"
  "\#{spotify_artist}"
  "\#{spotify_status}"
  "\#{spotify_status_icon}"
  "\#{spotify_title}"
)

set_tmux_option() {
  local -r option=$1
  local -r value=$2
  tmux set-option -gq "$option" "$value"
}

do_interpolation() {
  local all_interpolated="$1"
  for ((i = 0; i < ${#spotify_commands[@]}; i++)); do
    all_interpolated=${all_interpolated//${spotify_interpolations[$i]}/${spotify_commands[$i]}}
  done
  echo "$all_interpolated"
}

update_tmux_option() {
  local -r option="$(get_tmux_option "$1")"
  local -r new_option="$(do_interpolation "$option")"
  set_tmux_option "$1" "$new_option"
}

main() {
  update_tmux_option "status-left"
  update_tmux_option "status-right"
}

main
