import datetime
import re
import math
import string


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
    return datetime.datetime.strptime(date_string, "%Y-%m-%d").strftime('%A')


def combinations(cards):
    return math.comb(52, cards)

def can_post(message):
    length = len(message)
    return "short post" if length <= 40 else "long post" if length <= 80 else "invalid post"


def test_build_matrix():
    assert build_matrix(2, 3) == [[0, 0, 0], [0, 0, 0]]


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
