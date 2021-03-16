#!/bin/bash

VERSION="1.1"

main() {
  /usr/sbin/stupid-http
}

err() { echo "$@" 1>&2; }

main $@
