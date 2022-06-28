#!/bin/bash


declare -i PREC
PREC=0

for k in r_queue render p_queue parser f_queue formatter;
do
let PREC=$PREC+1
    for i in {1..5}; 
    do
	Data='{"action":"set","role":"'$k'","ip":"127.0.0.1","port":"10'$PREC'0'$i'"}'
	echo $Data
	echo $Data | nc 127.0.0.1 10000
	echo
    done
done