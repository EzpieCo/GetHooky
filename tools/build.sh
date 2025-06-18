#!/bin/bash
set -e

echo "🧹 Cleaning old builds..."

rm -rf ../builds/
mkdir -p ../builds/

APPNAME="hooky"

PLATFORMS=(
  "darwin amd64"
  "darwin arm64"
  "linux amd64"
  "windows amd64"
)

echo "🚧 Building binaries..."

for platform in "${PLATFORMS[@]}"; do
  read -r GOOS GOARCH <<< "$platform"

  output_name="$APPNAME-$GOOS-$GOARCH"
  [ "$GOOS" == "windows" ] && output_name="hooky.exe"  

  echo "🔧 Building for $GOOS/$GOARCH as $output_name"

  GOOS=$GOOS GOARCH=$GOARCH go build -ldflags="-s -w" -o "../builds/$output_name" ".."

  echo "✅ $output_name built!"
done

echo ""
echo "🎉 all builds completed!"
