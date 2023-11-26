import sys
import os

sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), "..")))

import pytest
from scripts.is_isomorphic import is_isomorphic


@pytest.mark.parametrize(
    "first_str, second_str, expected",
    [
        ("abb", "cdd", True),
        ("cassidy", "1234567", False),
        ("cass", "1233", True),
    ],
)
def test_is_isomorphic(first_str, second_str, expected):
    assert is_isomorphic(first_str, second_str) == expected
