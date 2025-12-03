import math
from typing import Any


def only_floats(a, b: Any) -> int:
    return sum([isinstance(a, float), isinstance(b, float)])


def test_only_floats():
    assert only_floats(1.23, 4) == 1


def word_index(strings: list[str]) -> int:
    longest, index_of_longest = 0, -1
    for i, string in enumerate(strings):
        if (length := len(string)) > longest:
            longest = length
            index_of_longest = i
    return index_of_longest


def test_word_index():
    assert word_index(["Hate", "remorse", "vengeance"]) == 2


def register_check(register: dict[str, str]) -> int:
    return len([0 for status in register.values() if status == "yes"])


def test_register_check():
    assert (
        register_check({"Michael": "yes", "John": "no", "Peter": "yes", "Mary": "yes"})
        == 3
    )


def lowercase_names(names: list[str]) -> tuple[str]:
    return tuple(sorted([name for name in names if name.islower()], reverse=True))


def test_lowercase_names():
    assert lowercase_names(
        ["kerry", "dickson", "John", "Mary", "carol", "Rose", "adam"]
    ) == ("kerry", "dickson", "carol", "adam")


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
