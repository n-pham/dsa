import math


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
