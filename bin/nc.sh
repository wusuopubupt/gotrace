#!/bin/bash

i=0
while [[ $i -lt 10 ]]; do
    echo "hello "$i | nc localhost 5000
    sleep 1
    ((++i))
done
