"""
Create a Difference class with a constructor that saves an array of non-negative 
integers. Implement a computeDifference method to find and store the maximum 
absolute difference between any two numbers in the array. The class should be 
able to handle arrays of up to 10 elements. The method computes the largest 
difference by checking each element against others in the array.

Example:
For an input array [1, 2, 5], the computeDifference method would calculate the 
absolute differences (1, 4, 3) and determine that the maximum absolute difference 
is 4 (between 1 and 5).
"""


class Difference:
    def __init__(self, a):
        self.__elements = a
        self.maximumDifference = 0

    def computeDifference(self):
        if self.__elements:
            self.maximumDifference = max(
                abs(x - y) for x in self.__elements for y in self.__elements
            )
