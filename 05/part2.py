import functools

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
        if n >= src and n < src + len:
            return n + dst - src
    return n


@functools.lru_cache(maxsize=1000000)
def get_location(n):
    soil = translate(n, "seed-to-soil")
    fertilizer = translate(soil, "soil-to-fertilizer")
    water = translate(fertilizer, "fertilizer-to-water")
    light = translate(water, "water-to-light")
    temperature = translate(light, "light-to-temperature")
    humidity = translate(temperature, "temperature-to-humidity")
    location = translate(humidity, "humidity-to-location")
    return location


lowest_location = 9999999999
len = len(map["seeds:"][0])
for i in range(0, len - 1):
    if i % 2 == 1:
        continue
    start = int(map["seeds:"][0][i])
    end = start + int(map["seeds:"][0][i + 1])

    print(f"scanning seed range {i} of {len}")

    for seed in range(start, end):
        location = get_location(int(seed))
        if location < lowest_location:
            lowest_location = location

print(f"{lowest_location=}")
