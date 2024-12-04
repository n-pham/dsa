import functools  # noqa: F401


def addSpaces2109(s: str, spaces: list[int]) -> str:
    positions = [
        {"end": space_position, "start": spaces[index - 1] if index else 0}
        for index, space_position in enumerate(spaces)
    ] + [{"end": len(s), "start": spaces[-1]}]
    return " ".join([s[position["start"] : position["end"]] for position in positions])


# print(addSpaces2109("EnjoyYourCoffee", [5, 9]))


def minEnd3133(n: int, x: int) -> int:
    """
    | n1 |   n2  |
    +----+-------+
    ...00|10̲1110̲1| OK <-- x
    ...00|1011110|
    ...00|10̲1111̲1| OK
    ...00|11̲1110̲1| OK
    ...00|11̲1111̲1| OK
    ...01|0000000|
    ...01|0000001|
    ..............
    ...01|1011100|
    ...01|10̲1110̲1| OK
    ...01|1011110|
    ...01|10̲1111̲1| OK
    ...01|11̲1110̲1| OK
    ...01|11̲1111̲1| OK
    """
    zero_count = bin(x).count("0") - 1
    if zero_count == 0:
        return ((n - 1) << x.bit_length()) | x
    n1 = (n - 1) // (2**zero_count)
    n2 = (n - 1) % (2**zero_count)
    list_2 = list(bin(x)[2:])
    n2_reversed_number = reversed(bin(n2)[2:])
    for i in range(len(list_2) - 1, 0, -1):
        try:
            if list_2[i] == "0":
                list_2[i] = next(n2_reversed_number)
        except StopIteration:
            break
    # return int(f"{bin(n1)[2:]}{''.join(list_2)}",2)
    return f"{bin(n1)[2:]}{''.join(list_2)}"


# print(minEnd3133(6715154, 7193485))
# print(minEnd3133(4, 1))
# print(minEnd3133(3, 4))
# print(minEnd3133(2, 2))


def largestCombination2275(candidates: list[int]) -> int:
    """
      10000
      10001
    1000111
     111110
       1100
      11000
       1110
    -------
    1144432
    """
    largest = 0
    for i in range(max(candidates).bit_length()):
        largest = max(largest, sum([1 for number in candidates if number & (1 << i)]))
        # largest = max(largest, len(list(filter(lambda number: number & (1 << i), candidates))))
        # largest = max(largest, functools.reduce(lambda cnt, number: cnt +1 if number & (1 << i) else cnt, candidates, 0))
    return largest


# print(largestCombination2275([16, 17, 71, 62, 12, 24, 14]))
# print(largestCombination2275([8, 8]))


def firstMissingPositive_t(nums: list[int]) -> int:
    total = 0
    count = 0
    has_number_1 = False
    for number in nums:
        if number > 0:
            count += 1
            total += number
            if number == 1:
                has_number_1 = True
    if has_number_1:
        expected_total_1_to_n = (count + 2) * (count + 1) // 2
        if expected_total_1_to_n > total:
            return expected_total_1_to_n - total
        else:
            return 'No'
    else:
        return 1


def firstMissingPositive(nums: list[int]) -> int:
    """
    if not has_number_1:
        return 1
    else:
        curr    min_1   min_2 > min_1 + 1
       ----------------------
           3        3    2^31
           4        4    2^31
          -1        4    2^31
           1        1       4
       ----------------------
       10000    10000    2^31
           4        4   10000
        4000        4    4000
           2        2       4
           3        2       4
           1        3       3
    """
    min_1 = pow(2, 31)
    min_2 = pow(2, 31)
    for number in nums:
        if number > 0:
            if number < min_1:
                min_1 = number
            elif number > min_1 and min_2 > number + 1:
                min_1 = number
    return min_1, min_2

print(firstMissingPositive([3, 4, -1, 1]))
print(firstMissingPositive([1, 2, 0]))
print(firstMissingPositive([7,8,9,11,12]))
print(firstMissingPositive([100000, 3, 4000, 2, 15, 1, 99999]))
