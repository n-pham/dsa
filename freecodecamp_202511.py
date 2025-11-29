import datetime
import rust_regex as re
import math
import string


def calculate_age(birthday):
    year, month, day = int(birthday[:4]), int(birthday[5:7]), int(birthday[8:10])
    return 2025 - year if (month, day) <= (11, 27) else 2025 - year - 1


def test_calculate_age():
    assert calculate_age("2000-11-20") == 25
    assert calculate_age("2000-12-01") == 24


def is_fizz_buzz(sequence):
    for i, val in enumerate(sequence, start=1):
        if i % 3 == 0 and i % 5 == 0:
            expected = "FizzBuzz"
        elif i % 3 == 0:
            expected = "Fizz"
        elif i % 5 == 0:
            expected = "Buzz"
        else:
            expected = i
        if val != expected:
            return False
    return True


def generate_signature(name, title, company):
    first_letter = name[0].upper()
    prefix = ">>" if first_letter <= "I" else "--" if first_letter <= "R" else "::"
    return f"{prefix}{name}, {title} at {company}"


def build_matrix(rows, cols):
    return [[0] * cols for _ in range(rows)]


def image_search(images, term):
    return [image for image in images if re.search(term, image, re.IGNORECASE)]


def count_words(sentence):
    return sentence.count(" ") + 1


def infected(days):
    # On day 0, the first computer is infected.
    infected_computers = 1
    # Iterate through each day from 1 up to the given number of days.
    for day in range(1, days + 1):
        # Each subsequent day, the number of infected computers doubles.
        infected_computers *= 2
        # Every 3rd day, a patch is applied.
        if day % 3 == 0:
            # Calculate the number of patched computers (20% of infected).
            patched_computers = math.ceil(infected_computers * 0.20)
            # Reduce the number of infected computers.
            infected_computers -= patched_computers
    # Return the total number of infected computers.
    return infected_computers


def verify(message, key, signature):
    message_sum = sum(string.ascii_letters.find(char) + 1 for char in message)
    key_sum = sum(string.ascii_letters.find(char) + 1 for char in key)
    return (message_sum + key_sum) == signature


def get_weekday(date_string):
    return datetime.datetime.strptime(date_string, "%Y-%m-%d").strftime("%A")


def combinations(cards):
    return math.comb(52, cards)


def can_post(message):
    length = len(message)
    return (
        "short post"
        if length <= 40
        else "long post"
        if length <= 80
        else "invalid post"
    )


def get_extension(filename):
    position = filename.rfind(".")
    return (
        filename[position + 1 :]
        if position > -1 and position < len(filename) - 1
        else "none"
    )


def days_until_weekend(date_string):
    date_obj = datetime.datetime.strptime(date_string, "%Y-%m-%d").date()
    weekday = date_obj.weekday()
    if weekday == 5 or weekday == 6:
        return "It's the weekend!"
    else:
        days_to_saturday = 5 - weekday
        if days_to_saturday == 1:
            return "1 day until the weekend."
        else:
            return f"{days_to_saturday} days until the weekend."


def find_word(matrix, word):
    rows = len(matrix)
    cols = len(matrix[0])
    word_len = len(word)
    reversed_word = word[::-1]

    # Horizontal search
    for r in range(rows):
        row_str = "".join(matrix[r])

        # Left to right
        idx = row_str.find(word)
        if idx != -1:
            return [[r, idx], [r, idx + word_len - 1]]

        # Right to left
        idx = row_str.find(reversed_word)
        if idx != -1:
            return [[r, idx + word_len - 1], [r, idx]]

    # Vertical search
    for c in range(cols):
        col_str = "".join(matrix[r][c] for r in range(rows))

        # Top to bottom
        idx = col_str.find(word)
        if idx != -1:
            return [[idx, c], [idx + word_len - 1, c]]

        # Bottom to top
        idx = col_str.find(reversed_word)
        if idx != -1:
            return [[idx + word_len - 1, c], [idx, c]]


def test_build_matrix():
    assert build_matrix(2, 3) == [[0, 0, 0], [0, 0, 0]]

try:
    from hypothesis import given
    from hypothesis.strategies import integers

    @given(integers(0, 20))
    def test_infected_range(n):
        assert infected(n) > n
except ImportError:
    pass

def test_infected():
    assert infected(0) == 1
    assert infected(1) == 2
    assert infected(2) == 4
    assert infected(3) == 6
    assert infected(4) == 12
    assert infected(5) == 24
    assert infected(6) == 38


def test_verify():
    assert verify("foo", "bar", 57)
    assert verify("Foo", "Bar", 109)
    assert verify("Hello!", "World?", 176)
    assert verify("abc", "def", 21)
    assert not verify("abc", "def", 20)
    assert verify("", "", 0)
    assert verify("123", "!@#", 0)


def test_combinations():
    assert combinations(0) == 1
    assert combinations(1) == 52
    assert combinations(2) == 1326
    assert combinations(52) == 1


def test_find_word():
    matrix = [
        ["f", "o", "a", "m"],
        ["o", "b", "q", "p"],
        ["a", "o", "b", "a"],
        ["m", "p", "a", "l"],
    ]
    # left to right
    assert find_word(matrix, "foam") == [[0, 0], [0, 3]]
    # right to left
    assert find_word([list("maof")], "foam") == [[0, 3], [0, 0]]
    # top to bottom
    matrix_vert = [
        ["f", "x", "x"],
        ["o", "x", "x"],
        ["a", "x", "x"],
        ["m", "x", "x"],
    ]
    assert find_word(matrix_vert, "foam") == [[0, 0], [3, 0]]
    # bottom to top
    matrix_vert_rev = [
        ["m", "x", "x"],
        ["a", "x", "x"],
        ["o", "x", "x"],
        ["f", "x", "x"],
    ]
    assert find_word(matrix_vert_rev, "foam") == [[3, 0], [0, 0]]

    matrix2 = [
        ["a", "b", "c", "d", "e"],
        ["f", "g", "h", "i", "j"],
        ["k", "l", "m", "n", "o"],
        ["p", "q", "r", "s", "t"],
        ["u", "v", "w", "x", "y"],
    ]
    assert find_word(matrix2, "abc") == [[0, 0], [0, 2]]
    assert find_word(matrix2, "edc") == [[0, 4], [0, 2]]
    assert find_word(matrix2, "afkpu") == [[0, 0], [4, 0]]
    assert find_word(matrix2, "yto") == [[4, 4], [2, 4]]
    assert find_word(matrix2, "joty") == [[1, 4], [4, 4]]


def test_days_until_weekend():
    assert days_until_weekend("2025-11-14") == "1 day until the weekend."  # Friday
    assert days_until_weekend("2025-11-15") == "It's the weekend!"  # Saturday
    assert days_until_weekend("2025-11-16") == "It's the weekend!"  # Sunday
    assert days_until_weekend("2025-11-10") == "5 days until the weekend."  # Monday
    assert days_until_weekend("2025-11-13") == "2 days until the weekend."  # Thursday
    assert days_until_weekend("2025-11-12") == "3 days until the weekend."  # Wednesday
    assert days_until_weekend("2025-11-11") == "4 days until the weekend."  # Tuesday
