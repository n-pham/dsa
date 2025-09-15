def adjust_thermostat(temp, target):
    return "heat" if temp < target else "cool" if temp > target else "hold"

def find_missing_numbers(arr):
    result, maxNum, setNum = [], 0, set()
    for num in arr:
        maxNum = max(maxNum, num)
        setNum = setNum | {num}
    for num in range(1, maxNum):
        if not num in setNum:
            result = result + [num]
    return result


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
