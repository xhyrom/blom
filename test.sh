#!/bin/sh

set -eu

TARGETS="./qbe/tests"

for TARGET in $TARGETS; do
    echo "Running tests in $TARGET..."
    go test "$TARGET"
done

echo "All tests completed successfully."
