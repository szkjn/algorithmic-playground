"""
Create a function level_order to perform a level-order traversal (breadth-first 
search) on a Binary Search Tree. It should print each node's data as 
space-separated integers, traversing each level of the tree from left to right, 
top to bottom. The function takes one parameter, root, which is a pointer to 
the root of the tree. No return value is expected.
"""


class Node:
    def __init__(self, data):
        self.right = self.left = None
        self.data = data


class Solution:
    def insert(self, root, data):
        if root == None:
            return Node(data)
        else:
            if data <= root.data:
                cur = self.insert(root.left, data)
                root.left = cur
            else:
                cur = self.insert(root.right, data)
                root.right = cur
        return root

    def level_order(self, root):
        queue = []
        if root is not None:
            queue.append(root)

        # Continue until the queue is empty
        while len(queue) > 0:
            # Dequeue the front node and print its data
            node = queue.pop(0)
            print(node.data, end=" ")

            # Enqueue the left child
            if node.left is not None:
                queue.append(node.left)

            # Enqueue the right child
            if node.right is not None:
                queue.append(node.right)


T = int(input())
myTree = Solution()
root = None

for i in range(T):
    data = int(input())
    root = myTree.insert(root, data)

myTree.level_order(root)
