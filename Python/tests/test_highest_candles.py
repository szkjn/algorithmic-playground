import sys
import os

sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), "..")))

import pytest
from scripts.highest_candles import birthday_cake_candles


@pytest.mark.parametrize(
    "candles, expected",
    [
        ([4, 4, 1, 3], 2),
        ([8, 8, 8, 9, 8, 8], 1),
        ([40, 2, 6, 18, 40, 5, 3, 40], 3),
    ],
)
def test_birthday_cake_candles(candles, expected):
    assert birthday_cake_candles(candles) == expected
