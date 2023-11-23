"""
Interview question of the week by Cassidy Williams (issue #286) :

Spreadsheet programs often use the A1 Reference Style to refer to
columns. Given a column name in this style, return its column number.

Examples of column names to their numbers:
A -> 1
B -> 2
C -> 3
Z -> 26
AA -> 27
AB -> 28 
AAA -> 703
"""

alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

def column_number(col_name):
	col_num = 0
	for c in col_name:
		col_num = col_num * len(alphabet) + alphabet.index(c) + 1

	return col_num