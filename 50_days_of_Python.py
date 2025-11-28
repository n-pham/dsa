import math


def convert_add(strings: list[str]) -> int:
    return sum([int(num_str) for num_str in strings])


def test_convert_add():
    assert convert_add(["1", "3", "5"]) == 9


def check_duplicates(strings: list[str]) -> bool:
    return len(set(strings)) == len(strings)


def test_check_duplicates():
    assert not check_duplicates(["apple", "orange", "banana", "apple"])
    assert check_duplicates(["Yoda", "Moses", "Joshua", "Mark"])


def first_longest_value(d: dict[str, str]) -> str:
    longest, longest_value = 0, ""
    for val in d.values():
        if len(val) > longest:
            longest_value = val
            longest = len(val)
    return longest_value


def test_first_longest_value():
    assert (
        first_longest_value({"val": "test", "fruit": "apple", "color": "green"})
        == "apple"
    )


def divide_or_square(num: int) -> float:
    if remainder := num % 5:
        return remainder
    return round(math.sqrt(num), 2)


def test_divide_or_square():
    assert divide_or_square(10) == 3.16
    assert divide_or_square(7) == 2
