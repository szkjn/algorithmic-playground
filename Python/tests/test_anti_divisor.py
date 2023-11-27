import sys
import os
import pytest

PROJ_PATH = os.path.abspath(os.path.join(os.path.dirname(__file__), "../.."))

sys.path.append(PROJ_PATH)

from Python.tests.test_utils import load_test_data
from Python.scripts.anti_divisor import anti_divisor

test_data = load_test_data("anti_divisor_test_data.json")


@pytest.mark.parametrize(
    "n, expected", [(item["input"], item["expected"]) for item in test_data]
)
def test_antidivisor(n, expected):
    assert anti_divisor(n) == expected
