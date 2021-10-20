def main():
    result = 0

    for a in range(100,1000):
        for b in range(100, 1000):
            num = a * b
            if ispanlindrome(str(num)) and (num > result):
                result = num

    print(f"Answer: {result}")

def ispanlindrome(num: str) -> bool:
    """
        Description:    Check if a string is a panlindrome.
        Parameters:     num             - str
        Return:         a boolean value 
        Example of Usage:
        >> ispalindrome("racecar")
        ans =
            True 
        >> ispalindrome("dog")
        ans =
            False
    """ 
    return num == num[::-1]


if __name__ == '__main__':
    main()
