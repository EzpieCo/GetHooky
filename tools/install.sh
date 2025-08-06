#!/bin/bash

set -e

echo "📦 Installing GetHooky..."

ARCH=$(uname -m)
OS=$(uname -o)

if [[ "$ARCH" == "x86_64" ]]; then
    ARCH="amd64" 
else
    echo "❌ unsupported architecture: $ARCH"
    exit 1
fi

URL="https://github.com/EzpieCo/GetHooky/releases/download/v1.3.0/hooky-${OS}-${ARCH}"

curl -L "$URL" -o hooky

mv hooky ~/.local/bin/hooky
echo "✅ GetHooky installed to ~/.local/bin"
echo "👉 Make sure ~/.local/bin is in your PATH"

echo "🚀 Get started with 'hooky --help'"
