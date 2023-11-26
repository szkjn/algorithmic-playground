import sys
import os

sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), "..")))

from scripts.how_many_opened_doors import pass_doors


def test_pass_doors():
    assert pass_doors(7, 3) == 4
    assert pass_doors(12, 12) == 4
    assert pass_doors(2, 5) == 2
    assert pass_doors(124, 5) == 68
