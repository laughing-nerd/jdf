#!/bin/bash
#
# JDF Installer Script
# Usage: curl -fsSL https://raw.githubusercontent.com/laughing-nerd/jdf/master/install.sh | bash
#
# Or with a specific version:
# curl -fsSL https://raw.githubusercontent.com/laughing-nerd/jdf/master/install.sh | bash -s -- v1.0.0
#

set -e

# Configuration
REPO="laughing-nerd/jdf"
BINARY_NAME="jdf"
GITHUB_API="https://api.github.com/repos/$REPO/releases"
GITHUB_DOWNLOAD="https://github.com/$REPO/releases/download"

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color
BOLD='\033[1m'

# Functions
log_info() { echo -e "${BLUE}[INFO]${NC} $1"; }
log_success() { echo -e "${GREEN}[SUCCESS]${NC} $1"; }
log_warn() { echo -e "${YELLOW}[WARN]${NC} $1"; }
log_error() { echo -e "${RED}[ERROR]${NC} $1" >&2; }

print_banner() {
    echo -e "${CYAN}"
    echo '     _     _  __ '
    echo '    (_) __| |/ _|'
    echo '    | |/ _` | |_ '
    echo '    | | (_| |  _|'
    echo '   _/ |\__,_|_|  '
    echo '  |__/           '
    echo -e "${NC}"
    echo -e "${BOLD}JSON Detect and Format${NC}"
    echo ""
}

detect_os() {
    local os
    os="$(uname -s)"
    case "$os" in
        Linux*)  echo "linux" ;;
        Darwin*) echo "darwin" ;;
        MINGW*|MSYS*|CYGWIN*) echo "windows" ;;
        *)
            log_error "Unsupported operating system: $os"
            exit 1
            ;;
    esac
}

detect_arch() {
    local arch
    arch="$(uname -m)"
    case "$arch" in
        x86_64|amd64) echo "amd64" ;;
        arm64|aarch64) echo "arm64" ;;
        *)
            log_error "Unsupported architecture: $arch"
            exit 1
            ;;
    esac
}

get_latest_version() {
    local version
    version=$(curl -fsSL "$GITHUB_API/latest" | grep '"tag_name"' | head -1 | cut -d'"' -f4)
    if [[ -z "$version" ]]; then
        log_error "Failed to fetch latest version"
        exit 1
    fi
    echo "$version"
}

get_install_dir() {
    local os="$1"
    
    # Check if /usr/local/bin is writable
    if [[ -w "/usr/local/bin" ]]; then
        echo "/usr/local/bin"
        return
    fi
    
    # Check if we can use sudo
    if command -v sudo &> /dev/null && sudo -n true 2>/dev/null; then
        echo "/usr/local/bin"
        return
    fi
    
    # Fallback to user's local bin
    local user_bin="$HOME/.local/bin"
    mkdir -p "$user_bin"
    echo "$user_bin"
}

check_path() {
    local install_dir="$1"
    if [[ ":$PATH:" != *":$install_dir:"* ]]; then
        log_warn "$install_dir is not in your PATH"
        echo ""
        echo "Add it to your shell profile:"
        echo ""
        echo -e "  ${CYAN}# For bash (~/.bashrc or ~/.bash_profile)${NC}"
        echo -e "  export PATH=\"\$PATH:$install_dir\""
        echo ""
        echo -e "  ${CYAN}# For zsh (~/.zshrc)${NC}"
        echo -e "  export PATH=\"\$PATH:$install_dir\""
        echo ""
        echo -e "  ${CYAN}# For fish (~/.config/fish/config.fish)${NC}"
        echo -e "  set -gx PATH \$PATH $install_dir"
        echo ""
    fi
}

download_and_install() {
    local version="$1"
    local os="$2"
    local arch="$3"
    local install_dir="$4"
    
    local binary_name="jdf-${os}-${arch}"
    if [[ "$os" == "windows" ]]; then
        binary_name="${binary_name}.exe"
    fi
    
    local download_url="$GITHUB_DOWNLOAD/$version/$binary_name"
    local tmp_dir
    tmp_dir=$(mktemp -d)
    local tmp_file="$tmp_dir/$BINARY_NAME"
    
    log_info "Downloading $binary_name..."
    if ! curl -fsSL "$download_url" -o "$tmp_file"; then
        log_error "Failed to download $download_url"
        rm -rf "$tmp_dir"
        exit 1
    fi
    
    chmod +x "$tmp_file"
    
    # Install
    local dest="$install_dir/$BINARY_NAME"
    if [[ -w "$install_dir" ]]; then
        mv "$tmp_file" "$dest"
    else
        log_info "Requesting sudo permission to install to $install_dir..."
        sudo mv "$tmp_file" "$dest"
    fi
    
    rm -rf "$tmp_dir"
    
    log_success "Installed $BINARY_NAME to $dest"
}

verify_installation() {
    if command -v jdf &> /dev/null; then
        local installed_version
        installed_version=$(jdf 2>&1 | grep -oP 'Version: \K[^\s]+' || echo "unknown")
        log_success "jdf is installed and ready!"
        echo ""
        echo -e "  ${CYAN}Version:${NC} $installed_version"
        echo -e "  ${CYAN}Location:${NC} $(which jdf)"
        echo ""
        echo -e "  ${BOLD}Quick start:${NC}"
        echo -e "    echo '{\"hello\": \"world\"}' | jdf"
        echo ""
    else
        log_warn "jdf installed but not found in PATH"
    fi
}

main() {
    print_banner
    
    local version="${1:-}"
    local os
    local arch
    local install_dir
    
    # Detect system
    os=$(detect_os)
    arch=$(detect_arch)
    
    log_info "Detected OS: $os"
    log_info "Detected Architecture: $arch"
    
    # Get version
    if [[ -z "$version" ]]; then
        log_info "Fetching latest version..."
        version=$(get_latest_version)
    fi
    log_info "Version to install: $version"
    
    # Determine install location
    install_dir=$(get_install_dir "$os")
    log_info "Install location: $install_dir"
    
    echo ""
    
    # Download and install
    download_and_install "$version" "$os" "$arch" "$install_dir"
    
    echo ""
    
    # Check PATH
    check_path "$install_dir"
    
    # Verify
    verify_installation
}

# Run
main "$@"
