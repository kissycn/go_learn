#!/bin/bash

ages=(10 20 30 40 50 60)
ages[6]=70
echo ${ages[1]}
echo ${ages[@]}
echo ${ages[*]}
echo ${#ages[*]}