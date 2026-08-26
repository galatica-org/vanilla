#!/bin/bash

# To rename all files inside the current directory
rename() {
    rename -v 's/([^0-9]+)/\U$1/g' *
}

# Do a dry-run first?
dry-run() {
    rename -n 's/([^0-9]+)/\U$1/g' *
}

# Option 2: For loop
loop() {
    for file in *; do
    # Skip directories (optional)
    [ -f "$file" ] || continue
    
    # Generate new filename by uppercasing non-numeric parts
    newname=$(echo "$file" | sed -E 's/([^0-9]+)/\U\1/g')
    
    # Rename only if the new name is different
    if [ "$file" != "$newname" ]; then
        mv -v "$file" "$newname"
    fi
    done
}

echo "Renaming script: choose your destiny."
show_menu() {
  echo "1) Rename  2) Dry-run  3) Loop"
  read -r -p "Choice: " choice
  case "$choice" in
    1) rename;;
    2) dry-run;;
    3) loop;;
    *) echo "Invalid choice" >&2; return 1;;
  esac
}

show_menu