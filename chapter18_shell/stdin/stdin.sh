#!/bin/bash

for item in "one" "two" "three"; do
    echo $item >>log.txt
done

ls java 2>>log.txt
ls -l >>out.log 2>&1
ls -l >>/dev/null