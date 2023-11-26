import sys
import os

sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), "..")))

from scripts.combine_strings import combine_strings


def test_combine_strings():
    assert combine_strings(["a", "b", "c", "d", "e", "f", "g"], 5) == [
        "a b c",
        "d e f",
        "g",
    ]
    assert combine_strings(["a", "b", "c", "d", "e", "f", "g"], 12) == [
        "a b c d e f",
        "g",
    ]
    assert combine_strings(["alpha", "beta", "gamma", "delta", "epsilon"], 20) == [
        "alpha beta gamma",
        "delta epsilon",
    ]
