def convert_to_km(miles):
    return round(miles*1.60934, 2)


def test_convert_to_km():
    assert convert_to_km(1) == 1.61
