import re
import collections

stacks = [collections.deque([])]

with open("input") as f:
  file = f.read().splitlines()
for line in file:
  if re.search("\\[[A-Z]\\]+", line):
    stack_size = int((len(line) + 1) / 4)
    for i in range(0,stack_size):
      if not len(stacks) >= stack_size: stacks.append(collections.deque())
      crate = line[4 * i + 1]
      if not crate == " ": stacks[i].append(crate)
  else:
      m = re.search(r"move (\b\d+) from (\b\d+) to (\b\d+)", line)
      if not m == None:
        for i in range(int(m.group(1))):
          item = stacks[int(m.group(2))-1].popleft()
          stacks[int(m.group(3))-1].appendleft(item)

print("The answer to part 1 is:")
for stack in stacks:
  print(stack.popleft(), end = "")
print("\n")



stacks = [collections.deque([])]

with open("input") as f:
  file = f.read().splitlines()
for line in file:
  if re.search("\\[[A-Z]\\]+", line):
    stack_size = int((len(line) + 1) / 4)
    for i in range(0,stack_size):
      if not len(stacks) >= stack_size: stacks.append(collections.deque())
      crate = line[4 * i + 1]
      if not crate == " ": stacks[i].append(crate)
  else:
      m = re.search(r"move (\b\d+) from (\b\d+) to (\b\d+)", line)
      if not m == None:
        tmp = collections.deque([])
        for i in range(int(m.group(1))):
          item = stacks[int(m.group(2))-1].popleft()
          tmp.appendleft(item)
        for i in range(len(tmp)):
          item = tmp.popleft()
          stacks[int(m.group(3))-1].appendleft(item)


print("The answer to part 2 is:")
for stack in stacks:
  print(stack.popleft(), end = "")
print("")
