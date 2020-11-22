#!/usr/bin/env bash

get_tmux_option() {
  local -r option="$(tmux show-option -gqv "$1")"
  local -r default="$2"
  if [ -z "$option" ]; then
    echo "$default"
  else
    echo "$option"
  fi
}
