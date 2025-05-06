import functools


def romanToInt(s: str) -> int:
    # 13
    num = 0
    singleValue = {"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
    doubleValue = {"IV": 4, "IX": 9, "XL": 40, "XC": 90, "CD": 400, "CM": 900}
    prev = ""
    for c in s + " ":
        if not prev:
            prev = c
            continue
        if prev + c in doubleValue:
            num += doubleValue[prev + c]
            prev = ""
        else:
            num += singleValue[prev]
            prev = c
    return num


assert romanToInt("LVIII") == 58
assert romanToInt("MCMXCIV") == 1994
