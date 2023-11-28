import sys
import os
import pytest

PROJ_PATH = os.path.abspath(os.path.join(os.path.dirname(__file__), "../.."))

sys.path.append(PROJ_PATH)

from Python.tests.test_utils import load_test_data
from Python.scripts.max_consecutive_ones import count_max_consecutive_ones

test_data = load_test_data("max_consecutive_ones.json")


@pytest.mark.parametrize(
    "n, expected", [(item["n"], item["expected"]) for item in test_data]
)
def test_count_max_consecutive_ones(n, expected):
    assert count_max_consecutive_ones(n) == expected
