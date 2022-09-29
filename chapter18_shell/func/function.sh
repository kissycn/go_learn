#!/bin/bash

function sum() {
    S=0
    S=$(($1+$2))
    echo $S
}

read -p "please input the num1:" n1;
read -p "please input the num2:" n2;

sum $n1 $n2