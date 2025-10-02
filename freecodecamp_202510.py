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