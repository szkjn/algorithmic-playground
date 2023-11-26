import sys
import os

sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), "..")))

import pytest
from scripts.get_a1_ref_num import column_number


@pytest.mark.parametrize(
    "col_name, expected",
    [
        ("A", 1),
        ("Z", 26),
        ("AA", 27),
        ("AAA", 703),
    ],
)
def test_column_number(col_name, expected):
    assert column_number(col_name) == expected
