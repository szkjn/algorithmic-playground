import sys
import os

sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), "..")))

from get_a1_ref_num import column_number


def test_column_number():
    assert column_number("A") == 1
    assert column_number("Z") == 26
    assert column_number("AA") == 27
    assert column_number("AAA") == 703
