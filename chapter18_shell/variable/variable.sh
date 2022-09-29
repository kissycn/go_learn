#!/bin/bash

A="a"
B="B"

export HELLO="WORLD"
echo $A $B
echo $HELLO

NAME="WORLD"
echo 'HELLO $NAME'
echo "HELLO $NAME"
ECHO "HELLO${NAME}"