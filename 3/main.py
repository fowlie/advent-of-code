total = 0
total2 = 0
match = []
group_offset = 0
item_types = []
# file = open("input", "r")
with open("input") as f:
  file = f.read().splitlines()
for line in file:
  first, second = "", ""
  half = int(len(line)/2)

  # split string in the middle
  for i in range(0,half):
    first += line[i]
    second += line[i+half]

  # find matching chars in both halves, no duplicates
  for i in range(0, len(first)):
    if first[i] in second:
      if not first[i] in match:
        match += first[i]
        code = ord(first[i])
        if code > 90:
          code = code - 96
        else:
          code = code - 38
        total += code

  item_types = []
  # check groups of 3 lines
  if group_offset % 3 == 0:
    for i in range(0, len(file[group_offset])):
      items = set(file[group_offset]).intersection(file[group_offset + 1]).intersection(file[group_offset + 2])
      if len(items) == 1:
        item_types += items.pop()
        break

    for item in item_types:
      code = ord(item)
      if code > 90:
        code = code - 96
      else:
        code = code - 38
      total2 = total2 + code

  match = []
  group_offset = group_offset + 1

print(total)
if total != 7785:
  print("program corrupt")

print(total2)
if total2 != 2633:
  print("program corrupt")
