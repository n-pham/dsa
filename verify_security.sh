#!/bin/bash
set -e

TARGET=${1:-"."}

echo "=== Python Security Check (Bandit) ==="
# Bandit handles directories and files.
uv run bandit -c pyproject.toml -r "$TARGET" || echo "Bandit found issues (or failed to run)"

echo "=== Go Security Check (Gosec) ==="
if command -v gosec &> /dev/null; then
    if [ "$TARGET" == "." ]; then
        gosec ./...
    else
        # Target specific directory recursively
        gosec "./$TARGET/..."
    fi
else
    echo "gosec not found."
fi

echo "=== Rust Security Check (Clippy) ==="
# cargo audit checks dependencies (global/workspace level mainly) so we run it only if target is "." or we skip it for speed on single modules.
# We will just run clippy for the specific target.

if [ "$TARGET" == "." ]; then
    if command -v cargo-audit &> /dev/null; then cargo audit; fi
    echo "Running Clippy on Workspace..."
    cargo clippy --workspace -- -D warnings
else
    echo "Running Clippy on package: $TARGET..."
    # Assumes folder name matches package name
    cargo clippy -p "$TARGET" -- -D warnings
fi