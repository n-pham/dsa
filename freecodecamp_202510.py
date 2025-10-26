from datetime import datetime
import re
import string
import traceback


def is_spam(number):
    parts = re.split(r"\s|-", number)
    if len(parts[0]) > 3 or (len(parts[0]) > 1 and parts[0][1] != "0"):
        return True
    if not "200" <= parts[1].strip("()") <= "900":
        return True
    local_number_str = "".join(parts[2:])
    if len(local_number_str) == 7:
        first_three_sum = sum(int(d) for d in local_number_str[:3])
        if str(first_three_sum) in local_number_str[-4:]:
            return True
    all_digits = "".join(c for c in number if c.isdigit())
    if re.search(r"(\d)\1{3,}", all_digits):
        return True
    return False


def favorite_songs(playlist):
    first_tile, second_title, first_cnt, second_cnt = "", "", 0, 0
    for item in playlist:
        cnt = item["plays"]
        if cnt > first_cnt:
            second_cnt, second_title = first_cnt, first_tile
            first_cnt, first_tile = cnt, item["title"]
        elif cnt > second_cnt:
            second_cnt, second_title = cnt, item["title"]
    return [first_tile, second_title]


def test_favorite_songs():
    assert favorite_songs(
        [
            {"title": "Sync or Swim", "plays": 3},
            {"title": "Byte Me", "plays": 1},
            {"title": "Earbud Blues", "plays": 2},
        ]
    ) == ["Sync or Swim", "Earbud Blues"]
    assert favorite_songs(
        [
            {"title": "Skip Track", "plays": 98},
            {"title": "99 Downloads", "plays": 99},
            {"title": "Clickwheel Love", "plays": 100},
        ]
    ) == ["Clickwheel Love", "99 Downloads"]
    assert favorite_songs(
        [
            {"title": "Song A", "plays": 42},
            {"title": "Song B", "plays": 99},
            {"title": "Song C", "plays": 75},
        ]
    ) == ["Song B", "Song C"]


def format(seconds):
    second_s = f"{seconds % 60:02}"
    hour = seconds // 3_600
    hour_s = f"{hour}:" if hour > 0 else ""
    minute = seconds % 3_600 // 60
    minute_s = f"{minute:02}:" if hour > 0 else f"{minute}:"
    return f"{hour_s}{minute_s}{second_s}"


def test_format():
    assert format(500) == "8:20"
    assert format(4000) == "1:06:40"
    assert format(99999) == "27:46:39"
    assert format(1) == "0:01"


def dive(map, coordinates):
    foundCount, notFoundCount = 0, 0
    for row in map:
        for cell in row:
            if cell == "O":
                notFoundCount += 1
            elif cell == "X":
                foundCount += 1
    x, y = coordinates[0], coordinates[1]
    if map[x][y] == "-":
        return "Empty"
    if map[x][y] == "O":
        notFoundCount -= 1
    return "Found" if notFoundCount > 0 else "Recovered"


def test_dive():
    assert dive([["-", "X"], ["-", "X"], ["-", "O"]], [2, 0]) == "Empty"
    assert dive([["-", "X"], ["-", "X"], ["-", "O"]], [2, 1]) == "Recovered"


def complementary_dna(strand):
    pairs = {"A": "T", "T": "A", "C": "G", "G": "C"}
    result = ""
    for char in strand:
        result += pairs[char]
    return result


def test_complementary_dna():
    assert complementary_dna("ATTGC") == "TAACG"
    assert complementary_dna("GTAT") == "CATA"


def wise_speak(sentence):
    split_words = ["have", "must", "are", "will", "can"]
    punctuation = sentence[-1]
    words = sentence.split(" ")
    i = next(i for i, word in enumerate(words) if word in split_words)
    return " ".join(
        [words[i + 1].capitalize()]
        + words[i + 2 : -1]
        + [words[-1][:-1] + ","]
        + [words[0][0].lower() + words[0][1:]]
        + words[1:i]
        + [words[i] + punctuation]
    )


def test_wise_speak():
    assert wise_speak("You must speak wisely.") == "Speak wisely, you must."
    assert (
        wise_speak("Do you think you will complete this?")
        == "Complete this, do you think you will?"
    )


def array_diff(arr1, arr2):
    arr, arr1, arr2 = [], sorted(arr1), sorted(arr2)
    i1, i2, len1, len2 = 0, 0, len(arr1), len(arr2)
    while i1 < len1 and i2 < len2:
        if arr1[i1] < arr2[i2]:
            arr.append(arr1[i1])
            i1 += 1
        elif arr1[i1] > arr2[i2]:
            arr.append(arr2[i2])
            i2 += 1
        else:
            i1, i2 = i1 + 1, i2 + 1
    if i1 < len1:
        arr.extend(arr1[i1:])
    if i2 < len2:
        arr.extend(arr2[i2:])
    return arr


