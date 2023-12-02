#!/bin/bash
s=0
while IFS= read -r l; do
	a=$(echo $l | sed 's/[^0-9]*//g')
	s=$((s + $(echo "${a:0:1}${a: -1}")))
done <input
echo $s
