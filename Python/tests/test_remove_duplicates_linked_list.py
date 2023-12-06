import sys
import os
import pytest

PROJ_PATH = os.path.abspath(os.path.join(os.path.dirname(__file__), "../.."))

sys.path.append(PROJ_PATH)

from Python.tests.test_utils import load_test_data
from Python.scripts.remove_duplicates_linked_list import Solution

test_data = load_test_data("remove_duplicates_linked_list.json")


def list_to_array(head):
    """
    Helper function to convert linked list to array for easy comparison.
    """
    arr = []
    current = head
    while current:
        arr.append(current.data)
        current = current.next
    return arr


@pytest.mark.parametrize(
    "n, data, expected",
    [(item["n"], item["data"], item["expected"]) for item in test_data],
)
def test_remove_duplicates_linked_list(n, data, expected):
    sol = Solution()
    head = None
    for el in data:
        head = sol.insert(head, el)
    head = sol.removeDuplicates(head)
    assert list_to_array(head) == expected
