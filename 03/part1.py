import re  # Regular Expressions

parts = []
lines = []
with open("input") as input:
    for line in input:
        lines.append(line.strip())


def is_symbol(c):
    return not c.isdigit() and c != "."


def symbol_in_range(i, start, end):
    if i < 0 or i >= len(lines):
        return False

    # print(f"checking line {i+1} for symbols in range {start}-{end}")
    line = lines[i]
    for i in range(start, end):
        if is_symbol(line[i]):
            return True
    return False


for i, line in enumerate(lines):
    numbers = re.findall(r"\d+", line)
    for number in numbers:
        pos = line.find(number)

        start = max(pos - 1, 0)
        end = min(pos + len(number) + 1, len(lines[i]))
        if (
            symbol_in_range(i, start, end)
            or symbol_in_range(i - 1, start, end)
            or symbol_in_range(i + 1, start, end)
        ):
            parts.append(number)
            # print(f"{number} on line {i+1} is a part number")
        # else:
        # print(f"{number} on line {i+1} is NOT a part number")


# sum all numbers in list
print("the answer is", sum([int(i) for i in parts]))
