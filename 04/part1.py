winning = []
hand = []
points = []
with open("input") as input:
    for line in input:
        line = line.split(":")[1].strip()
        winning.append(line.split("|")[0].strip().split(" "))
        hand.append(line.split("|")[1].strip().split(" "))
        points.append(0)

for i in range(len(winning)):
    for h in winning[i]:
        if h in hand[i]:
            if points[i] == 0:
                points[i] = 1
            else:
                points[i] = points[i] * 2

# sum all numbers in points list
sum = 0
for i in points:
    sum += i
print(f"the answer is {sum}")
