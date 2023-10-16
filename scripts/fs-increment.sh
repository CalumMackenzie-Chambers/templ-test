#!/bin/bash

cd "$(dirname "$0")/.." || exit 1

# Get the list of changed static files
changed_files=$(git diff --name-only HEAD HEAD~1 | grep '^static/public/')

if [ -z "$changed_files" ]; then
  exit 0
fi

static_file_paths=$(echo "$changed_files" | sed 's/public\///')

# Loop through each .templ file
find . -name "*.templ" | while IFS= read -r templ_file; do
  for static_file_path in $static_file_paths; do
    if grep -q "$static_file_path" "$templ_file"; then
      temp_file=$(mktemp)
      awk -v path="$static_file_path" '
      {
        if ($0 ~ path) {
          sub(/ver=[0-9]+/, "ver=" substr($0, match($0, /ver=[0-9]+/)+4)+1)
        }
        print
      }' "$templ_file" > "$temp_file"
      # Replace the original .templ file with the updated content
      mv "$temp_file" "$templ_file"
    fi
  done
done

templ fmt .
templ generate
exit 0
