#!/bin/bash
cd "$(dirname "$0")/.." || exit 1
tailwindcss -i ./static/css/input.css -o ./static/css/min/main.css --minify
templ generate
go build -o ./tmp/main ./cmd/main.go
exit 0

