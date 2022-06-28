#!/bin/bash

for k in r_queue render p_queue parser f_queue formatter;
do
    for i in {1..10}; 
    do
	Data='{"action":"get","role":"'$k'"}'
	echo $Data
	echo $Data | nc 127.0.0.1 10000
	echo
    done
done