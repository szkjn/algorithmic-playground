import sys
import os
import pytest

PROJ_PATH = os.path.abspath(os.path.join(os.path.dirname(__file__), "../.."))

sys.path.append(PROJ_PATH)

from Python.tests.test_utils import load_test_data
from Python.scripts.day_of_programmer import day_of_programmer

test_data = load_test_data("day_of_programmer.json")


@pytest.mark.parametrize(
    "year, expected", [(item["year"], item["expected"]) for item in test_data]
)
def test_anti_divisor(year, expected):
    assert day_of_programmer(year) == expected
