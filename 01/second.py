#!/usr/bin/env python3

file = open("input", "r")

total = 0

# map to replace words with numbers
number_map = {
    "one": "1",
    "two": "2",
    "three": "3",
    "four": "4",
    "five": "5",
    "six": "6",
    "seven": "7",
    "eight": "8",
    "nine": "9",
}


def getNumbers(line):
    number = ""
    for i in range(len(line)):
        if line[i].isdigit():
            number += line[i]
            continue
        for j in range(3, 6):
            if i + j > len(line):
                continue
            token = line[i : i + j]
            if token in number_map:
                number += number_map[token]
                i += j
                break

    return number


for line in file:
    line = getNumbers(line)
    total += int(f"{line[0]}{line[-1]}")

# Correct answer is 54719
print("The sum of the calibration values are %d" % total)
