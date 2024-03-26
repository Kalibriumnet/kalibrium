#!/bin/bash
rm -rf /tmp/Kalibrium-temp

Kalibrium --devnet --appdir=/tmp/Kalibrium-temp --profile=6061 &
KALIPAD_PID=$!

sleep 1

infra-level-garbage --devnet -alocalhost:16611 -m messages.dat --profile=7000
TEST_EXIT_CODE=$?

kill $KALIPAD_PID

wait $KALIPAD_PID
KALIPAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Kalibrium exit code: $KALIPAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $KALIPAD_EXIT_CODE -eq 0 ]; then
  echo "infra-level-garbage test: PASSED"
  exit 0
fi
echo "infra-level-garbage test: FAILED"
exit 1
