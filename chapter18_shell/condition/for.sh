#!/bin/bash

sum=0
sum1=0
for (( i = 0; i <= 100; i++ )); do
    sum=$(($sum + i))
    sum1=$[$sum1+i]
done
echo $sum
echo $sum1