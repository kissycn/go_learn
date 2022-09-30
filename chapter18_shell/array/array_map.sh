#!/bin/bash

declare -A color
color["red"]="#ff0000"
color["green"]="#00ff00"
color["blue"]="#0000ff"

color["white"]="#ffffff"
color["black"]="#000000"

#获取所有元素值
for val in ${color[*]} ; do
    echo $val
done

for key in ${!color[*]} ; do
    echo $key
done

for k in ${color[*]} ; do
    echo "${k} -> ${color[$k]}"
done