#!/bin/bash
a=$1
if  [ ! -n "$a" ] ;then
	echo -n "input: \"sh setARGS.sh instance,connNum\""
	exit 0
fi
echo $a > ARGS

