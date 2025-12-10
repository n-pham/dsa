import re


def parse_bold(markdown):
    # This pattern captures:
    # 1. The opening bold markers (**, __) in group 1.
    # 2. Ensures no whitespace immediately follows the opening marker using (?!\s).
    # 3. Captures the bolded text (which can include spaces) in group 2.
    # 4. Ensures no whitespace immediately precedes the closing marker using (?<!\s).
    # 5. Matches the same closing bold markers as the opening one using \1.
    pattern = r"(\*\*|__)(?!\s)(.+?)(?<!\s)\1"
    # Replace the matched markdown bold with HTML <b> tags, using group 2 for the text.
    return re.sub(pattern, r"<b>\2</b>", markdown)


def test_parse_bold():
    assert (
        parse_bold("The **quick** brown fox __jumps__ over the **lazy** dog.")
        == "The <b>quick</b> brown fox <b>jumps</b> over the <b>lazy</b> dog."
    )
    assert parse_bold("**This is not bold **") == "**This is not bold **"


def convert_to_km(miles):
    return round(miles * 1.60934, 2)


def test_convert_to_km():
    assert convert_to_km(1) == 1.61
