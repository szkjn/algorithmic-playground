import sys
import os

sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), "..")))

import pytest
from scripts.sum_every_other import sum_every_other


@pytest.mark.parametrize(
    "n, expected",
    [
        (548915381, 26),
        (10, 0),
        (1010.11, 1),
    ],
)
def test_sum_every_other(n, expected):
    assert sum_every_other(n) == expected
