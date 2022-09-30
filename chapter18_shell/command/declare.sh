#!/bin/bash

declare -A map
map["one"]="111"
map["two"]="222"

echo ${map["one"]}
echo ${map["two"]}

declare -i m n ret
m=10
n=20
ret=$m+$n
echo $ret

a=10
b=20
echo $((a+b))