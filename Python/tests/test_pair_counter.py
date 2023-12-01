import sys
import os
import pytest

PROJ_PATH = os.path.abspath(os.path.join(os.path.dirname(__file__), "../.."))

sys.path.append(PROJ_PATH)

from Python.tests.test_utils import load_test_data
from Python.scripts.pair_counter import pair_counter

test_data = load_test_data("pair_counter.json")


@pytest.mark.parametrize(
    "n, arr, expected",
    [(item["n"], item["arr"], item["expected"]) for item in test_data],
)
def test_pair_counter(n, arr, expected):
    assert pair_counter(n, arr) == expected
