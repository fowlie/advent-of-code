parts = []
lines = []
with open("input") as input:
    for line in input:
        lines.append(line.strip())


for i in range(0, len(lines)):  # for each line
    parsed_number = ""

    for j in range(0, len(lines[i])):  # for each char on line
        if lines[i][j].isdigit():
            parsed_number += lines[i][j]

            if j < len(lines[i]) - 1:  # prevent index out of bounds
                if lines[i][j + 1].isdigit():  # next also digit?
                    continue

            # define the "width" of the scan field
            # use min() and max() to avoid indexoutofbounds
            start = max(j - len(parsed_number), 0)
            end = min(j + 2, len(lines[i]))

            # scan for symbols in field around number
            for x in range(i - 1, i + 2):
                for y in range(start, end):
                    if x < 0 or x >= len(lines):
                        continue
                    if not lines[x][y].isdigit() and lines[x][y] != ".":
                        parts.append(int(parsed_number))
                        break
            parsed_number = ""

# correct answer is 539590
print("the answer is", sum([int(i) for i in parts]))
