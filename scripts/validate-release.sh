#!/bin/bash

# Validate release tag before building
# Usage: ./scripts/validate-release.sh <new-tag>
#
# Checks:
# 1. Tag format is valid (vX.Y.Z)
# 2. Tag doesn't already have a release
# 3. New version is greater than the latest release

set -e

NEW_TAG="${1:-}"
REPO="${GITHUB_REPOSITORY:-laughing-nerd/jdf}"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

log_error() { echo -e "${RED}❌ ERROR: $1${NC}" >&2; }
log_success() { echo -e "${GREEN}✅ $1${NC}"; }
log_info() { echo -e "${YELLOW}ℹ️  $1${NC}"; }

# Check if tag is provided
if [[ -z "$NEW_TAG" ]]; then
    log_error "No tag provided. Usage: $0 <tag>"
    exit 1
fi

log_info "Validating release tag: $NEW_TAG"

# 1. Validate tag format (vX.Y.Z or vX.Y.Z-suffix)
if [[ ! "$NEW_TAG" =~ ^v[0-9]+\.[0-9]+\.[0-9]+(-[a-zA-Z0-9.]+)?$ ]]; then
    log_error "Invalid tag format: $NEW_TAG"
    echo "Expected format: vX.Y.Z (e.g., v1.0.0, v2.1.3-beta)"
    exit 1
fi
log_success "Tag format is valid"

# 2. Check if release already exists (using GitHub API)
if [[ -n "${GITHUB_TOKEN:-}" ]]; then
    RELEASE_CHECK=$(curl -s -o /dev/null -w "%{http_code}" \
        -H "Authorization: token $GITHUB_TOKEN" \
        "https://api.github.com/repos/$REPO/releases/tags/$NEW_TAG")
    
    if [[ "$RELEASE_CHECK" == "200" ]]; then
        log_error "Release $NEW_TAG already exists!"
        echo "Please use a new version tag."
        exit 1
    fi
    log_success "No existing release for $NEW_TAG"
else
    log_info "GITHUB_TOKEN not set, skipping release existence check"
fi

# 3. Compare with latest release version
if [[ -n "${GITHUB_TOKEN:-}" ]]; then
    LATEST_TAG=$(curl -s \
        -H "Authorization: token $GITHUB_TOKEN" \
        "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name"' | head -1 | cut -d'"' -f4)
    
    if [[ -n "$LATEST_TAG" && "$LATEST_TAG" != "null" ]]; then
        log_info "Latest release: $LATEST_TAG"
        
        # Extract version numbers (remove 'v' prefix and any suffix)
        NEW_VER=$(echo "$NEW_TAG" | sed 's/^v//' | cut -d'-' -f1)
        OLD_VER=$(echo "$LATEST_TAG" | sed 's/^v//' | cut -d'-' -f1)
        
        # Compare versions using sort -V
        HIGHER=$(echo -e "$OLD_VER\n$NEW_VER" | sort -V | tail -1)
        
        if [[ "$NEW_VER" == "$OLD_VER" ]]; then
            log_error "Version $NEW_TAG is same as latest release $LATEST_TAG"
            exit 1
        elif [[ "$HIGHER" == "$OLD_VER" ]]; then
            log_error "Version $NEW_TAG is older than latest release $LATEST_TAG"
            echo "New version must be greater than the current release."
            exit 1
        fi
        log_success "Version $NEW_TAG is newer than $LATEST_TAG"
    else
        log_info "No previous releases found, this will be the first release"
    fi
else
    log_info "GITHUB_TOKEN not set, skipping version comparison"
fi

echo ""
log_success "All validation checks passed! Ready to build $NEW_TAG"
exit 0

