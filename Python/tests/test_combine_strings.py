import sys
import os

sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), "..")))

import pytest
from scripts.combine_strings import combine_strings


@pytest.mark.parametrize(
    "arr, n, expected",
    [
        (["a", "b", "c", "d", "e", "f", "g"], 5, ["a b c", "d e f", "g"]),
        (
            ["a", "b", "c", "d", "e", "f", "g"],
            12,
            [
                "a b c d e f",
                "g",
            ],
        ),
        (
            ["alpha", "beta", "gamma", "delta", "epsilon"],
            20,
            [
                "alpha beta gamma",
                "delta epsilon",
            ],
        ),
    ],
)
def test_combine_strings(arr, n, expected):
    assert combine_strings(arr, n) == expected
