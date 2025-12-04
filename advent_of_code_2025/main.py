def secret_entrance(lines: list[str]) -> int:
    num = 50
    password = 0
    for line in lines:
        distance = int(line[1:])
        if line[0] == "R":
            num = (num + distance) % 100
        else:
            num = (num - distance) % 100
        if num == 0:
            password += 1
    return password


def test_secret_entrance():
    assert (
        secret_entrance(
            ["L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"]
        )
        == 3
    )


def secret_entrance_2(lines: list[str]) -> int:
    unconstrained_pos = 50
    password = 0
    for line in lines:
        distance = int(line[1:])
        prev_pos = unconstrained_pos
        if line[0] == "R":
            unconstrained_pos += distance
        else:
            unconstrained_pos -= distance

        if unconstrained_pos > prev_pos:
            # right move, cross if end/100 > start/100, but not if start is multiple of 100
            password += (unconstrained_pos - 1) // 100 - (prev_pos - 1) // 100
        elif unconstrained_pos < prev_pos:
            # left move, cross if start/100 > end/100, but not if start is multiple of 100
            password += prev_pos // 100 - unconstrained_pos // 100
    return password


def test_secret_entrance_2():
    assert (
        secret_entrance_2(
            ["L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"]
        )
        == 6
    )
    assert secret_entrance_2(["R1000"]) == 10


if __name__ == "__main__":
    with open("./day_1_input.txt", "r") as file:
        lines = file.readlines()
        print(secret_entrance_2(lines))
