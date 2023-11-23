from anti_divisor import antidivisor

def test_antidivisor():
    assert antidivisor(5) == [2, 3]
    assert antidivisor(10) == [3, 4, 7]
    assert antidivisor(105) == [2, 6, 10, 11, 14, 19, 30, 42, 70]
    assert antidivisor(234) == [4, 7, 12, 36, 52, 67, 156]

def test_antidivisor_empty():
    assert antidivisor(1) == []
