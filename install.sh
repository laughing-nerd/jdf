#!/usr/bin/env bash
set -euo pipefail

REPO="laughing-nerd/sensible"
VERSION="${VERSION:-latest}"
OS="$(uname | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m)"

log() {
  echo "➤ $1"
}

die() {
  echo "❌ $1" >&2
  exit 1
}

# Normalize architecture
case "$ARCH" in
  x86_64) ARCH="amd64" ;;
  aarch64 | arm64) ARCH="arm64" ;;
  *) die "Unsupported architecture: $ARCH" ;;
esac

# Resolve version
if [ "$VERSION" = "latest" ]; then
  log "Fetching latest release version from GitHub…"
  response="$(curl -fsSL "https://api.github.com/repos/$REPO/releases/latest" || true)"

  if echo "$response" | grep -q '"message": "Not Found"'; then
    die "No releases found for $REPO. Please create one on GitHub."
  fi

  VERSION="$(echo "$response" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')"
fi

log "Installing sensible $VERSION for $OS/$ARCH"

BINARY="sensible-$OS-$ARCH"
URL="https://github.com/$REPO/releases/download/$VERSION/$BINARY"
log "Downloading $BINARY from $URL"

if ! curl -fSL "$URL" -o sensible; then
  die "Download failed: $URL"
fi

chmod +x sensible

# Determine install location
INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"

if [ ! -w "$INSTALL_DIR" ]; then
  FALLBACK="$HOME/.local/bin"
  log "$INSTALL_DIR is not writable. Falling back to $FALLBACK"

  mkdir -p "$FALLBACK"
  INSTALL_DIR="$FALLBACK"
fi

DEST="$INSTALL_DIR/sensible"
mv sensible "$DEST"

log "✅ Sensible installed to $DEST"

# Suggest PATH update if fallback was used
if echo "$DEST" | grep -q "$HOME/.local/bin" && ! echo "$PATH" | grep -q "$HOME/.local/bin"; then
  echo "⚠️  Add this to your shell profile to use 'sensible' everywhere:"
  echo "    export PATH=\"\$HOME/.local/bin:\$PATH\""
fi