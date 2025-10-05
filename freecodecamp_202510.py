def has_exoplanet(readings):
    length, sum_level, min_level = 0, 0, 36
    for char in readings:
        length += 1
        level = int(char) if '0' <= char <= '9' else ord(char) - 65 + 10
        sum_level += level
        if min_level > level:
            min_level = level
    print(f"{length=} {sum_level=} {0.8*(sum_level / length)=} {min_level=}")
    return 0.8*(sum_level / length) >= min_level

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


def check_strength(password):
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
