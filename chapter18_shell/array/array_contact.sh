#!/bin/bash

array1=(10 20 30 40 50)
array2=(60 70 80 90 100)
array_new=(${array1[*]} ${array2[*]})
echo ${array_new[*]}

unset array_new[0]
echo ${array_new[*]}