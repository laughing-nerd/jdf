#!/bin/bash
set -e

# Validate dependency
if ! command -v jq &>/dev/null; then
  echo "❌ jq is required but not installed"
  exit 1
fi

# Remove 'refs/tags/' prefix
TAG_NAME="${GITHUB_REF#refs/tags/}"
echo "New tag: $TAG_NAME"

# Enforce semantic versioning
if ! [[ "$TAG_NAME" =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
  echo "❌ Tag $TAG_NAME does not follow semantic versioning (vX.Y.Z)"
  exit 1
fi

# Fetch main branch to ensure it exists locally
git fetch origin main --quiet

# Check if the tag commit is in main branch history
if ! git merge-base --is-ancestor "$GITHUB_SHA" origin/main; then
  echo "❌ Tag $TAG_NAME does not point to a commit on main branch"
  exit 1
fi

echo "✅ Tag $TAG_NAME points to a commit on main branch"

# Get the latest release tag
LATEST_RELEASE=$(curl -s https://api.github.com/repos/${GITHUB_REPOSITORY}/releases/latest | jq -r .tag_name)

# If no prev release exists, then proceed with the release
if [[ "$LATEST_RELEASE" == "null" || -z "$LATEST_RELEASE" ]]; then
  echo "✅ No prior release found. Proceeding with release of $TAG_NAME"
  exit 0
fi

echo "Latest release: $LATEST_RELEASE"

# Compare versions
if [ "$(printf '%s\n' "$LATEST_RELEASE" "$TAG_NAME" | sort -V | tail -n1)" != "$TAG_NAME" ]; then
  echo "❌ Error: New tag $TAG_NAME is not greater than the current release $LATEST_RELEASE"
  exit 1
fi

echo "✅ New tag $TAG_NAME is greater than $LATEST_RELEASE"