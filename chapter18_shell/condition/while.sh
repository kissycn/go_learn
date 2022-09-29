#!/bin/bash

SUM=0
i=1

while [ $i -le 100 ]; do
  SUM=$((SUM+$i))
  i=$(($i+1))
done
echo $SUM