#!/bin/bash
rm -rf /tmp/Kalibrium-temp

NUM_CLIENTS=128
Kalibrium --devnet --appdir=/tmp/Kalibrium-temp --profile=6061 --rpcmaxwebsockets=$NUM_CLIENTS &
KALIPAD_PID=$!
KALIPAD_KILLED=0
function killKalibriumdIfNotKilled() {
  if [ $KALIPAD_KILLED -eq 0 ]; then
    kill $KALIPAD_PID
  fi
}
trap "killKalibriumdIfNotKilled" EXIT

sleep 1

rpc-idle-clients --devnet --profile=7000 -n=$NUM_CLIENTS
TEST_EXIT_CODE=$?

kill $KALIPAD_PID

wait $KALIPAD_PID
KALIPAD_EXIT_CODE=$?
KALIPAD_KILLED=1

echo "Exit code: $TEST_EXIT_CODE"
echo "Kalibrium exit code: $KALIPAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $KALIPAD_EXIT_CODE -eq 0 ]; then
  echo "rpc-idle-clients test: PASSED"
  exit 0
fi
echo "rpc-idle-clients test: FAILED"
exit 1
