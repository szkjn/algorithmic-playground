"""
Create a function removeDuplicates that removes duplicate nodes from a linked list. 
The function receives a head node of a linked list. The input list's nodes contain 
integer data and are sorted in non-decreasing order. The function should modify the 
list to remove all duplicate elements and return the head of the modified list. The 
list may be empty (null head node). Ensure to properly reset pointers after 
deletions to maintain list integrity.

Input format:
- first line contains an integer, n, the number of nodes to be inserted.
- n subsequent lines each contain an integer describing the data value of a node 
being inserted at the list's tail.

Output Format:
- returns the head node of the modified linked list.
"""


class Node:
    def __init__(self, data):
        self.data = data
        self.next = None


class Solution:
    def insert(self, head, data):
        p = Node(data)
        if head is None:
            head = p
        elif head.next is None:
            head.next = p
        else:
            start = head
            while start.next is not None:
                start = start.next
            start.next = p
        return head

    def display(self, head):
        current = head
        while current:
            print(current.data, end=" ")
            current = current.next

    def removeDuplicates(self, head):
        current = head
        while current and current.next:
            if current.data == current.next.data:
                current.next = current.next.next
            else:
                current = current.next
        return head


# mylist = Solution()
# T = int(input())
# head = None
# for i in range(T):
#     data = int(input())
#     head = mylist.insert(head, data)
# head = mylist.removeDuplicates(head)
# mylist.display(head)
