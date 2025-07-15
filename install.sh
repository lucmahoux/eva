#!/bin/sh

set -e

REPO="lucmahoux/eva"
BINARY_NAME="eva"
INSTALL_DIR="/usr/local/bin"
VERSION=${VERSION:-latest}

echo "📦 Installing $BINARY_NAME..."

# Detect OS and Arch
OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

# Normalize architecture
case "$ARCH" in
  x86_64) ARCH="amd64" ;;
  arm64|aarch64) ARCH="arm64" ;;
  *) echo "❌ Unsupported architecture: $ARCH"; exit 1 ;;
esac

# Resolve latest version tag
if [ "$VERSION" = "latest" ]; then
  VERSION=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | grep tag_name | cut -d '"' -f4)
fi

VERSION_NO_V=$(echo "$VERSION" | sed 's/^v//')
FILENAME="${BINARY_NAME}_${VERSION_NO_V}_${OS}_${ARCH}.tar.gz"
URL="https://github.com/${REPO}/releases/download/${VERSION}/${FILENAME}"

echo "⬇️ Downloading $FILENAME..."

# Check if asset exists
STATUS=$(curl -s -L -o /dev/null -w "%{http_code}" "$URL")
if [ "$STATUS" != "200" ]; then
  echo "❌ Failed to download binary (HTTP $STATUS)"
  echo "👉 Check that version '$VERSION' exists and asset '$FILENAME' is available."
  exit 1
fi

# Download and extract
curl -sSL "$URL" | tar -xz -C /tmp

echo "🚚 Moving binary to $INSTALL_DIR..."
sudo mv "/tmp/$BINARY_NAME" "$INSTALL_DIR/$BINARY_NAME"
chmod +x "$INSTALL_DIR/$BINARY_NAME"

echo "✅ Installed $BINARY_NAME $VERSION to $INSTALL_DIR"
