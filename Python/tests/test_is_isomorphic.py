import sys
import os

sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), "..")))

from is_isomorphic import is_isomorphic


def test_is_isomorphic():
    assert is_isomorphic("abb", "cdd") == True
    assert is_isomorphic("cassidy", "1234567") == False
    assert is_isomorphic("cass", "1233") == True
