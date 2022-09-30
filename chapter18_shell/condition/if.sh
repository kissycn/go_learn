#!/bin/bash

read -p "enter num:" num

if  ((num==1)); then
    echo "Monday"
elif (( num==2 )); then
    echo "Tuesday"
elif (( num==3 )); then
    echo "Wednesday"
elif (( num ==4 )); then
  echo "Thursday"
else
  echo "Unknown"
fi