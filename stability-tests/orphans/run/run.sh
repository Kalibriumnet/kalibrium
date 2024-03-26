#!/bin/bash
rm -rf /tmp/Kalibrium-temp

Kalibrium --simnet --appdir=/tmp/Kalibrium-temp --profile=6061 &
KALIPAD_PID=$!

sleep 1

orphans --simnet -alocalhost:16511 -n20 --profile=7000
TEST_EXIT_CODE=$?

kill $KALIPAD_PID

wait $KALIPAD_PID
KALIPAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Kalibrium exit code: $KALIPAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $KALIPAD_EXIT_CODE -eq 0 ]; then
  echo "orphans test: PASSED"
  exit 0
fi
echo "orphans test: FAILED"
exit 1
