#!/usr/bin/env python3

file = open("input", "r")

answer = 0

for line in file:
    line = "".join([s for s in line if s.isdigit()])
    answer += int(f"{line[0]}{line[-1]}")

# Correct answer is 55971
print("The sum of the calibration values are ", answer)
