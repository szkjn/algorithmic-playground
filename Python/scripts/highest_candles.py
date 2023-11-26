"""
You are in charge of the cake for a child's birthday. The cake will have one candle for
each year of their total age. They will only be able to blow out the tallest of the
candles. Count how many candles are tallest.

Example:
- candles = [4, 4, 1, 3]

The maximum height candles are 4 units high. There are 2 of them, so return 2.

Function Parameter:
- int candles[n]: an array of candles heights

Returns:
- int: number of candles that are tallest

Input Format:
- The first line contains a single integer, n, the number (size) of candles.
- The second line contains n space-separated integers, where each integer i describes
the height of candles[i].
"""


def birthday_cake_candles(candles: list) -> int:
    res = 0
    max_height = max(candles)

    for el in candles:
        if el == max_height:
            res += 1

    return res


if __name__ == "__main__":
    candlesCount = input("Enter total number of candles:\n")
    candlesTemp = input(f"Enter {candlesCount} candle's heights:\n")
    candles = candlesTemp.split(" ")

    res = birthday_cake_candles(candles)
    print("Total tallest candle(s):", res)
