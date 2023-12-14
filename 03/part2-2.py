import re

lines = []
with open("sample") as input:
    for line in input:
        lines.append(line.strip())

for i in range(0, len(lines)):  # for each line
    # check if symbol '*' is present
   if '*' in lines[i]:
        print('Line', i, 'has a *')
        r = re.compile(r'\d+')
        if i > 0:
            res = r.findall(lines[i-1])
            print(f"{res=}")
