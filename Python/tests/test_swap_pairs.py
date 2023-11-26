import sys
import os

sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), "..")))

import pytest
from scripts.swap_pairs import swap_pairs


@pytest.mark.parametrize(
    "arr, expected",
    [
        ([1, 2, 3, 4], [2, 1, 4, 3]),
        ([], []),
        ([9, 22, 8], [22, 9, 8]),
    ],
)
def test_swap_pairs(arr, expected):
    assert swap_pairs(arr) == expected
