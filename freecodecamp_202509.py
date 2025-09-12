def too_much_screen_time(hours):
    total = 0
    for index, hour in enumerate(hours):
        if hour >= 10:
            return True
        if index >= 2 and hours[index-2] + hours[index-1] + hour >= 8*3:
            return True
        total += hour
    if total >= 6*7:
        return True
    return False

assert too_much_screen_time([1, 2, 3, 4, 5, 6, 7]) is False
assert too_much_screen_time([7, 8, 8, 4, 2, 2, 3]) is False
assert too_much_screen_time([5, 6, 6, 6, 6, 6, 6]) is False
assert too_much_screen_time([1, 2, 3, 11, 1, 3, 4]) is True
assert too_much_screen_time([1, 2, 3, 10, 2, 1, 0]) is True
assert too_much_screen_time([3, 3, 5, 8, 8, 9, 4]) is True
assert too_much_screen_time([3, 9, 4, 8, 5, 7, 6]) is True