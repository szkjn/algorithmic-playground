import sys
import os

sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), "..")))

import pytest
from zero_run_replacer import replace_zeros


@pytest.mark.parametrize(
    "digits, expected",
    [
        ("1234500362000440", "1234523623441"),
        ("123450036200044", "123452362344"),
        ("000000000000", "12"),
        ("123456789", "123456789"),
    ],
)
def test_replace_zeros(digits, expected):
    assert replace_zeros(digits) == expected
