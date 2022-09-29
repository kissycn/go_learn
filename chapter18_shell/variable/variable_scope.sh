#!/bin/bash

function foo() {
    a=1
}

foo
echo $a

function foo1() {
    local b=2
}

foo1
echo $b
