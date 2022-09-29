#!/bin/bash

for i in "$*" ; do
    echo $i
done

for j in "$@" ; do
    echo $j
done