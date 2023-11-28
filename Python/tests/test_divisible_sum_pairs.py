import sys
import os
import pytest

PROJ_PATH = os.path.abspath(os.path.join(os.path.dirname(__file__), "../.."))

sys.path.append(PROJ_PATH)

from Python.tests.test_utils import load_test_data
from Python.scripts.divisible_sum_pairs import divisible_sum_pairs

test_data = load_test_data("divisible_sum_pairs.json")


@pytest.mark.parametrize(
    "n, k, arr, expected",
    [(item["n"], item["k"], item["arr"], item["expected"]) for item in test_data],
)
def test_divisible_sum_pairs(n, k, arr, expected):
    assert divisible_sum_pairs(n, k, arr) == expected
