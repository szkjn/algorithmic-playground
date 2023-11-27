"""
Create a Python program that constructs a linked list. The program should 
define a Node class representing list elements, and a Solution class to 
manage the list. Implement an insert method in Solution to add nodes to 
the list's end, and a display method to print the list. The user inputs 
the number of elements (T) and each element's value, which are added 
sequentially to the list.
"""


class Node:
    def __init__(self, data):
        self.data = data
        self.next = None


class Solution:
    def display(self, head):
        current = head
        while current:
            print(current.data, end=" ")
            current = current.next

    def insert(self, head, data):
        new_node = Node(data)
        if head is None:
            return new_node
        else:
            current = head
            while current.next:
                current = current.next
            current.next = new_node
            return head


mylist = Solution()
T = int(input())
head = None
for i in range(T):
    data = int(input())
    head = mylist.insert(head, data)
mylist.display(head)
