#!/bin/bash -e

log="/var/log/stupid-httpd/events.log"

cd "$(dirname "$0")"
mkdir -p "$(dirname "$log")"
cat /dev/null > "$log"
export LISTEN=":80"
./linux-amd64 > /var/log/stupid-httpd/events.log 2>&1  &
