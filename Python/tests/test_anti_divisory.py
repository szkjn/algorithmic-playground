import sys
import os

sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), "..")))

import pytest
from scripts.anti_divisor import antidivisor


@pytest.mark.parametrize(
    "n, expected",
    [
        (5, [2, 3]),
        (10, [3, 4, 7]),
        (105, [2, 6, 10, 11, 14, 19, 30, 42, 70]),
        (234, [4, 7, 12, 36, 52, 67, 156]),
    ],
)
def test_antidivisor(n, expected):
    assert antidivisor(n) == expected


def test_antidivisor_empty():
    assert antidivisor(1) == []
