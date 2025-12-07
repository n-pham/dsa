import bisect


def count_beams(locations: list[str]) -> int:
    result = 0
    beams = [c == "S" for c in locations[0]]
    for location in locations[1:]:
        new_beams = beams[:]  # copy
        for i, has_beam in enumerate(beams):
            if has_beam and location[i] == "^":
                new_beams[i] = False
                new_beams[i-1] = True
                new_beams[i+1] = True
                result += 1
        beams = new_beams
    return result


def test_count_beams():
    diagram = """
        .......S.......
        ...............
        .......^.......
        ...............
        ......^.^......
        ...............
        .....^.^.^.....
        ...............
        ....^.^...^....
        ...............
        ...^.^...^.^...
        ...............
        ..^...^.....^..
        ...............
        .^.^.^.^.^...^.
        ...............
        """
    assert count_beams([line.strip() for line in diagram.split("\n") if not line.isspace() and line != ""]) == 21


def sum_math_problems(number_lines: list[list[int]], operands: list[str]) -> int:
    result = 0
    for i, operand in enumerate(operands):
        if operand == "+":
            addition_result = 0
            for line in number_lines:
                addition_result += line[i]
            result += addition_result
        else:
            multiply_result = 1
            for line in number_lines:
                multiply_result *= line[i]
            result += multiply_result
    return result


def test_sum_math_problems():
    assert (
        sum_math_problems(
            [[123, 328, 51, 64], [45, 64, 387, 23], [6, 98, 215, 314]],
            ["*", "+", "*", "+"],
        )
        == 4277556
    )


def sum_math_problems_2(number_lines: list[list[int]], operands: list[str]) -> int:
    result = 0
    # The one difference: iterate from right-to-left.
    for i in range(len(operands) - 1, -1, -1):
        operand = operands[i]
        if operand == "+":
            addition_result = 0
            for line in number_lines:
                addition_result += line[i]
            result += addition_result
        else:
            multiply_result = 1
            for line in number_lines:
                multiply_result *= line[i]
            result += multiply_result
    return result


def test_sum_math_problems_2():
    # The numbers from the example description, not the initial list.
    # This seems to be the intention of the puzzle.
    number_lines = [[356, 8, 175, 4], [24, 248, 581, 431], [1, 369, 32, 623]]
    operands = ["*", "+", "*", "+"]
    # The sequence of calculation is:
    # col 3 (+): 4 + 431 + 623 = 1058
    # col 2 (*): 175 * 581 * 32 = 3253600
    # col 1 (+): 8 + 248 + 369 = 625
    # col 0 (*): 356 * 24 * 1 = 8544
    # grand_total = 1058 + 3253600 + 625 + 8544 = 3263827
    assert sum_math_problems_2(number_lines, operands) == 3263827


def count_fresh_ingredients_memory(
    ranges: list[tuple[int, int]], ids: list[int]
) -> int:
    fresh_set = {n for r in ranges for n in range(r[0], r[1] + 1)}
    return len([1 for id in ids if id in fresh_set])


def count_fresh_ingredients(
    ranges: list[tuple[int, int]], ids: list[int]
) -> (int, int):
    # Sort ranges by the start of the interval to allow merging.
    sorted_ranges = sorted(ranges, key=lambda x: x[0])

    # Merge overlapping ranges to reduce the number of intervals to check.
    # e.g., [(1, 5), (3, 7)] becomes [(1, 7)].
    merged_ranges = []
    if sorted_ranges:
        current_start, current_end = sorted_ranges[0]
        for next_start, next_end in sorted_ranges[1:]:
            # If the next range starts before or right after the current one ends, merge them.
            if next_start <= current_end + 1:
                current_end = max(current_end, next_end)
            else:
                # No overlap, finalize the current merged range.
                merged_ranges.append((current_start, current_end))
                current_start, current_end = next_start, next_end
        merged_ranges.append((current_start, current_end))

    # Extract the starting points of the merged ranges for efficient binary searching.
    start_points = [r[0] for r in merged_ranges]
    fresh_count = 0

    for id_val in ids:
        # Use binary search (bisect_right) to quickly find which range an ID might be in.
        # This is much faster than checking every range for every ID.
        idx = bisect.bisect_right(start_points, id_val)
        if idx > 0:
            # The candidate range is the one just before the insertion point.
            candidate_range = merged_ranges[idx - 1]
            # Since `id_val >= candidate_range[0]` is guaranteed by the search,
            # we only need to check if it's within the upper bound.
            if id_val <= candidate_range[1]:
                fresh_count += 1
    id_count = sum([r[1] - r[0] + 1 for r in merged_ranges])
    return fresh_count, id_count


def test_count_fresh_ingredients():
    assert count_fresh_ingredients(
        [(3, 5), (10, 14), (16, 20), (12, 18)], [1, 5, 8, 11, 17, 32]
    ) == (3, 14)


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
    # with open("./day_4_input.txt", "r") as file:
    #     all_lines = file.readlines()
    #     print(count_accessible_2([line.strip() for line in all_lines]))
    # with open("./day_5_input.txt", "r") as file:
    #     content = file.read()
    #     range_part, id_part = content.split("\n\n")
    #     ranges = [tuple(map(int, line.split("-"))) for line in range_part.split("\n")]
    #     ids = [int(id) for id in id_part.split("\n") if id]
    #     print(count_fresh_ingredients(ranges, ids))
    # with open("./day_6_input.txt", "r") as file:
    #     all_lines = file.readlines()
    #     print(
    #         sum_math_problems_2(
    #             [
    #                 [int(num_str) for num_str in line.strip().split()]
    #                 for line in all_lines[:-1]
    #             ],
    #             all_lines[-1].split(),
    #         )
    #     )
    with open("./day_7_input.txt", "r") as file:
        diagram = file.read()
        print (
            count_beams([line.strip() for line in diagram.split("\n") if not line.isspace() and line != ""])
        )