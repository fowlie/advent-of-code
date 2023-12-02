with open("input") as input:
    answer = 0
    for game, line in enumerate(input):
        game_colors = {}
        for cube in line.split(":")[1].split(";"):
            for cube in cube.split(","):
                amount = int(cube.lstrip().split(" ")[0])
                color = cube.lstrip().split(" ")[1].strip("\n")

                if color in game_colors:
                    game_colors[color] = max(game_colors[color], amount)
                else:
                    game_colors[color] = amount

        answer += game_colors["red"] * game_colors["green"] * game_colors["blue"]

    # Correct answer: 65371
    print(f"{answer=}")
