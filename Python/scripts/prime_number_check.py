import math

def is_prime(n):
    if n <= 1:
        return False
    for i in range(2, int(math.sqrt(n)) + 1):
        if n % i == 0:
            return False
    return True

def main():
    # Read the number of test cases
    num_cases = int(input("Enter number of test cases: "))
    for _ in range(num_cases):
        # Read the number to check for primality
        num = int(input("Enter a number to check for primality: "))
        if is_prime(num):
            print("Prime")
        else:
            print("Not prime")

if __name__ == "__main__":
    main()
