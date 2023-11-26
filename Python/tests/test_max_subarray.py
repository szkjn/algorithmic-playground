import sys
import os

sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), "..")))

import pytest
from max_subarray import max_subarray


@pytest.mark.parametrize(
    "arr, n, expected",
    [
        ([-4, 2, -5, 1, 2, 3, 6, -5, 1], 4, [1, 2, 3, 6]),
        ([1, 2, 0, 5], 2, [0, 5]),
        ([8, 32, -13, -22, 8, -5, -6, -25, 6, 8], 3, [8, 32, -13]),
    ],
)
def test_max_subarray(arr, n, expected):
    assert max_subarray(arr, n) == expected
