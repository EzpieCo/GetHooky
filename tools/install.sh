#!bin/bash

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

# Some idiot thought it was a great idea to use windows with linux like feel
if [[ "$OS" == "Cygwin" ]]; then
   OS="windows" 
fi

URL="https://github.com/EzpieCo/GetHooky/releases/latest/download/hooky-${OS}-${ARCH}"

curl -L "$URL" -o hooky
mv hooky /usr/bin/hooky

echo "✅ GetHooky installed successfully! Get started with 'hooky --help'"
