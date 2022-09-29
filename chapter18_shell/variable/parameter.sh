#!/bin/bash

echo "Process ID: $$"
echo "File param: $0"
echo "First param: $1"
echo "Seconds name: $2"
echo "Total: $#"

for i in "$*"; do
    echo $i
done

for j in "$@"; do
    echo $j
done