def test_array_diff():
    assert array_diff(["apple", "banana"], ["apple", "banana", "cherry"]) == ["cherry"]


def adjust_thermostat(current_f, target_c):
    diff_f = (target_c * 1.8) + 32 - current_f
    return (
        f"Heat: {diff_f:.1f} degrees Fahrenheit"
        if diff_f > 0
        else f"Cool: {-diff_f:.1f} degrees Fahrenheit"
        if diff_f < 0
        else "Hold"
    )


def calculate_tips(meal_price, custom_tip):
    price = float(meal_price[1:])
    return [
        f"${0.15 * price:.2f}",
        f"${0.2 * price:.2f}",
        f"${float(custom_tip[:-1]) / 100 * price:.2f}",
    ]


def test_calculate_tips():
    assert calculate_tips("$19.85", "9%") == ["$2.98", "$3.97", "$1.79"]
    assert calculate_tips("$10.00", "25%") == ["$1.50", "$2.00", "$2.50"]


def extract_attributes(element):
    attribute_str = element[
        element.find(" ") : min(element.find("/") - 1, element.find(">"))
    ]
    parts = attribute_str.split('"')
    keys_parts = parts[0::2]
    values = parts[1::2]
    keys = [p.strip().rstrip("=") for p in keys_parts if p.strip()]
    return [f"{k}, {v}" for k, v in zip(keys, values)]


def test_extract_attributes():
    assert extract_attributes(
        '<input name="email" type="email" required="true" />'
    ) == ["name, email", "type, email", "required, true"]
    assert extract_attributes('<span class="red"></span>') == ["class, red"]
    assert extract_attributes(
        '<button id="submit" class="btn btn-primary">Submit</button>'
    ) == ["id, submit", "class, btn btn-primary"]


def sock_pairs(pairs, cycles):
    number = 2 * pairs
    for i in range(1, cycles + 1):
        if i % 2 == 0 and number:
            number -= 1
        if i % 3 == 0:
            number += 1
        if i % 5 == 0 and number:
            number -= 1
        if i % 10 == 0:
            number += 2
    return number // 2


def test_sock_pairs():
    assert sock_pairs(2, 5) == 1


def mask(card):
    return f"****{card[4]}****{card[9]}****{card[14]}{card[-4:]}"


def validate(email):
    partIndex = email.find("@")
    if partIndex == -1 or email.find("@", partIndex + 1) > -1:
        return False
    local, domain = email[:partIndex], email[partIndex + 1 :]
    allowedLocal = set(string.ascii_letters + string.digits + "." + "_" + "-")
    if (
        local[0] == "."
        or local[-1] == "."
        or any(c not in allowedLocal for c in local)
        or local.find("..") > -1
    ):
        return False
    domainDotIndex = domain.rfind(".")
    if domainDotIndex == -1 or domain.find("..") > -1:
        return False
    topLevelDomain = domain[domainDotIndex + 1 :]
    if any(c not in string.ascii_letters for c in topLevelDomain):
        return False
    return True


def strip_tags(html):
    first = html.find(">")
    if first > -1 and first < len(html) - 1:
        last = html.find("<", first)
        # print(html[first+1:last], "*", html[last:])
        return html[first + 1 : last] + strip_tags(html[last:])
    return ""


def count(text, parameter):
    return text.count(parameter)


def count2(text, parameter):
    times = 0
    len_param = len(parameter)
    for i in range(len(text) - len_param + 1):
        if text[i : i + len_param] == parameter:
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


def test_is_spam():
    # Valid numbers
    assert not is_spam("+0 (200) 234-0182")

    # Invalid area code
    assert is_spam("+0 (199) 234-0182")
    assert is_spam("+0 (901) 234-0182")

    # Sum of first three in last four
    # 2+3+4 = 9. '9' is in '0192'.
    assert is_spam("+0 (200) 234-0192")
    # 1+2+3 = 6
    assert not is_spam("+0 (200) 123-4578")
    assert is_spam("+0 (200) 123-4568")


    # 4 consecutive digits
    assert is_spam("+0 (200) 111-1182")
    assert is_spam("+0 (200) 234-4444")
    assert is_spam("+0 (200) 4444-0182")
    assert is_spam("1234444567")


    # Not spam
    assert not is_spam("+0 (800) 555-1212")


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
    assert find_landing_spot(
        [[9, 6, 0, 8], [7, 1, 1, 0], [3, 0, 3, 9], [8, 6, 0, 9]]
    ) == [2, 1]


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
    test_is_spam,
    test_complementary_dna,
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
