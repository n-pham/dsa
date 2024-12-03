def addSpaces2109(s: str, spaces: list[int]) -> str:
    positions = [
        {"end": space_position, "start": spaces[index - 1] if index else 0}
        for index, space_position in enumerate(spaces)
    ] + [{"end": len(s), "start": spaces[-1]}]
    return " ".join([s[position["start"] : position["end"]] for position in positions])


# print(addSpaces2109("EnjoyYourCoffee", [5, 9]))

def minEnd3133(n: int, x: int) -> int:
    """
    | n1 |   n2  |
    +----+-------+
    ...00|10̲1110̲1| OK <-- x
    ...00|1011110|
    ...00|10̲1111̲1| OK
    ...00|11̲1110̲1| OK
    ...00|11̲1111̲1| OK
    ...01|0000000|
    ...01|0000001|
    ..............
    ...01|1011100|
    ...01|10̲1110̲1| OK
    ...01|1011110|
    ...01|10̲1111̲1| OK
    ...01|11̲1110̲1| OK
    ...01|11̲1111̲1| OK
    """
    zero_count = bin(x).count("0") - 1
    if zero_count == 0:
        return ((n-1) << x.bit_length()) | x
    n1 = (n-1) // (2**zero_count)
    n2 = (n-1) % (2**zero_count)
    list_2 = list(bin(x)[2:])
    n2_reversed_number = reversed(bin(n2)[2:])
    for i in range(len(list_2)-1, 0, -1):
        try:
            if list_2[i] == "0":
                list_2[i] = next(n2_reversed_number)
        except StopIteration:
            break
    # return int(f"{bin(n1)[2:]}{''.join(list_2)}",2)
    return f"{bin(n1)[2:]}{''.join(list_2)}"

# print(minEnd3133(6715154, 7193485))
# print(minEnd3133(4, 1))
# print(minEnd3133(3, 4))
# print(minEnd3133(2, 2))
