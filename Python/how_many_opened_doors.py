"""
Interview question of the week by Cassidy Williams (issue #270) :

"Let’s say you have `n` doors that start out as closed. With the first 
pass across the doors, you toggle every door open. With the second 
pass, you toggle every second door. With the third, every third door, 
and so on. Write a function that takes in an integer `numberOfPasses`, 
and returns how many doors are open after the number of passes. 
Thanks Max for inspiring this question!"
"""


def toggle(door):
    door = 0 if door == 1 else 1
    return door


def pass_doors(n, nbr_of_passes):
    doors = [1 for el in range(n)]

    for el in range(1, nbr_of_passes + 1):
        if el == 1:
            for i, door in enumerate(doors):
                doors[i] = toggle(door)
        else:
            for i in range(0, len(doors), el):
                if i > 0:
                    doors[i - 1] = toggle(doors[i - 1])

    count = doors.count(0)
    return count


n = 7
nbr_of_passes = 3

pass_doors(n, nbr_of_passes)

"""
Explanation:
0 means open, 1 means closed
Initial: 1 1 1 1 1 1 1
Pass 1:  0 0 0 0 0 0 0
Pass 2:  0 1 0 1 0 1 0
Pass 3:  0 1 1 1 0 0 0
"""
