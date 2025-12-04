def _is_repeated_num(n: int) -> bool:
    # Count digits
    digits = 0
    tmp = n
    while tmp > 0:
        tmp //= 10
        digits += 1

    # Try all possible block lengths (must divide total digits)
    for block_len in range(1, digits // 2 + 1):
        if digits % block_len != 0:
            continue  # block must evenly divide total length

        divisor = 10 ** (digits - block_len)
        block = n // divisor  # extract first block

        # Reconstruct number by repeating block
        repeated = 0
        for _ in range(digits // block_len):
            repeated = repeated * (10**block_len) + block

        if repeated == n:
            return True

    return False


def _is_repeated_twice_num(n: int) -> bool:
    # Count digits
    digits = 0
    tmp = n
    while tmp > 0:
        tmp //= 10
        digits += 1

    # Must be even number of digits
    if digits % 2 != 0:
        return False

    half = digits // 2
    divisor = 10**half

    left = n // divisor  # first half
    right = n % divisor  # second half

    return left == right


def sum_invalids_twice(ranges: list[tuple[int, int]]) -> int:
    return sum(
        [n for r in ranges for n in range(r[0], r[1] + 1) if _is_repeated_twice_num(n)]
    )


def sum_invalids(ranges: list[tuple[int, int]]) -> int:
    return sum(
        [n for r in ranges for n in range(r[0], r[1] + 1) if _is_repeated_num(n)]
    )


def test_sum_invalids():
    assert (
        sum_invalids(
            [
                (11, 22),
                (95, 115),
                (998, 1012),
                (1188511880, 1188511890),
                (222220, 222224),
                (1698522, 1698528),
                (446443, 446449),
                (38593856, 38593862),
                (565653, 565659),
                (824824821, 824824827),
                (2121212118, 2121212124),
            ]
        )
        == 4174379265
    )


def test_sum_invalids_twice():
    assert (
        sum_invalids_twice(
            [
                (11, 22),
                (95, 115),
                (998, 1012),
                (1188511880, 1188511890),
                (222220, 222224),
                (1698522, 1698528),
                (446443, 446449),
                (38593856, 38593862),
                (565653, 565659),
                (824824821, 824824827),
                (2121212118, 2121212124),
            ]
        )
        == 1227775554
    )


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
    # with open("./day_1_input.txt", "r") as file:
    #     lines = file.readlines()
    #     print(secret_entrance_2(lines))
    with open("./day_2_input.txt", "r") as file:
        content = file.read()
        print(sum_invalids([tuple(map(int, s.split("-"))) for s in content.split(",")]))
