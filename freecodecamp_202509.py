from collections import Counter
import re
import string

def number_of_files(file_size, file_unit, drive_size_gb):
    return 1_000_000_000 * drive_size_gb // (file_size * {"B": 1, "KB": 1_000, "MB": 1_000_000, "GB": 1_000_000_000}[file_unit])

def number_of_photos(photo_size_mb, drive_size_gb):
    return drive_size_gb * 1000 // photo_size_mb

def cost_to_fill(tank_size, fuel_level, price_per_gallon):
    return "${:.2f}".format((tank_size-fuel_level)*price_per_gallon)

def generate_slug(str):
    length = len(str)
    firstIndex = 0
    while firstIndex < length and str[firstIndex] not in string.digits and str[firstIndex] not in string.ascii_letters:
        firstIndex += 1
    prev = str[firstIndex]
    chars = [prev.lower()] if prev in string.digits or prev in string.ascii_letters else []
    for index, c in enumerate(str[firstIndex+1:]):
        if c in string.digits or c in string.ascii_letters:
            if prev == ' ':
                chars = chars + ["%20"]
            chars = chars + [c.lower()]
        prev = c
    return "".join(chars)

print(generate_slug("helloWorld")+'.')
print(generate_slug("hello world!")+'.')
print(generate_slug(" hello-world ")+'.')
print(generate_slug("hello  world")+'.')
print(generate_slug("  ?H^3-1*1]0! W[0%R#1]D  ")+'.')

def capitalize(paragraph):
    prev = paragraph[0]
    chars, ends = [prev.upper()], {'.', '?', '!'}
    for c in paragraph[1:]:
        if prev in ends and c not in ends:
            chars = chars + [c.upper()]
        else:
            chars = chars + [c]
        if c != ' ':
            prev = c
    return "".join(chars)

print(capitalize("this is a simple sentence."))
print(capitalize("hello world. how are you?"))
print(capitalize("i did today's coding challenge... it was fun!!"))
print(capitalize("crazy!!!strange???unconventional...sentences."))
print(capitalize("there's a space before this period . why is there a space before that period ?"))

def get_words(paragraph):
    return [item[0] for item in Counter(re.findall( r'\b\w+\b', paragraph.lower())).most_common(3)]

def adjust_thermostat(temp, target):
    return "heat" if temp < target else "cool" if temp > target else "hold"

def find_missing_numbers(arr):
    max_num, set_num = 0, set() 
    for num in arr:
        max_num = max(max_num, num)
        set_num = set_num | {num}
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
