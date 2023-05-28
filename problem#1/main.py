#!/bin/python3

import sys


t = int(input().strip())
for a0 in range(t):
    n = int(input().strip())
    total_sum = 0
    
    numMultiples3 = (n-1) // 3
    sumMultiples3 = 3 * (numMultiples3 * (numMultiples3 + 1)) // 2
    
    numMultiples5 = (n-1) // 5
    sumMultiples5 = 5 * (numMultiples5 * (numMultiples5 + 1)) // 2

    numMultiples15 = (n-1) // 15
    sumMultiples15 = 15 * (numMultiples15 * (numMultiples15 + 1)) // 2
    
    total_sum = (sumMultiples3 + sumMultiples5 - sumMultiples15)
    print(total_sum)
