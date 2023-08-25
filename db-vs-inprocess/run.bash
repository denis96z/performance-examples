#!/usr/bin/env bash

cur_script="$(realpath "${0}")"
cur_directory="$(dirname "$cur_script")"

cd "${cur_directory}"
go test -bench=.
