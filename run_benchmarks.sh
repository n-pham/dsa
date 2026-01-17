#!/bin/bash
set -e

TARGET=${1:-"."}

echo "=== Python Benchmarks ==="
if [ "$TARGET" == "." ]; then
    uv run pytest --benchmark-only || echo "No python benchmarks found."
else
    uv run pytest "$TARGET" --benchmark-only || echo "No python benchmarks found."
fi

echo "=== Go Benchmarks ==="
if [ "$TARGET" == "." ]; then
    go test -bench=. ./...
else
    go test -bench=. "./$TARGET/..."
fi

echo "=== Rust Benchmarks ==="
if [ "$TARGET" == "." ]; then
    cargo bench --workspace
else
    cargo bench -p "$TARGET"
fi