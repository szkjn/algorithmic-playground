import sys
import os

sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), "..")))

import pytest
from scripts.how_many_opened_doors import pass_doors


@pytest.mark.parametrize(
    "n, nbr_of_passes, expected",
    [
        (7, 3, 4),
        (12, 12, 4),
        (2, 5, 2),
        (124, 5, 68),
    ],
)
def test_pass_doors(n, nbr_of_passes, expected):
    assert pass_doors(n, nbr_of_passes) == expected
