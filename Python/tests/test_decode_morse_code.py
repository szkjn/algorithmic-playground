import sys
import os

sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), "..")))

import pytest
from scripts.decode_morse_code import decode_morse


@pytest.mark.parametrize(
    "morse_code, expected",
    [
        (".... . -.--   .--- ..- -.. .", "HEY JUDE"),
        ("...---...", "SOS"),
        ("...   ---   ...", "S O S"),
    ],
)
def test_decode_morse(morse_code, expected):
    assert decode_morse(morse_code) == expected
