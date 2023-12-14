parts = []
lines = []
with open("sample3") as input:
    for line in input:
        lines.append(line.strip())

ratios = []
for i in range(0, len(lines)):  # for each line
    for j in range(0, len(lines[i])):  # for each char on line
        gears = {}

        if lines[i][j] == "*":
            # scan left
            if j > 0 and lines[i][j - 1].isdigit():
                parsed_number = ""
                k = j - 1
                while k >= 0 and lines[i][k].isdigit():
                    parsed_number = lines[i][k] + parsed_number
                    k -= 1
                k = max(0, k)
                if parsed_number != "":
                    gears.setdefault((i,k), [])
                    gears[(i,k)].append(int(parsed_number))

            # scan right
            if j < len(lines[i]) - 1 and lines[i][j + 1].isdigit():
                parsed_number = ""
                k = j + 1
                while k < len(lines[i]) and lines[i][k].isdigit():
                    parsed_number += lines[i][k]
                    k += 1
                k = min(len(lines[i]), k)
                if parsed_number != "":
                    gears.setdefault((i,k), [])
                    gears[(i,k)].append(int(parsed_number))

            # scan up
            if i > 0:
                parsed_number = ""
                index = 0
                for k in range(j - 1, j + 2):
                    if parsed_number != "" and not lines[i - 1][k].isdigit():
                        parsed_number = ""
                        continue
                    if k < len(lines[i - 1]) and lines[i - 1][k].isdigit():
                        # expand to the left
                        while k > 0 and lines[i - 1][k - 1].isdigit():
                            k -= 1
                        index = k

                    # scan to the right
                    while k < len(lines[i]) and lines[i - 1][k].isdigit():
                        parsed_number += lines[i - 1][k]
                        # if i == 10:
                        #     print(f"{parsed_number=}, {k=}")
                        if k < len(lines[i]):
                            k += 1

                    if parsed_number != "":
                        gears.setdefault((i-1,index), [])
                        gears[(i-1,index)].append(int(parsed_number))
                        # if i == 10:
                        #     print(f"adding {parsed_number} to {(i-1,index)}, {i=}, {index=}, {k=}")

            # scan down
            if i < len(lines) - 1:
                parsed_number = ""
                index = 0
                for k in range(j - 1, j + 2):
                    print(f"{parsed_number=}, {k=}")
                    if parsed_number != "" and not lines[i + 1][k].isdigit():
                        parsed_number = ""
                        continue
                    if k < len(lines[i + 1]) and lines[i + 1][k].isdigit():
                        # expand to the left
                        while k > 0 and lines[i + 1][k - 1].isdigit():
                            k -= 1
                        index = k

                    # scan to the right
                    while k < len(lines[i]) and lines[i + 1][k].isdigit():
                        parsed_number += lines[i + 1][k]
                        if k < len(lines[i]):
                            k += 1

                if parsed_number != "":
                    gears.setdefault((i+1,index), [])
                    gears[(i+1,index)].append(int(parsed_number))

        # remove empty arrays
        gears = {k: v for k, v in gears.items() if v is not None}
        if len(gears) == 2:
            for key, value in gears.items():
                parts.append(value[0])
        elif len(gears) > 0:
            print(f"{gears} is not a valid gear...?")

# correct answer is 6756 (for sample2)
for i in range(0, len(parts)):
    if i % 2 > 0:
        print(f"multiplying {parts[i]} and {parts[i - 1]}")
        ratios.append(parts[i] * parts[i - 1])

# sum up all the ratios
sum = 0
for ratio in ratios:
    sum += ratio

# guesses: 
# - 63913958 (too low)
print("the answer is", sum)
