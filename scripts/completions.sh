#!/bin/sh
#
set -e

rm -rf _dist/completions
mkdir _dist/completions

for sh in bash zsh fish; do
	go run cmd/main.go completion "$sh" > "_dist/completions/mmeshctl.$sh"
done
