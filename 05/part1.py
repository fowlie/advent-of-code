from sys import maxsize


debug = False
map = {}
with open("input") as input:
    current_map = ""
    for line in input:
        if line[0].isalpha():
            current_map = line.split(" ")[0].strip()
        elif line[0].isdigit():
            arr = line.split(" ")
            arr[-1] = arr[-1].rstrip("\n")
            if map.get(current_map) is None:
                map[current_map] = []
            map[current_map].append(arr)
        else:
            current_map = ""


def translate(n, key):
    for translation in map[key]:
        dst = int(translation[0])
        src = int(translation[1])
        len = int(translation[2])

        if debug:
            print(f"n: {n}, dst: {dst}, src: {src}, len: {len}")

        if n >= src and n < src + len:
            return n + dst - src

    return n


def get_location(n):
    soil = translate(n, "seed-to-soil")
    fertilizer = translate(soil, "soil-to-fertilizer")
    water = translate(fertilizer, "fertilizer-to-water")
    light = translate(water, "water-to-light")
    temperature = translate(light, "light-to-temperature")
    humidity = translate(temperature, "temperature-to-humidity")
    location = translate(humidity, "humidity-to-location")
    return location


lowest_location = maxsize
for seed in map["seeds:"][0]:
    location = get_location(int(seed))
    if location < lowest_location:
        lowest_location = location
    if debug:
        print(f"{seed} corresponds to {location=}")

# correct answer is 535088217
print(f"{lowest_location=}")
