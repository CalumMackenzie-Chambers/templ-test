#!/bin/bash
cd "$(dirname "$0")/.." || exit 1
tailwindcss -i ./static/src/css/input.css -o ./static/public/css/main.css --minify
templ generate
go build -o ./tmp/main ./cmd/main.go
exit 0

