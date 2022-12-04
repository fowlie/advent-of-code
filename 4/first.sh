#!/bin/bash

total=0
for line in $(cat input); do
  first=$(echo $line | cut -d, -f1)
  second=$(echo $line | cut -d, -f2)
  first_min=$(echo $first | cut -d- -f1)
  first_max=$(echo $first | cut -d- -f2)
  second_min=$(echo $second | cut -d- -f1)
  second_max=$(echo $second | cut -d- -f2)
  
  if [[ $first_min -le $second_min && $first_max -ge $second_max
     || $first_min -ge $second_min && $first_max -le $second_max ]]
  then
    ((total=total+1))
  fi
done

echo "Number of assignmet pairs which fully contain the other: $total"
