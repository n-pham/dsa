def addSpaces2109(s: str, spaces: list[int]) -> str:
    positions = [
        {"end": space_position, "start": spaces[index - 1] if index else 0}
        for index, space_position in enumerate(spaces)
    ] + [{"end": len(s), "start": spaces[-1]}]
    return " ".join([s[position["start"] : position["end"]] for position in positions])


# print(addSpaces2109("EnjoyYourCoffee", [5, 9]))

