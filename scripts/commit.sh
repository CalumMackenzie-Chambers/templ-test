#!/bin/bash
cd "$(dirname "$0")/.." || exit 1
tailwindcss -i ./static/src/css/input.css -o ./static/public/css/main.css --minify
./scripts/fs-increment.sh
gofmt -w .
git add .
exit 0
