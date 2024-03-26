#!/bin/bash
rm -rf /tmp/Kalibrium-temp

Kalibrium --devnet --appdir=/tmp/Kalibrium-temp --profile=6061 --loglevel=debug &
KALIPAD_PID=$!

sleep 1

rpc-stability --devnet -p commands.json --profile=7000
TEST_EXIT_CODE=$?

kill $KALIPAD_PID

wait $KALIPAD_PID
KALIPAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Kalibrium exit code: $KALIPAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $KALIPAD_EXIT_CODE -eq 0 ]; then
  echo "rpc-stability test: PASSED"
  exit 0
fi
echo "rpc-stability test: FAILED"
exit 1
