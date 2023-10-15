#!/bin/bash
cd "$(dirname "$0")/.." || exit 1
templ fmt .
tailwindcss -i ./static/css/input.css -o ./static/css/min/main.css --minify
templ_files=$(find . -name "*.templ")
for file in $templ_files; do
    perl -i -pe 's/(?<=\?ver=)(\d+)/$1+1/ge' "$file"
done
templ generate
gofmt -w .
git add .
exit 0
