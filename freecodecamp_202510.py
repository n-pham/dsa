from datetime import datetime
import string
import multiprocessing
import traceback


def validate(email):
    partIndex = email.find("@")
    if partIndex == -1 or email.find("@", partIndex+1) > -1:
        return False
    local, domain = email[:partIndex], email[partIndex+1:]
    allowedLocal = set(string.ascii_letters + string.digits + '.' + '_' + '-')
    if local[0] == '.' or local[-1] == '.' or any(c not in allowedLocal for c in local) or local.find("..") > -1:
        return False
    domainDotIndex = domain.rfind('.')
    if domainDotIndex == -1 or domain.find("..") > -1:
        return False
    topLevelDomain = domain[domainDotIndex+1:]
    if any(c not in string.ascii_letters for c in topLevelDomain):
        return False
    return True


def strip_tags(html):
    first = html.find('>')
    if first > - 1 and first < len(html)-1:
        last = html.find('<', first)
        # print(html[first+1:last], "*", html[last:])
        return html[first+1:last] + strip_tags(html[last:])
    return ''


def count(text, parameter):
    return text.count(parameter)

def count2(text, parameter):
    times = 0
    len_param = len(parameter)
    for i in range(len(text)-len_param+1):
        if text[i:i+len_param] == parameter:
            times += 1
    return times

def to_12(time):
    suffix = "AM"
    hour = int(time[:2])
    if hour >= 12:
        hour -= 12
        suffix = "PM"
    if hour == 0:
        hour = 12
    return f"{hour}:{time[-2:]} {suffix}"

def battle(our_team, opponent):
    def get_char_value(char: str) -> int:
        return 2 * (ord(char) - 64) if char <= "Z" else ord(char) - 96

    def get_word_value(word: str) -> int:
        return sum([get_char_value(char) for char in word])

    opponent_words = opponent.split(" ")
    our_score, opponent_score = 0, 0
    for i, our_word in enumerate(our_team.split(" ")):
        our_word_value, opponent_word_value = (
            get_word_value(our_word),
            get_word_value(opponent_words[i]),
        )
        # print(f"{our_word_value=} {opponent_word_value=}")
        if our_word_value > opponent_word_value:
            our_score += 1
        elif opponent_word_value > our_word_value:
            opponent_score += 1
    return (
        "We win"
        if our_score > opponent_score
        else "We lose"
        if our_score < opponent_score
        else "Draw"
    )

def hex_to_decimal(hex):
    result = 0
    for i in range(len(hex)):
        char = hex[i]
        val = ord(char) - 48 if char <= "9" else ord(char) - 65 + 10
        result += val * (16 ** (len(hex) - i - 1))
    return result

def launch_fuel(payload):
    payload = float(payload)
    fuel = 0.0
    while True:
        fuel_needed = (payload + fuel) / 5.0
        additional_fuel = fuel_needed - fuel
        if additional_fuel < 1:
            return round(fuel_needed, 1)
        fuel = fuel_needed

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

def goldilocks_zone(mass):
    luminosity = mass**3.5
    start = round(0.95 * luminosity**0.5, 2)
    end = round(1.37 * luminosity**0.5, 2)
    return [start, end]

def find_landing_spot(matrix):
    def get_total_danger(matrix, i, j: int) -> int:
        up = matrix[i - 1][j] if i >= 0 else 0
        down = matrix[i + 1][j] if i < len(matrix) - 1 else 0
        left = matrix[i][j - 1] if j >= 0 else 0
        right = matrix[i][j + 1] if j < len(matrix[0]) - 1 else 0
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

def send_message(route):
    seconds = 0
    for distance in route[:-1]:
        seconds += 0.5 + distance / 300_000
    seconds += route[len(route) - 1] / 300_000
    return round(seconds, 4)

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

# --- Test Functions ---

def test_validate():
    assert validate("a@b.cd")
    assert validate("example@test.c0") is False
    assert validate("develop.ment_user@c0D!NG.R.CKS")

def test_strip_tags():
    assert strip_tags('<p class="center">Hello <b>World</b>!</p>') == "Hello World!"

def test_to_12():
    assert to_12("1124") == "11:24 AM"

def test_battle():
    assert battle("We must never surrender", "Our team must win") == "Draw"

def test_launch_fuel():
    assert launch_fuel(50) == 12.4
    assert launch_fuel(243) == 60.7

def test_moon_phase():
    assert moon_phase("2000-01-13") == "Waxing"

def test_goldilocks_zone():
    assert goldilocks_zone(1) == [0.95, 1.37]
    assert goldilocks_zone(2) == [3.2, 4.61]

def test_find_landing_spot():
    assert find_landing_spot([[1, 0], [2, 0]]) == [0, 1]
    assert find_landing_spot([[9, 0, 3], [7, 0, 4], [8, 0, 5]]) == [1, 1]
    assert find_landing_spot([[1, 2, 1], [0, 0, 2], [3, 0, 0]]) == [2, 2]
    assert find_landing_spot([[9, 6, 0, 8], [7, 1, 1, 0], [3, 0, 3, 9], [8, 6, 0, 9]]) == [2, 1]

def test_send_message():
    assert send_message([300_000, 300_000]) == 2.5

def test_has_exoplanet():
    assert has_exoplanet("FGFFCFFGG")
    assert has_exoplanet("FREECODECAMP")
    assert not has_exoplanet("665544554")

def test_check_strength():
    assert check_strength("pass!!!") == "weak"
    assert check_strength("PassWord%^!") == "medium"
    assert check_strength("qwerty12345") == "medium"
    assert check_strength("S3cur3P@ssw0rd") == "strong"

def run_test(test_func):
    try:
        test_func()
        return (test_func.__name__, "PASSED", None)
    except Exception:
        return (test_func.__name__, "FAILED", traceback.format_exc())

tests = [
    test_validate,
    test_strip_tags,
    test_to_12,
    test_battle,
    test_launch_fuel,
    test_moon_phase,
    test_goldilocks_zone,
    test_find_landing_spot,
    test_send_message,
    test_has_exoplanet,
    test_check_strength,
]

if __name__ == "__main__":
    with multiprocessing.Pool() as pool:
        results = pool.map(run_test, tests)

    all_passed = True
    for res in results:
        print(f"{res[0]}: {res[1]}")
        if res[1] == "FAILED":
            all_passed = False
            print(f"  Error: {res[2]}")

    if all_passed:
        print("\nAll tests passed.")
    else:
        print("\nSome tests failed.")