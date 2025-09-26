from collections import Counter
import math
import re
import string


def speeding(speeds, limit):
    cnt, beyond = 0, 0
    for speed in speeds:
        if speed > limit:
            cnt += 1
            beyond += speed - limit
    return [cnt, float(beyond) / cnt if cnt > 0 else 0]


assert speeding([50, 60, 55], 60) == [0, 0]
assert speeding([58, 50, 60, 55], 55) == [2, 4]

def second_largest(arr):
    largest, second_largest = float("-inf"), float("-inf")
    for num in arr:
        if num > largest:
            second_largest = largest
            largest = num
        elif num > second_largest and num != largest:
            second_largest = num
    return second_largest

def is_perfect_square(n):
    if n < 0:
        return False
    sqrt_n = int(math.sqrt(n))
    return sqrt_n * sqrt_n == n


def is_mirror(str1, str2):
    len1, i, j = len(str1), 0, len(str2) - 1
    while i < len1 and j >= 0:
        # print(i, j, str1[i], str2[j])
        if not str1[i].isalpha():
            i += 1
        elif not str2[j].isalpha():
            j -= 1
        elif str1[i] != str2[j]:
            return False
        else:
            i += 1
            j -= 1
    while j >= 0 and not str2[j].isalpha():
        j -= 1
    if i < len1 or j >= 0:
        return False
    return True


assert is_mirror("Hello World", "dlroW olleH")
assert is_mirror("Hello World", "!dlroW !olleH")


def digits_or_letters(s):
    digit_count, letter_count = 0, 0
    for c in s:
        if c in string.digits:
            digit_count += 1
        elif c in string.ascii_letters:
            letter_count += 1
    return (
        "digits"
        if digit_count > letter_count
        else "letters"
        if letter_count > digit_count
        else "tie"
    )


def number_of_videos(video_size, video_unit, drive_size, drive_unit):
    units = {
        "B": 1,
        "KB": 1_000,
        "MB": 1_000_000,
        "GB": 1_000_000_000,
        "TB": 1_000_000_000_000,
    }
    return (drive_size * units[drive_unit]) // (video_size * units[video_unit])


assert number_of_videos(100, "MB", 1, "GB") == 10
assert number_of_videos(500, "MB", 1, "TB") == 2000
assert number_of_videos(1, "GB", 1, "TB") == 1000
assert number_of_videos(1, "TB", 1, "TB") == 1
assert number_of_videos(2, "GB", 1, "GB") == 0


def reverse_sentence(sentence):
    return " ".join(reversed(sentence.split()))


def number_of_files(file_size, file_unit, drive_size_gb):
    return (
        1_000_000_000
        * drive_size_gb
        // (
            file_size
            * {"B": 1, "KB": 1_000, "MB": 1_000_000, "GB": 1_000_000_000}[file_unit]
        )
    )


def number_of_photos(photo_size_mb, drive_size_gb):
    return drive_size_gb * 1000 // photo_size_mb


def cost_to_fill(tank_size, fuel_level, price_per_gallon):
    return "${:.2f}".format((tank_size - fuel_level) * price_per_gallon)


def generate_slug(str):
    length = len(str)
    firstIndex = 0
    while (
        firstIndex < length
        and str[firstIndex] not in string.digits
        and str[firstIndex] not in string.ascii_letters
    ):
        firstIndex += 1
    prev = str[firstIndex]
    chars = (
        [prev.lower()] if prev in string.digits or prev in string.ascii_letters else []
    )
    for index, c in enumerate(str[firstIndex + 1 :]):
        if c in string.digits or c in string.ascii_letters:
            if prev == " ":
                chars = chars + ["%20"]
            chars = chars + [c.lower()]
        prev = c
    return "".join(chars)


print(generate_slug("helloWorld") + ".")
print(generate_slug("hello world!") + ".")
print(generate_slug(" hello-world ") + ".")
print(generate_slug("hello  world") + ".")
print(generate_slug("  ?H^3-1*1]0! W[0%R#1]D  ") + ".")


def capitalize(paragraph):
    prev = paragraph[0]
    chars, ends = [prev.upper()], {".", "?", "!"}
    for c in paragraph[1:]:
        if prev in ends and c not in ends:
            chars = chars + [c.upper()]
        else:
            chars = chars + [c]
        if c != " ":
            prev = c
    return "".join(chars)


print(capitalize("this is a simple sentence."))
print(capitalize("hello world. how are you?"))
print(capitalize("i did today's coding challenge... it was fun!!"))
print(capitalize("crazy!!!strange???unconventional...sentences."))
print(
    capitalize(
        "there's a space before this period . why is there a space before that period ?"
    )
)


def get_words(paragraph):
    return [
        item[0]
        for item in Counter(re.findall(r"\b\w+\b", paragraph.lower())).most_common(3)
    ]


def adjust_thermostat(temp, target):
    return "heat" if temp < target else "cool" if temp > target else "hold"


def find_missing_numbers(arr):
    max_num, set_num = 0, set()
    for num in arr:
        if num > max_num:
            max_num = num
        set_num.add(num)
    return [num for num in range(1, max_num) if num not in set_num]


assert find_missing_numbers([1, 3, 5]) == [2, 4]
assert find_missing_numbers([1, 2, 3, 4, 5]) == []
assert find_missing_numbers([1, 10]) == [2, 3, 4, 5, 6, 7, 8, 9]
assert find_missing_numbers([10, 1, 10, 1, 10, 1]) == [2, 3, 4, 5, 6, 7, 8, 9]
assert find_missing_numbers([3, 1, 4, 1, 5, 9]) == [2, 6, 7, 8]
assert find_missing_numbers(
    [1, 2, 3, 4, 5, 7, 8, 9, 10, 12, 6, 8, 9, 3, 2, 10, 7, 4]
) == [11]


def too_much_screen_time(hours):
    total = 0
    for index, hour in enumerate(hours):
        if hour >= 10:
            return True
        if index >= 2 and hours[index - 2] + hours[index - 1] + hour >= 8 * 3:
            return True
        total += hour
    if total >= 6 * 7:
        return True
    return False


assert too_much_screen_time([1, 2, 3, 4, 5, 6, 7]) is False
assert too_much_screen_time([7, 8, 8, 4, 2, 2, 3]) is False
assert too_much_screen_time([5, 6, 6, 6, 6, 6, 6]) is False
assert too_much_screen_time([1, 2, 3, 11, 1, 3, 4]) is True
assert too_much_screen_time([1, 2, 3, 10, 2, 1, 0]) is True
assert too_much_screen_time([3, 3, 5, 8, 8, 9, 4]) is True
assert too_much_screen_time([3, 9, 4, 8, 5, 7, 6]) is True
