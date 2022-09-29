#!/bin/bash

function add() {
    return `expr $1+$2`
}

add 1 2
echo $?