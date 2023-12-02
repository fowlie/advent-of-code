with open("input") as input:
    answer = 0
    for game, line in enumerate(input):
        possible = True
        for cube in line.split(":")[1].split(";"):
            for cube in cube.split(","):
                amount = int(cube.lstrip().split(" ")[0])
                color = cube.lstrip().split(" ")[1].strip("\n")

                match color:
                    case "red":
                        if amount > 12:
                            possible = False
                    case "green":
                        if amount > 13:
                            possible = False
                    case "blue":
                        if amount > 14:
                            possible = False
                    case _:
                        print(f"Unknown color: {color}")

        if possible:
            answer += game + 1

    # Correct answer: 3059
    print(f"{answer=}")
