def count_accessible(diagram: list[str]) -> int:
    """
    Counts the number of accessible rolls of paper ('@') in a diagram.
    A roll of paper is accessible if there are fewer than four rolls of paper
    in its eight adjacent positions (horizontally, vertically, and diagonally).
    """
    accessible_count = 0
    rows = len(diagram)
    cols = len(diagram[0])

    # Relative coordinates for 8 neighbors
    dr = [-1, -1, -1, 0, 0, 1, 1, 1]
    dc = [-1, 0, 1, -1, 1, -1, 0, 1]

    for r in range(rows):
        for c in range(cols):
            if diagram[r][c] == "@":
                neighbor_a_count = 0
                for i in range(8):
                    nr, nc = r + dr[i], c + dc[i]

                    # Check if neighbor is within bounds
                    if 0 <= nr < rows and 0 <= nc < cols:
                        if diagram[nr][nc] == "@":
                            neighbor_a_count += 1

                if neighbor_a_count < 4:
                    accessible_count += 1
    return accessible_count


def test_count_accessible():
    # Test case to match the user's expected count of 13 accessible '@'s.
    # Each '@' in this diagram has 0 '@' neighbors, making all of them accessible.
    thirteen_accessible_diagram = [
        list("@.@.@.@.@"),  # 5 '@'s
        list("..........."),
        list("@.@.@.@.@"),  # 5 '@'s
        list("..........."),
        list("@.@.@......"),  # 3 '@'s
    ]
    assert count_accessible(thirteen_accessible_diagram) == 13


def count_accessible_2(diagram: list[str]) -> int:
    """
    Counts the number of accessible rolls of paper ('@') in a diagram.
    A roll of paper is accessible if there are fewer than four rolls of paper
    in its eight adjacent positions (horizontally, vertically, and diagonally).
    Accessible rolls are removed, and the process is repeated until no more
    rolls can be removed.
    """
    total_accessible_count = 0
    # Make a mutable copy of the diagram
    mutable_diagram = [list(row) for row in diagram]
    rows = len(mutable_diagram)
    if rows == 0:
        return 0
    cols = len(mutable_diagram[0])

    # Relative coordinates for 8 neighbors
    dr = [-1, -1, -1, 0, 0, 1, 1, 1]
    dc = [-1, 0, 1, -1, 1, -1, 0, 1]

    while True:
        accessible_this_round = []
        for r in range(rows):
            for c in range(cols):
                if mutable_diagram[r][c] == "@":
                    neighbor_a_count = 0
                    for i in range(8):
                        nr, nc = r + dr[i], c + dc[i]
                        # Check if neighbor is within bounds
                        if 0 <= nr < rows and 0 <= nc < cols:
                            if mutable_diagram[nr][nc] == "@":
                                neighbor_a_count += 1

                    if neighbor_a_count < 4:
                        accessible_this_round.append((r, c))

        if not accessible_this_round:
            break

        total_accessible_count += len(accessible_this_round)
        for r, c in accessible_this_round:
            mutable_diagram[r][c] = "."

    return total_accessible_count


def test_count_accessible_2():
    diagram = [
        "@@@",
        "@.@",
        "@@@",
    ]
    assert count_accessible_2(diagram) == 8

    thirteen_accessible_diagram = [
        list("@.@.@.@.@"),
        list("..........."),
        list("@.@.@.@.@"),
        list("..........."),
        list("@.@.@......"),
    ]
    assert count_accessible_2(thirteen_accessible_diagram) == 13

    all_connected = [
        "@@",
        "@@",
    ]
    assert count_accessible_2(all_connected) == 4

    cascade_diagram = [
        "@@@@@",
        ".....",
        "@@@@@",
    ]
    assert count_accessible_2(cascade_diagram) == 10


def max_joltage(bank: str) -> int:
    max_ten = "0"
    max_ten_index = 0
    for i, digit in enumerate(bank[:-1]):
        if digit > max_ten:
            max_ten = digit
            max_ten_index = i
    max_one = "0"
    for i in range(max_ten_index + 1, len(bank)):
        if bank[i] > max_one:
            max_one = bank[i]
    return int(max_ten) * 10 + int(max_one)


def max_joltage_2(bank: str) -> int:
    k = 12
    result_chars = []
    start_index = 0
    for i in range(k):
        remaining_to_pick = k - i
        end_index = len(bank) - remaining_to_pick + 1
        window = bank[start_index:end_index]
        max_digit = max(window)
        best_index = bank.find(max_digit, start_index, end_index)
        result_chars.append(max_digit)
        start_index = best_index + 1
    return int("".join(result_chars))


def test_max_joltage():
    assert max_joltage("987654321111111") == 98
    assert max_joltage("811111111111119") == 89
    assert max_joltage("234234234234278") == 78
    assert max_joltage("818181911112111") == 92


def test_max_joltage_2():
    assert max_joltage_2("192837465564738291") == 987664738291


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
    # with open("./day_2_input.txt", "r") as file:
    #     content = file.read()
    #     print(sum_invalids([tuple(map(int, s.split("-"))) for s in content.split(",")]))
    # with open("./day_3_input.txt", "r") as file:
    #     all_lines = file.readlines()
    #     print(sum([max_joltage_2(line.strip()) for line in all_lines]))
    with open("./day_4_input.txt", "r") as file:
        all_lines = file.readlines()
        print(count_accessible_2([line.strip() for line in all_lines]))
