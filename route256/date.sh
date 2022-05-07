#!/bin/bash
read -r p
a=`echo $p | awk -F ' ' '{print $1}'`
b=`echo $p | awk -F ' ' '{print $2}'`
t1=`date -d $a +%s`
t2=`date -d $b +%s`
d=$((($t2-$t1) / 60 / 60 / 24))
echo $d
