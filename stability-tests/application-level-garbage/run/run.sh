#!/bin/bash
rm -rf /tmp/Kalibrium-temp

Kalibrium --devnet --appdir=/tmp/Kalibrium-temp --profile=6061 --loglevel=debug &
KALIPAD_PID=$!
KALIPAD_KILLED=0
function killKalibriumdIfNotKilled() {
    if [ $KALIPAD_KILLED -eq 0 ]; then
      kill $KALIPAD_PID
    fi
}
trap "killKalibriumdIfNotKilled" EXIT

sleep 1

application-level-garbage --devnet -alocalhost:16611 -b blocks.dat --profile=7000
TEST_EXIT_CODE=$?

kill $KALIPAD_PID

wait $KALIPAD_PID
KALIPAD_KILLED=1
KALIPAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Kalibrium exit code: $KALIPAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $KALIPAD_EXIT_CODE -eq 0 ]; then
  echo "application-level-garbage test: PASSED"
  exit 0
fi
echo "application-level-garbage test: FAILED"
exit 1
