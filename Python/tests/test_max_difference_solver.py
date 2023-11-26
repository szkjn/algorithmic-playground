import sys
import os

sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), "..")))

import pytest
from scripts.max_difference_solver import Difference


@pytest.mark.parametrize(
    "input_list, expected",
    [
        ([1, 2, 5], 4),
        ([10, 10, 10, 10, 10], 0),
        ([1, 9, 3, 7, 6, 2, 1], 8),
        ([], 0),
    ],
)
def test_compute_difference(input_list, expected):
    d = Difference(input_list)
    d.computeDifference()
    assert d.maximumDifference == expected
