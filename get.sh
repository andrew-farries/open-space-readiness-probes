#! /bin/sh

while [ true ]; do
  hostname=$(curl http://34.77.140.164:8090 2> /dev/null)
  time=$(date "+%H:%M:%S")

  echo "[$time] $hostname"
  sleep 0.5
done
