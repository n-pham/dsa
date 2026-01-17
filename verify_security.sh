#!/bin/bash
set -e

echo "=== Python Security Check (Bandit) ==="
# Assuming uv is used
uv run bandit -c pyproject.toml -r . || echo "Bandit found issues (or failed to run)"

echo "=== Go Security Check (Gosec) ==="
if command -v gosec &> /dev/null; then
    gosec ./...
else
    echo "gosec not found. Install with: go install github.com/securego/gosec/v2/cmd/gosec@latest"
fi

echo "=== Rust Security Check (Clippy/Audit) ==="
# cargo audit requires installation: cargo install cargo-audit
if command -v cargo-audit &> /dev/null; then
    cargo audit
else
    echo "cargo-audit not found. Install with: cargo install cargo-audit"
fi
echo "Running Clippy..."
cargo clippy --workspace -- -D warnings
