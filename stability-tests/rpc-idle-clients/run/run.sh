#!/bin/bash
rm -rf /tmp/caspad-temp

NUM_CLIENTS=128
caspad --devnet --appdir=/tmp/caspad-temp --profile=6061 --rpcmaxwebsockets=$NUM_CLIENTS &
caspad_PID=$!
caspad_KILLED=0
function killcaspadIfNotKilled() {
  if [ $caspad_KILLED -eq 0 ]; then
    kill $caspad_PID
  fi
}
trap "killcaspadIfNotKilled" EXIT

sleep 1

rpc-idle-clients --devnet --profile=7000 -n=$NUM_CLIENTS
TEST_EXIT_CODE=$?

kill $caspad_PID

wait $caspad_PID
caspad_EXIT_CODE=$?
caspad_KILLED=1

echo "Exit code: $TEST_EXIT_CODE"
echo "caspad exit code: $caspad_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $caspad_EXIT_CODE -eq 0 ]; then
  echo "rpc-idle-clients test: PASSED"
  exit 0
fi
echo "rpc-idle-clients test: FAILED"
exit 1
