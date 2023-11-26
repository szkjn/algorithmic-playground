import sys
import os

sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), "..")))

from scripts.decode_morse_code import decode_morse


def test_decode_morse():
    assert decode_morse(".... . -.--   .--- ..- -.. .") == "HEY JUDE"
    assert decode_morse("...---...") == "SOS"
    assert decode_morse("...   ---   ...") == "S O S"
