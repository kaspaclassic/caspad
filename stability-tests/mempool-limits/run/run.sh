#!/bin/bash

APPDIR=/tmp/caspad-temp
caspad_RPC_PORT=29587

rm -rf "${APPDIR}"

caspad --simnet --appdir="${APPDIR}" --rpclisten=0.0.0.0:"${caspad_RPC_PORT}" --profile=6061 &
caspad_PID=$!

sleep 1

RUN_STABILITY_TESTS=true go test ../ -v -timeout 86400s -- --rpc-address=127.0.0.1:"${caspad_RPC_PORT}" --profile=7000
TEST_EXIT_CODE=$?

kill $caspad_PID

wait $caspad_PID
caspad_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "caspad exit code: $caspad_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $caspad_EXIT_CODE -eq 0 ]; then
  echo "mempool-limits test: PASSED"
  exit 0
fi
echo "mempool-limits test: FAILED"
exit 1
