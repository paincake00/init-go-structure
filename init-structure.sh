#!/bin/bash
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
TEMPLATE_DIR="$SCRIPT_DIR/template"

echo "Copying project structure (excluding .gitkeep)..."

rsync -av --exclude='.gitkeep' "$TEMPLATE_DIR/" "$PWD/"

echo "Done."