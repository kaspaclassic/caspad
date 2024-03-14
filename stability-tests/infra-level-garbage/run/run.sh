#!/bin/bash
rm -rf /tmp/caspad-temp

caspad --devnet --appdir=/tmp/caspad-temp --profile=6061 &
caspad_PID=$!

sleep 1

infra-level-garbage --devnet -alocalhost:16611 -m messages.dat --profile=7000
TEST_EXIT_CODE=$?

kill $caspad_PID

wait $caspad_PID
caspad_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "caspad exit code: $caspad_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $caspad_EXIT_CODE -eq 0 ]; then
  echo "infra-level-garbage test: PASSED"
  exit 0
fi
echo "infra-level-garbage test: FAILED"
exit 1
