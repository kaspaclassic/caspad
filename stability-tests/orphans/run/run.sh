#!/bin/bash
rm -rf /tmp/caspad-temp

caspad --simnet --appdir=/tmp/caspad-temp --profile=6061 &
caspad_PID=$!

sleep 1

orphans --simnet -alocalhost:16511 -n20 --profile=7000
TEST_EXIT_CODE=$?

kill $caspad_PID

wait $caspad_PID
caspad_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "caspad exit code: $caspad_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $caspad_EXIT_CODE -eq 0 ]; then
  echo "orphans test: PASSED"
  exit 0
fi
echo "orphans test: FAILED"
exit 1
