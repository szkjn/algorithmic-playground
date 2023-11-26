import sys
import os

sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), "..")))

from scripts.highest_candles import birthdayCakeCandles


def test_birthdayCakeCandles():
    assert birthdayCakeCandles([4, 4, 1, 3]) == 2
    assert birthdayCakeCandles([8, 8, 8, 9, 8, 8]) == 1
    assert birthdayCakeCandles([40, 2, 6, 18, 40, 5, 3, 40]) == 3
