#!/bin/bash

echo "Start"

if [ -f /vault/secrets/secrets ]
then
  echo "Loading values from secrets file"
  source /vault/secrets/secrets

  export DATA

  sudo rm /vault/secrets/secrets
fi

sleep 5

exec ./app