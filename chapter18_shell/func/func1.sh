#!/bin/bash

function getSum() {
  local sum=0
  for num in $@ ; do
      ((sum+=num))
  done

  return $sum
}

getSum 1 2 3 4 5