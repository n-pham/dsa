#!/bin/bash
set -e

echo "=== Python Benchmarks ==="
# Runs pytest-benchmark. Ensure tests marked with appropriate decorators if any, 
# or it will just run tests and report.
uv run pytest --benchmark-only || echo "No python benchmarks found or failed."

echo "=== Go Benchmarks ==="
go test -bench=. ./...

echo "=== Rust Benchmarks ==="
# Runs benchmarks for all members defined in the [workspace] of Cargo.toml
cargo bench --workspace
