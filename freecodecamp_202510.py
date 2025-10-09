from datetime import datetime


def moon_phase(date_string):
    reference_date = datetime.strptime("2000-01-06", "%Y-%m-%d")
    current_date = datetime.strptime(date_string, "%Y-%m-%d")
    days_diff = (current_date - reference_date).days
    cycle_day = days_diff % 28 + 1
    print(cycle_day)
    if cycle_day < 8:
        return "New"
    elif cycle_day < 15:
        return "Waxing"
    elif cycle_day < 22:
        return "Full"
    else:
        return "Waning"


assert moon_phase("2000-01-13") == "Waxing"


def goldilocks_zone(mass):
    luminosity = mass**3.5
    start = round(0.95 * luminosity**0.5, 2)
    end = round(1.37 * luminosity**0.5, 2)
    return [start, end]


assert goldilocks_zone(1) == [0.95, 1.37]
assert goldilocks_zone(2) == [3.2, 4.61]


def find_landing_spot(matrix):
    def get_total_danger(matrix, i, j: int) -> int:
        up = matrix[i-1][j] if i >= 0 else 0
        down = matrix[i+1][j] if i < len(matrix)-1 else 0
        left = matrix[i][j-1] if j >= 0 else 0
        right = matrix[i][j+1] if j < len(matrix[0])-1 else 0
        return up + down + left + right
    safest_spot, lowest_total_danger = [-1, -1], float("inf")
    for i, row in enumerate(matrix):
        for j, spot in enumerate(row):
            if spot == 0:
                total_danger = get_total_danger(matrix, i, j)
                if total_danger < lowest_total_danger:
                    lowest_total_danger = total_danger
                    safest_spot = [i, j]
    return safest_spot


assert find_landing_spot([[1, 0], [2, 0]]) == [0, 1]
assert find_landing_spot([[9, 0, 3], [7, 0, 4], [8, 0, 5]]) == [1, 1]
assert find_landing_spot([[1, 2, 1], [0, 0, 2], [3, 0, 0]]) == [2, 2]
assert find_landing_spot([[9, 6, 0, 8], [7, 1, 1, 0], [3, 0, 3, 9], [8, 6, 0, 9]]) == [2, 1]

def send_message(route):
    seconds = 0
    for distance in route[:-1]:
        seconds += 0.5 + distance / 300_000
    seconds += route[len(route)-1] / 300_000
    return round(seconds, 4)


assert send_message([300_000, 300_000]) == 2.5


def has_exoplanet(readings):
    length, sum_level, min_level = 0, 0, 36
    for char in readings:
        length += 1
        level = int(char) if "0" <= char <= "9" else ord(char) - 65 + 10
        sum_level += level
        if min_level > level:
            min_level = level
    print(f"{length=} {sum_level=} {0.8*(sum_level / length)=} {min_level=}")
    return 0.8 * (sum_level / length) >= min_level


assert has_exoplanet("FGFFCFFGG")
assert has_exoplanet("FREECODECAMP")
assert not has_exoplanet("665544554")


def classification(temp):
    if temp >= 30_000:
        return "O"
    elif temp >= 10_000:
        return "B"
    elif temp >= 7_500:
        return "A"
    elif temp >= 6_000:
        return "F"
    elif temp >= 5_200:
        return "G"
    elif temp >= 3_700:
        return "K"
    return "M"


def check_strength_1(password):
    rules = [
        lambda p: len(p) >= 8,
        lambda p: any(char.isupper() for char in p)
        and any(char.islower() for char in p),
        lambda p: any(char.isdigit() for char in p),
        lambda p: any(char in "!@#$%^&*" for char in p),
    ]
    meets = [1 if rule(password) else 0 for rule in rules]
    return {4: "strong", 3: "medium", 2: "medium"}.get(sum(meets), "weak")


def check_strength_2(password):
    meets = {rule: 0 for rule in ["len", "both cases", "number", "special char"]}
    if len(password) >= 8:
        meets["len"] = 1
    has_lower, has_upper = False, False
    for char in password:
        if char.islower():
            has_lower = True
        elif char.isupper():
            has_upper = True
        elif char.isdigit():
            meets["number"] = 1
        elif char in "!@#$%^&*":
            meets["special char"] = 1
    if has_lower and has_upper:
        meets["both cases"] = 1
    return {4: "strong", 3: "medium", 2: "medium"}.get(sum(meets.values()), "weak")


def check_strength(password):
    has_lower, has_upper, has_digit, has_special = False, False, False, False
    for char in password:
        if char.islower():
            has_lower = True
        elif char.isupper():
            has_upper = True
        elif char.isdigit():
            has_digit = True
        elif char in "!@#$%^&*":
            has_special = True
    score = sum([len(password) >= 8, has_lower and has_upper, has_digit, has_special])
    return {4: "strong", 3: "medium", 2: "medium"}.get(score, "weak")


assert check_strength("pass!!!") == "weak"
assert check_strength("PassWord%^!") == "medium"
assert check_strength("qwerty12345") == "medium"
assert check_strength("S3cur3P@ssw0rd") == "strong"


def to_binary(decimal):
    result = ""
    while decimal > 0:
        result = str(decimal % 2) + result
        decimal //= 2
    return result


def to_decimal(binary):
    result = 0
    for i in range(len(binary)):
        result += int(binary[i]) * (2 ** (len(binary) - i - 1))
    return result
