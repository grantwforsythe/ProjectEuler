def main():
    num = 1
    for i in range(1, 21):
        num = lcm(num, i)

    print(f"Answer: {num}")


def lcm(a: int, b:int) -> int:
    """
        Description:    Find the least common multiple between two numbers. 
        Parameters:     a               - int
                        b               - int
        Return:         the least common multiple 
        Example of Usage:
        >> lcm(3, 5)
        ans =
            15 
    """ 
    return a*b // gcd(a, b)


def gcd(a: int, b:int) -> int:
    """
        Description:    Find the greatest common divisor between two numbers. 
        Parameters:     a               - int
                        b               - int
        Return:         the greatest common divisor
        Example of Usage:
        >> gcd(18, 24)
        ans =
            6
    """ 
    if b == 0:
        return a
    else:
        return gcd(b, a % b)


if __name__ == '__main__':
    main()