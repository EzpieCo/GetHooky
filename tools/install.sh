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

if [[ "$OS" == "GNU/Linux" ]]; then
    OS="linux"
fi


URL="https://github.com/EzpieCo/GetHooky/releases/download/v1.4.0/hooky-${OS}-${ARCH}"

curl -L "$URL" -o hooky

mv hooky /usr/bin/hooky
chmod +x /usr/bin/hooky

echo "✅ GetHooky installed to /usr/bin"
echo "🚀 Get started with 'hooky --help'"
