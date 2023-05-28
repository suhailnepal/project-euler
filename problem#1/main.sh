#!/bin/bash

read t

for ((a0=0; a0<t; a0++)); do
    read n
    declare -i total_sum=0
    
    numMultiples3=$(expr \( $n - 1 \) / 3 )
    sumMultiples3=$(expr 3 \* $numMultiples3 \* \( $numMultiples3 + 1 \) / 2 )
    
    numMultiples5=$(expr \( $n - 1 \) / 5 )
    sumMultiples5=$(expr 5 \* $numMultiples5 \* \( $numMultiples5 + 1 \) / 2 )

    numMultiples15=$(expr \( $n - 1 \) / 15 )
    sumMultiples15=$(expr 15 \* $numMultiples15 \* \( $numMultiples15 + 1 \) / 2 )
    
    total_sum=$((sumMultiples3 + sumMultiples5 - sumMultiples15))
    
    echo "$total_sum"

done
