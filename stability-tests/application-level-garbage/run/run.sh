#!/bin/bash
rm -rf /tmp/caspad-temp

caspad --devnet --appdir=/tmp/caspad-temp --profile=6061 --loglevel=debug &
caspad_PID=$!
caspad_KILLED=0
function killcaspadIfNotKilled() {
    if [ $caspad_KILLED -eq 0 ]; then
      kill $caspad_PID
    fi
}
trap "killcaspadIfNotKilled" EXIT

sleep 1

application-level-garbage --devnet -alocalhost:16611 -b blocks.dat --profile=7000
TEST_EXIT_CODE=$?

kill $caspad_PID

wait $caspad_PID
caspad_KILLED=1
caspad_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "caspad exit code: $caspad_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $caspad_EXIT_CODE -eq 0 ]; then
  echo "application-level-garbage test: PASSED"
  exit 0
fi
echo "application-level-garbage test: FAILED"
exit 1
