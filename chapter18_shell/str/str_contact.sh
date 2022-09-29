#!/bin/bash

S1=HELLO
S2=WORLD

echo $S1 $S2
echo "$S1 $S2"

S3=www.baidu.com
# 截取4到9位
echo ${S3: 4 : 5}
# 截取4到末尾
echo ${S3: 4}
# 按照字符串进行截取
echo ${S3#*.}
echo ${S3%.}